package datastore

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type config struct {
	mu     sync.RWMutex
	values map[string]json.RawMessage
}

func (c *config) update(data map[string]json.RawMessage) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.values == nil {
		c.values = make(map[string]json.RawMessage)
	}

	for k, v := range data {
		parts := strings.Split(k, ":")
		if len(parts) != 2 {
			return fmt.Errorf("key %s has wrong format. Expected one `:`", k)
		}

		if parts[0] != "core/config" {
			continue
		}

		id, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("key %s has wrong format. Expected id to be int not %s", k, parts[1])
		}

		var decoded struct {
			Key   string          `json:"key"`
			Value json.RawMessage `json:"value"`
		}
		if err := json.Unmarshal(v, &decoded); err != nil {
			return fmt.Errorf("decoding core/config:%d: %w", id, err)
		}

		c.values[decoded.Key] = decoded.Value
	}
	return nil
}

func (c *config) ConfigValue(key string, v interface{}) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return json.Unmarshal(c.values[key], v)
}
