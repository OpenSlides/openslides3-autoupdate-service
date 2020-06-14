package autoupdate

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

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
	closed     chan struct{}
	topic      *topic.Topic
}

// New create a new autoupdate instance.
func New(datastore Datastore, restricter Restricter) (*Autoupdate, error) {
	closed := make(chan struct{})

	a := &Autoupdate{
		datastore:  datastore,
		closed:     closed,
		restricter: restricter,
		topic:      topic.New(topic.WithClosed(closed), topic.WithStartID(uint64(datastore.CurrentID()))),
	}

	go func() {
		for {
			keys, changeID, err := datastore.KeysChanged(a.closed)
			if err != nil {
				var closing interface {
					Closing()
				}
				if errors.As(err, &closing) {
					return
				}

				log.Printf("Error: Can not receive new data: %v", err)
				continue
			}

			tid := a.topic.Publish(keys...)

			// When the received data is to new, all the missing data is received at
			// once. If more then one change id was skipped, the missing ids have to be
			// created in the toppic with dummy items.
			for tid < uint64(changeID) {
				tid = a.topic.Publish()
			}

			// if the topic id is bigger then the change id, then something is broken.
			// There is no way to recover safely from this.
			if tid != uint64(changeID) {
				log.Panicf("topic id differs from change id. Topic: %d, changeid %d", tid, changeID)
			}
		}
	}()

	return a, nil
}

// Close stops the autoupdate service.
//
// This method is not save for concurrent use. It can only be called once.
func (a *Autoupdate) Close() {
	close(a.closed)
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

	// The connection was closed or the server is exiting.
	if changedKeys == nil {
		return false, nil, 0, nil
	}

	data := a.datastore.GetMany(changedKeys)
	a.restricter.Restrict(uid, data)
	return false, data, int(newChangeID), nil
}
