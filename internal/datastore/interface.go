package datastore

import (
	"encoding/json"
)

// RedisConn calls all needed redis commands.
type RedisConn interface {
	FullData() (data map[string]json.RawMessage, max int, min int, err error)
	Update(<-chan struct{}) ([]byte, error)
	ChangedKeys(from, to int) ([]string, error)
	Data(keys []string) (map[string]json.RawMessage, error)
}

// ProjectorCallable knows how to build the projector data for an element.
type ProjectorCallable interface {
	Build(element json.RawMessage, pid int) (json.RawMessage, error)
}
