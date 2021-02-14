package autoupdate

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ostcar/topic"
)

// Autoupdate holds the state of the serice.
//
// The Service caches all data in its newest version. This means, that only the
// newst data ist returned, even when old change ids are requested. The change
// id system only desices, which keys have changed.
type Autoupdate struct {
	datastore  Datastore
	restricter Restricter
	closed     <-chan struct{}
	topic      *topic.Topic

	pccMu                    sync.Mutex
	projectorConnectionCount int
}

// New create a new autoupdate instance.
func New(datastore Datastore, restricter Restricter, closed <-chan struct{}) (*Autoupdate, error) {
	a := &Autoupdate{
		datastore:  datastore,
		closed:     closed,
		restricter: restricter,
		topic:      topic.New(topic.WithClosed(closed), topic.WithStartID(uint64(datastore.CurrentID()))),
	}

	go func() {
		for {
			keys, changeID, err := datastore.KeysChanged()
			if err != nil {
				var closing interface {
					Closing()
				}
				if errors.As(err, &closing) {
					return
				}

				var reset interface {
					Reset()
				}
				if errors.As(err, &reset) {
					a.reset()
					continue
				}

				log.Printf("Autoupdate: Can not receive new data: %v", err)
				time.Sleep(5 * time.Second)
				continue
			}

			// When the received data is to new, all the missing data is
			// received at once. If more then one change id was skipped, the
			// missing ids have to be created in the topic with dummy items.
			tid := a.topic.LastID()
			for tid < uint64(changeID)-1 {
				tid = a.topic.Publish()
			}

			tid = a.topic.Publish(keys...)

			// if the topic id is different then the change id, then something
			// is broken. There is no way to recover safely from this.
			if tid != uint64(changeID) {
				log.Panicf("topic id differs from change id. Topic: %d, changeid %d", tid, changeID)
			}
		}
	}()

	return a, nil
}

// Receive returns all changed data and the new changeid since the given change
// id. If there is no new data, then this method blocks until the context is
// done, the service is closed or new data is received.
//
// If the returned value is nil, then the context or the service was closed.
//
// The returned data is restricted for the given uid.
func (a *Autoupdate) Receive(ctx context.Context, uid int, changeID int) (bool, map[string]json.RawMessage, int, error) {
	if changeID == 0 || changeID < a.datastore.LowestID() {
		// The changeID is lower then the lowest change_id in redis. Return all data.
		data := a.datastore.GetAll()
		a.restricter.Restrict(uid, data)
		return true, data, int(a.topic.LastID()), nil
	}

	newChangeID, changedKeys, err := a.topic.Receive(ctx, uint64(changeID))
	if err != nil {
		var unknownID topic.UnknownIDError
		if !errors.As(err, &unknownID) {
			return false, nil, 0, fmt.Errorf("get changed keys from topic: %w", err)
		}

		// ID is not in memory, ask redis.
		newChangeID = unknownID.FirstID
		changedKeys, err = a.datastore.ChangedKeys(changeID, int(newChangeID))
		if err != nil {
			return false, nil, 0, fmt.Errorf("get changed keys from redis: %w", err)
		}
	}

	if len(changedKeys) == 0 {
		return false, nil, int(newChangeID), nil
	}

	data := a.datastore.GetMany(changedKeys)
	a.restricter.Restrict(uid, data)
	return false, data, int(newChangeID), nil
}

// Projectors returns the renderd data for a list of projectors. The attribute
// pids i sthe list of requested projectors. Only the projectors that changed
// are returned.
//
// The attribute tid tells the last version the caller has seen. 0 Means, no
// data has to be seen, so all requested projectors are returned.
//
// This method blocks until the service is closed, the given context exists or
// there are data to return.
func (a *Autoupdate) Projectors(ctx context.Context, tid uint64, pids []int) (ntid uint64, data map[int]json.RawMessage, cid int, err error) {
	rdata := make(map[int]json.RawMessage)
	for len(rdata) == 0 {
		tid, data, err = a.datastore.ProjectorData(ctx, tid)
		if err != nil {
			return 0, nil, 0, err
		}

		for _, pid := range pids {
			v, ok := data[pid]
			if !ok {
				// Requested projector does not exist.
				continue
			}

			rdata[pid] = v
		}
	}
	return tid, rdata, int(a.topic.LastID()), nil
}

func (a *Autoupdate) reset() {
	oldTopic := a.topic
	a.topic = topic.New(topic.WithClosed(a.closed), topic.WithStartID(uint64(a.datastore.CurrentID())))

	// Send an empty message on the old topic to wake up all clients.
	oldTopic.Publish()
}
