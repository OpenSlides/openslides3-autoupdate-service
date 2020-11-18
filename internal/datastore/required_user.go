package datastore

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

type requiredUser struct {
	mu sync.RWMutex

	// requiredUser maps element ids to a set of needed user ids.
	requiredUser map[string]map[int]bool
	// requiredPerm maps element ids to the required permission to have.
	requiredPerm map[string]string

	// users maps every required user id to a list of permissions, so
	// for each user it could be easily queried, which permissions are
	// needed to require the user.
	users map[int][]string

	callables map[string]func(json.RawMessage) (map[int]bool, string, error)
}

func (r *requiredUser) update(data map[string]json.RawMessage) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.requiredUser == nil {
		r.requiredUser = make(map[string]map[int]bool)
		r.requiredPerm = make(map[string]string)
	}

	var usersChanged bool

	for elementID, v := range data {
		parts := strings.Split(elementID, ":")
		if len(parts) != 2 {
			return fmt.Errorf("key %s has wrong format. Expected one `:`", elementID)
		}

		c, ok := r.callables[parts[0]]
		if !ok {
			continue
		}

		usersChanged = true

		if v == nil {
			delete(r.requiredUser, elementID)
			delete(r.requiredPerm, elementID)
			continue
		}

		uids, perm, err := c(v)
		if err != nil {
			return fmt.Errorf("calculate required users for %s: %w", elementID, err)
		}

		r.requiredUser[elementID] = uids
		r.requiredPerm[elementID] = perm
	}

	if !usersChanged {
		return nil
	}

	r.users = make(map[int][]string)

	for elementID, v := range r.requiredUser {
		for uid := range v {
			r.users[uid] = append(r.users[uid], r.requiredPerm[elementID])
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
