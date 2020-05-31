package receiver

import "encoding/json"

// RedisConn calls all needed redis commands.
type RedisConn interface {
	FullData() (map[string]json.RawMessage, int, int, error)
	Update() ([]byte, error)
}
