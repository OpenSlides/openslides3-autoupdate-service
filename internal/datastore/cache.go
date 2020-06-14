package datastore

import (
	"encoding/json"
	"sync"
)

type cache struct {
	mu   sync.RWMutex
	data map[string]json.RawMessage
}

func (c *cache) update(changed map[string]json.RawMessage) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.data == nil {
		c.data = make(map[string]json.RawMessage)
	}

	for k, v := range changed {
		if v == nil {
			delete(c.data, k)
			continue
		}

		c.data[k] = v
	}
}

// forkeys returns all data for the given keys.
//
// If a key does not exist in the cache, the returned data contains that key
// with a nil value.
//
// Creates a copy of all data.
func (c *cache) forKeys(keys ...string) map[string]json.RawMessage {
	c.mu.RLock()
	defer c.mu.RUnlock()

	data := make(map[string]json.RawMessage, len(keys))
	for _, key := range keys {
		v := c.data[key]
		data[key] = append(v[:0:0], v...)
	}
	return data
}

// all returns all data from the cache.
//
// Creates a copy of all data.
func (c *cache) all() map[string]json.RawMessage {
	c.mu.RLock()
	defer c.mu.RUnlock()

	data := make(map[string]json.RawMessage, len(c.data))
	for k, v := range c.data {
		data[k] = append(v[:0:0], v...)
	}
	return data
}
