package projector

import "encoding/json"

// Datastore are the methods a usual projector slide needs.
type Datastore interface {
	Get(string) json.RawMessage
}

// Callable knows how to build the projector data for an element.
type Callable interface {
	Build(element json.RawMessage, pid int) (json.RawMessage, error)
}

// CallableFunc is a functions that implements the Callable interface.
type CallableFunc func(element json.RawMessage, pid int) (json.RawMessage, error)

// Build builds the data of a projector.
func (f CallableFunc) Build(element json.RawMessage, pid int) (json.RawMessage, error) {
	return f(element, pid)
}
