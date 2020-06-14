package autoupdate

import (
	"encoding/json"
)

// Datastore holds the current data.
type Datastore interface {
	LowestID() int
	CurrentID() int
	KeysChanged(chan struct{}) ([]string, int, error)
	GetMany([]string) map[string]json.RawMessage
	GetAll() map[string]json.RawMessage
	ChangedKeys(from, to int) ([]string, error)
}

// Restricter restricts data for one user.
type Restricter interface {
	Restrict(uid int, data map[string]json.RawMessage)
}
