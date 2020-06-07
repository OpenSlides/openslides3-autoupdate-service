package autoupdate

import (
	"encoding/json"
)

// Datastore holds the current data.
type Datastore interface {
	LowestID() int
	KeysChanged() ([]string, int, error)
	GetMany([]string) map[string]json.RawMessage
	GetAll() map[string]json.RawMessage
}
