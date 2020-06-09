package restricter

import (
	"encoding/json"
	"log"
	"strings"
)

// Restricter can restrict some data for an user.
type Restricter struct {
	datastore Datastore
	elements  map[string]Element
}

// New initializes a Restricter
func New(datastore Datastore, elements map[string]Element) *Restricter {
	return &Restricter{
		datastore: datastore,
		elements:  elements,
	}
}

// Restrict changes the data for the given user.
func (r *Restricter) Restrict(uid int, data map[string]json.RawMessage) {
	for k, v := range data {
		parts := strings.Split(k, ":")

		e, ok := r.elements[parts[0]]
		if !ok {
			delete(data, k)
			continue
		}

		restricted, err := e.Restrict(uid, v)
		if err != nil {
			log.Printf("Can not restrict key %s for user %d: %v", k, uid, err)
			delete(data, k)
			continue
		}

		if restricted == nil {
			delete(data, k)
			continue
		}
		data[k] = restricted
	}
}

// ElementFunc converts a simple element restricter func to a element
// restricter.
type ElementFunc func(int, json.RawMessage) (json.RawMessage, error)

// Restrict calls the ElementFunc.
func (f ElementFunc) Restrict(u int, data json.RawMessage) (json.RawMessage, error) {
	return f(u, data)
}

// BasePermission returns a generator to create simple Elements that only check
// one permission.
func BasePermission(h HasPermer) func(perm string) ElementFunc {
	return func(perm string) ElementFunc {
		return func(u int, data json.RawMessage) (json.RawMessage, error) {
			if h.HasPerm(u, perm) {
				return data, nil
			}
			return nil, nil
		}
	}
}

// ForAll gets read access for everybody.
var ForAll ElementFunc = func(_ int, data json.RawMessage) (json.RawMessage, error) {
	return data, nil
}
