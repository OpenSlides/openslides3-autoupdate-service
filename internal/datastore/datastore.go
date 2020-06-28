// Package datastore receives and caches the data from redis and the worker.
package datastore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

// Datastore holds the connection to OpenSlides and Redis.
type Datastore struct {
	osAddr      string
	redisConn   RedisConn
	cache       *cache
	minChangeID int
	closed      <-chan struct{}

	mu          sync.RWMutex
	maxChangeID int

	hasPerm
	requiredUser
	projectors
	config
}

// New returns an initialized Datastore instance.
func New(osAddr string, redisConn RedisConn, requiredUsers map[string]func(json.RawMessage) ([]int, string, error), projectorSlides map[string]projector.Callable, closed <-chan struct{}) (*Datastore, error) {
	fd, max, min, err := redisConn.FullData()
	if err != nil {
		return nil, fmt.Errorf("get startdata from redis: %w", err)
	}

	d := &Datastore{
		osAddr:       osAddr,
		redisConn:    redisConn,
		cache:        new(cache),
		minChangeID:  min,
		maxChangeID:  max,
		requiredUser: requiredUser{callables: requiredUsers},
		closed:       closed,
	}

	// TODO: fix circular dependency between datastore and projector.
	d.projectors = projectors{ds: d, callables: projectorSlides, closed: closed}

	if err := d.update(fd, max); err != nil {
		return nil, fmt.Errorf("initial datastore update: %w", err)
	}

	return d, nil
}

// LowestID returns the lowest id in the datastore.
func (d *Datastore) LowestID() int {
	return d.minChangeID
}

// CurrentID returns the highest id in the datastore.
func (d *Datastore) CurrentID() int {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.maxChangeID
}

// KeysChanged blocks until there is new data. It updates the internal cache and
// returns the changed keys and the new change id.
//
// If closing is closed then it return nil, 0, nil.
func (d *Datastore) KeysChanged(closing <-chan struct{}) ([]string, int, error) {
	var rawData []byte
	var err error
	for rawData == nil {
		// Update() blocks until there is new data. But when there is no new
		// data for an hour, then it returns with nil.
		rawData, err = d.redisConn.Update(closing)
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
	keys := make([]string, 0, len(sData.Elements))
	for k, v := range sData.Elements {
		if bytes.Equal(v, []byte(`null`)) {
			// Deleted elements.
			sData.Elements[k] = nil
		}
		keys = append(keys, k)
	}

	if changeID > d.maxChangeID+1 {
		// Data is to new. Get the data in between.
		data, err := d.receive(d.maxChangeID, changeID-1)
		if err != nil {
			return nil, 0, fmt.Errorf("receive missing data from %d to %d: %w", d.maxChangeID, changeID-1, err)
		}

		for k := range data {
			keys = append(keys, k)
		}

		if err := d.update(data, changeID-1); err != nil {
			return nil, 0, fmt.Errorf("updating cache from %d to %d: %w", d.maxChangeID, changeID-1, err)
		}
	}

	if changeID < d.maxChangeID+1 {
		// Data already known. Try the next.
		return d.KeysChanged(closing)
	}

	if err := d.update(sData.Elements, changeID); err != nil {
		return nil, 0, fmt.Errorf("updating cache: %w", err)
	}

	return keys, changeID, nil
}

// ChangedKeys returns the keys that have changed between from and to from
// redis. from is not inclusive, to is inclusiv.
func (d *Datastore) ChangedKeys(from, to int) ([]string, error) {
	keys, err := d.redisConn.ChangedKeys(from, to)
	if err != nil {
		return nil, fmt.Errorf("get changed keys: %w", err)
	}
	return keys, err
}

// Get sets the attribute v to the value the collection:id. Returns an error
// with the method `DoesNotExist()` if the value does not exist.
//
// v has to be a pointer.
func (d *Datastore) Get(collection string, id int, v interface{}) error {
	e := d.cache.get(fmt.Sprintf("%s:%d", collection, id))
	if e == nil {
		return doesNotExist(fmt.Sprintf("%s:%d", collection, id))
	}
	return json.Unmarshal(e, v)
}

// GetMany returns the values for the given keys.
func (d *Datastore) GetMany(keys []string) map[string]json.RawMessage {
	return d.cache.forKeys(keys...)
}

// GetCollection gets all elements of one collection.
func (d *Datastore) GetCollection(collection string) []json.RawMessage {
	// TODO: maybe build an index?

	var elements []json.RawMessage
	prefix := collection + ":"
	for key, value := range d.cache.all() {
		if !strings.HasPrefix(key, prefix) {
			continue
		}
		elements = append(elements, value)
	}
	return elements
}

// GetModels returns each element from collection that is in the ids slide.
func (d *Datastore) GetModels(collection string, ids []int) []json.RawMessage {
	// TODO: maybe build an index?

	set := make(map[int]bool)
	for _, id := range ids {
		set[id] = true
	}

	var elements []json.RawMessage
	for key, value := range d.cache.all() {
		parts := strings.Split(key, ":")
		if len(parts) != 2 || parts[0] != collection {
			continue
		}

		id, err := strconv.Atoi(parts[1])
		if err != nil || !set[id] {
			continue
		}

		elements = append(elements, value)
	}
	return elements
}

// GetAll returns all data.
func (d *Datastore) GetAll() map[string]json.RawMessage {
	return d.cache.all()
}

// update updates the cache. It is not save for concourent use.
func (d *Datastore) update(data map[string]json.RawMessage, changeID int) error {
	d.cache.update(data)

	d.mu.Lock()
	d.maxChangeID = changeID
	d.mu.Unlock()

	if err := d.hasPerm.update(data); err != nil {
		return fmt.Errorf("updating hasPerm: %w", err)
	}

	if err := d.requiredUser.update(data); err != nil {
		return fmt.Errorf("updating requiredUser: %w", err)
	}

	if err := d.config.update(data); err != nil {
		return fmt.Errorf("updating config: %w", err)
	}

	if err := d.projectors.update(data); err != nil {
		return fmt.Errorf("update projectors: %w", err)
	}

	log.Printf("Datastore on change_id %d", changeID)
	return nil
}

// receive is used to get missing data. It returns all data between higher
// "from" and lower or equal "to".
func (d *Datastore) receive(from, to int) (data map[string]json.RawMessage, err error) {
	keys, err := d.redisConn.ChangedKeys(from, to)
	if err != nil {
		return nil, fmt.Errorf("get changed keys: %w", err)
	}

	if len(keys) == 0 {
		return nil, nil
	}

	data, err = d.redisConn.Data(keys)
	if err != nil {
		return nil, fmt.Errorf("get data: %w", err)
	}
	return data, nil
}
