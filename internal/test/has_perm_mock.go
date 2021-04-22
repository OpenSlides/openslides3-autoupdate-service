package test

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// HasPermMock implements the restricter.HasPermer interface.
type HasPermMock struct {
	IsSuperuser bool
	Perms       []string
	Groups      map[int]bool
	Data        map[string]json.RawMessage
}

// HasPerm returns true, if the given perm is in the list of Perms.
func (h *HasPermMock) HasPerm(_ int, perm string) bool {
	if h.IsSuperuser {
		return true
	}
	for _, p := range h.Perms {
		if p == perm {
			return true
		}
	}
	return false
}

// IsSuperadmin returns, if the user is superadmin.
func (h *HasPermMock) IsSuperadmin(_ int) bool {
	return h.IsSuperuser
}

// InGroups checks if one of the id in groups is also in HasPermMock.Groups
func (h *HasPermMock) InGroups(_ int, groups []int) bool {
	if h.IsSuperuser {
		return true
	}

	for _, g := range groups {
		if h.Groups[g] {
			return true
		}
	}
	return false
}

// UserRequired checks for the given groups.
func (h *HasPermMock) UserRequired(uid int) []string {
	perms := make(map[string]bool)
	for _, e := range exampleRequiredUser {
		for _, id := range e.ids {
			if id == uid {
				perms[e.perm] = true
			}
		}
	}

	keys := make([]string, 0, len(perms))
	for k := range perms {
		keys = append(keys, k)
	}
	return keys
}

// Get tells, if the given user has the ...
func (h *HasPermMock) Get(collection string, id int, v interface{}) error {
	elementID := collection + ":" + strconv.Itoa(id)
	e := h.Data[elementID]
	if e == nil {
		return fmt.Errorf("doesNotExistError: %s", elementID)
	}
	return json.Unmarshal(e, v)
}

// ConfigValue is currently a stub until it is actually needed by the tests
func (h *HasPermMock) ConfigValue(key string, v interface{}) error {
	return nil
}
