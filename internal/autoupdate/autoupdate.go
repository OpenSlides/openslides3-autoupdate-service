package autoupdate

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/ostcar/topic"
)

// Autoupdate holds the state of the serice.
//
// The Service caches all data in its newest version. This means, that only the
// newst data ist returned, even when old change ids are requested. The change
// id system only desices, which keys have changed.
type Autoupdate struct {
	receiver    Receiver
	cache       *cache
	minChangeID int
	closed      chan struct{}
	topic       *topic.Topic

	mu          sync.RWMutex
	maxChangeID int
}

// New create a new autoupdate instance.
func New(receiver Receiver) (*Autoupdate, error) {
	fd, min, max, err := receiver.StartData()
	if err != nil {
		return nil, fmt.Errorf("receiving startup data: %w", err)
	}

	closed := make(chan struct{})

	a := &Autoupdate{
		receiver:    receiver,
		minChangeID: min,
		maxChangeID: max,
		cache:       &cache{data: fd},
		closed:      closed,
		topic:       topic.New(topic.WithClosed(closed), topic.WithStartID(uint64(max))),
	}

	// Background task, that updates the data.
	go func() {
		for {
			data, changeID, err := a.receiver.Update()
			if err != nil {
				log.Printf("Can not receive updated data: %v", err)
				continue
			}

			if changeID > a.maxChangeID+1 {
				// Data is to new. Get the data in between.
				data, err := a.receiver.Receive(a.maxChangeID, changeID-1)
				if err != nil {
					log.Printf("Can not receive data: %v", err)
					continue
				}
				a.update(data, changeID-1)
			}

			if changeID < a.maxChangeID+1 {
				// Data already known.
				continue
			}

			a.update(data, changeID)
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
func (a *Autoupdate) Receive(ctx context.Context, uid int, changeID int) (map[string]json.RawMessage, int, error) {
	a.mu.RLock()
	maxChangeID := a.maxChangeID
	a.mu.RUnlock()

	if changeID >= maxChangeID {
		nid, keys, err := a.topic.Receive(ctx, uint64(changeID))
		if err != nil {
			return nil, 0, fmt.Errorf("get changed keys: %w", err)
		}

		if keys == nil {
			// service or connection is closed.
			return nil, 0, nil
		}

		return a.cache.forKeys(keys), int(nid), nil
	}

	if changeID == 0 || changeID < a.minChangeID {
		return a.cache.all(), maxChangeID, nil
	}

	keys, err := a.receiver.ChangedKeys(changeID, maxChangeID)
	if err != nil {
		return nil, 0, fmt.Errorf("receive changed keys: %v", err)
	}

	return a.cache.forKeys(keys), maxChangeID, nil
}

// update updates the cache. It is not save for concourent use.
func (a *Autoupdate) update(data map[string]json.RawMessage, changeID int) {
	a.cache.update(data)

	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
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

	a.mu.Lock()
	a.maxChangeID = changeID
	a.mu.Unlock()
}
