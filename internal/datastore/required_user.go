package datastore

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

type requiredUser struct {
	mu       sync.RWMutex
	users    map[int]bool
	requires map[string][]int

	callables map[string]func(json.RawMessage) ([]int, error)
}

func (r *requiredUser) update(data map[string]json.RawMessage) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users = make(map[int]bool)

	for k, v := range data {
		parts := strings.Split(k, ":")
		if len(parts) != 2 {
			return fmt.Errorf("key %s has wrong format. Expected one `:`", k)
		}

		c, ok := r.callables[parts[0]]
		if !ok {
			continue
		}

		uids, err := c(v)
		if err != nil {
			return fmt.Errorf("calculate required users for %s: %w", k, err)
		}

		for _, uid := range uids {
			r.users[uid] = true
		}

	}

	return nil
}

func (r *requiredUser) UserRequired(uid int) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.users[uid]
}
