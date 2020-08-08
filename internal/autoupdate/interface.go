package autoupdate

import (
	"context"
	"encoding/json"
)

// Datastore holds the current data.
type Datastore interface {
	LowestID() int
	CurrentID() int
	KeysChanged() ([]string, int, error)
	GetMany([]string) map[string]json.RawMessage
	GetAll() map[string]json.RawMessage
	ChangedKeys(from, to int) ([]string, error)
	Projectors(ctx context.Context, tid uint64) (uint64, map[int]json.RawMessage, error)
}

// Restricter restricts data for one user.
type Restricter interface {
	Restrict(uid int, data map[string]json.RawMessage)
}
