package restricter

import "encoding/json"

// Datastore gives the restricter the full.
type Datastore interface {
	GetMany([]string) map[string]json.RawMessage
	GetAll() map[string]json.RawMessage
}

// Element knows how to restrict one element.
type Element interface {
	Restrict(int, json.RawMessage) (json.RawMessage, error)
}
