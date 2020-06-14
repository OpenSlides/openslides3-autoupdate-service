package datastore

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

const (
	groupDefaultPK = 1
	groupAdminPK   = 2
)

type hasPerm struct {
	mu        sync.RWMutex
	groupPerm map[int]map[string]bool
	userGroup map[int][]int
}

func (h *hasPerm) HasPerm(uid int, perm string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if uid == 0 {
		return h.groupPerm[groupDefaultPK][perm]
	}

	for _, groupID := range h.userGroup[uid] {
		if groupID == groupAdminPK {
			// User is in admin group.
			return true
		}

		if h.groupPerm[groupID][perm] {
			return true
		}
	}
	return false
}

func (h *hasPerm) InGroups(uid int, groups []int) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()

	set := make(map[int]bool)
	for _, gid := range groups {
		set[gid] = true
	}

	for _, groupID := range h.userGroup[uid] {
		if set[groupID] {
			return true
		}
	}

	return h.IsSuperadmin(uid)
}

func (h *hasPerm) IsSuperadmin(uid int) bool {
	for _, groupID := range h.userGroup[uid] {
		if groupID == groupAdminPK {
			return true
		}
	}

	return false
}

func (h *hasPerm) update(data map[string]json.RawMessage) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	for k, v := range data {
		parts := strings.Split(k, ":")
		if len(parts) != 2 {
			return fmt.Errorf("key %s has wrong format. Expected one `:`", k)
		}

		id, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("key %s has wrong format. Expected id to be int not %s", k, parts[1])
		}

		switch parts[0] {
		case "users/user":
			err = h.updateUserGroup(id, v)
		case "users/group":
			err = h.updateGroupPerm(id, v)
		}
		if err != nil {
			return fmt.Errorf("update %s: %w", k, err)
		}
	}
	return nil
}

func (h *hasPerm) updateGroupPerm(id int, data json.RawMessage) error {
	if h.groupPerm == nil {
		h.groupPerm = make(map[int]map[string]bool)
	}

	if data == nil {
		// Group deleted.
		delete(h.groupPerm, id)
		return nil
	}

	var group struct {
		Permissions []string `json:"permissions"`
	}
	if err := json.Unmarshal(data, &group); err != nil {
		return fmt.Errorf("unmarshal group: %w", err)
	}

	// Reset perms. It would be faster to reuse the map
	h.groupPerm[id] = make(map[string]bool)

	for _, perm := range group.Permissions {
		h.groupPerm[id][perm] = true
	}
	return nil
}

func (h *hasPerm) updateUserGroup(id int, data json.RawMessage) error {
	if h.userGroup == nil {
		h.userGroup = make(map[int][]int)
	}

	if data == nil {
		// User deleted.
		delete(h.userGroup, id)
		return nil
	}

	var user struct {
		ID       int   `json:"id"`
		GroupsID []int `json:"groups_id"`
	}

	if err := json.Unmarshal(data, &user); err != nil {
		return fmt.Errorf("unmarshal user: %w", err)
	}

	h.userGroup[id] = user.GroupsID

	if len(user.GroupsID) == 0 {
		h.userGroup[id] = []int{groupDefaultPK}
	}
	return nil
}
