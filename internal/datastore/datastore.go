// Package datastore receives and caches the data from redis and the worker.
package datastore

import (
	"encoding/json"
	"fmt"
	"sync"
)

// Datastore holds the connection to OpenSlides and Redis.
type Datastore struct {
	osAddr      string
	redisConn   RedisConn
	cache       *cache
	minChangeID int

	mu          sync.RWMutex
	maxChangeID int

	hasPerm
}

// New returns an initialized Datastore instance.
func New(osAddr string, redisConn RedisConn) (*Datastore, error) {
	// TODO: Handle case, that data is not ready.
	fd, min, max, err := redisConn.FullData()
	if err != nil {
		return nil, fmt.Errorf("get startdata from redis: %w", err)
	}

	d := &Datastore{
		osAddr:      osAddr,
		redisConn:   redisConn,
		cache:       new(cache),
		minChangeID: min,
		maxChangeID: max,
	}

	d.updateCache(fd, max)

	return d, nil
}

// LowestID returns the highest known change id.
func (d *Datastore) LowestID() int {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.maxChangeID
}

// KeysChanged blocks until there is new data. It updates the internal cache and
// returns the changed keys and the new change id.
func (d *Datastore) KeysChanged() ([]string, int, error) {
	var rawData []byte
	var err error
	for rawData == nil {
		// Update() blocks until there is new data. But when there is no new
		// data for an hour, then it returns with nil.
		rawData, err = d.redisConn.Update()
		if err != nil {
			return nil, 0, fmt.Errorf("get autoupdate data: %w", err)
		}
	}

	var sData struct {
		Elements map[string]json.RawMessage `json:"elements"`
		ChangeID int                        `json:"change_id"`
	}

	if err := json.Unmarshal(rawData, &sData); err != nil {
		return nil, 0, fmt.Errorf("parse data from redis: %w", err)
	}

	changeID := sData.ChangeID

	if changeID > d.maxChangeID+1 {
		// Data is to new. Get the data in between.
		data, err := d.receive(d.maxChangeID, changeID-1)
		if err != nil {
			return nil, 0, fmt.Errorf("receive missing data from %d to %d: %w", d.maxChangeID, changeID-1, err)
		}

		if err := d.updateCache(data, changeID-1); err != nil {
			return nil, 0, fmt.Errorf("updating cache: %w", err)
		}
	}

	if changeID < d.maxChangeID+1 {
		// Data already known. Try the next.
		return d.KeysChanged()
	}

	if err := d.updateCache(sData.Elements, changeID); err != nil {
		return nil, 0, fmt.Errorf("updating cache: %w", err)
	}

	keys := make([]string, 0, len(sData.Elements))
	for k := range sData.Elements {
		keys = append(keys, k)
	}
	return keys, changeID, nil
}

// GetMany returns the values for the given keys.
func (d *Datastore) GetMany(keys []string) map[string]json.RawMessage {
	return d.cache.forKeys(keys...)
}

// GetAll returns all data.
func (d *Datastore) GetAll() map[string]json.RawMessage {
	return d.cache.all()
}

// update updates the cache. It is not save for concourent use.
func (d *Datastore) updateCache(data map[string]json.RawMessage, changeID int) error {
	d.cache.update(data)

	d.mu.Lock()
	d.maxChangeID = changeID
	d.mu.Unlock()

	if err := d.hasPerm.update(data); err != nil {
		return fmt.Errorf("updating hasPerm: %w", err)
	}
	return nil
}

// receive is used to get missing data. It returns all data between higher
// "from" and lower or equal "to".
func (d *Datastore) receive(from, to int) (data map[string]json.RawMessage, err error) {
	keys, err := d.redisConn.ChangedKeys(from, to)
	if err != nil {
		return nil, fmt.Errorf("get changed keys: %w", err)
	}

	data, err = d.redisConn.Data(keys)
	if err != nil {
		return nil, fmt.Errorf("get data: %w", err)
	}
	return data, nil
}
