package datastore

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

type requiredUser struct {
	mu sync.RWMutex

	requiredUser map[string][]int
	requiredPerm map[string]string

	// users holds a calculated version of usersRequired that can be accessed in
	// constant time.
	users map[int][]string

	callables map[string]func(json.RawMessage) ([]int, string, error)
}

func (r *requiredUser) update(data map[string]json.RawMessage) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.requiredUser == nil {
		r.requiredUser = make(map[string][]int)
		r.requiredPerm = make(map[string]string)
	}

	var usersChanged bool

	for k, v := range data {
		parts := strings.Split(k, ":")
		if len(parts) != 2 {
			return fmt.Errorf("key %s has wrong format. Expected one `:`", k)
		}

		c, ok := r.callables[parts[0]]
		if !ok {
			continue
		}

		usersChanged = true

		if v == nil {
			delete(r.requiredUser, k)
			delete(r.requiredPerm, k)
			continue
		}

		uids, perm, err := c(v)
		if err != nil {
			return fmt.Errorf("calculate required users for %s: %w", k, err)
		}

		r.requiredUser[k] = uids
		r.requiredPerm[k] = perm
	}

	if !usersChanged {
		return nil
	}

	r.users = make(map[int][]string)

	for k, v := range r.requiredUser {
		for _, uid := range v {
			// Make sure each perm is only once in the slide. For smal slides,
			// this is faster then using a set.
			var inSlide bool
			for _, perm := range r.users[uid] {
				if perm == r.requiredPerm[k] {
					inSlide = true
					break
				}
			}
			if !inSlide {
				r.users[uid] = append(r.users[uid], r.requiredPerm[k])
			}
		}
	}

	return nil
}

// UserRequired tells, if the uid is required by another element.
//
// An empty result means, that the uid is not required.
//
// If the user is required, a list of permissions is returned. If the client has
// one of the permission, it should see the user.
func (r *requiredUser) UserRequired(uid int) []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.users == nil {
		return nil
	}

	return r.users[uid]
}
