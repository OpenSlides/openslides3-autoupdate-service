package test

import (
	"encoding/json"
)

// RedisMock implements the datastore.RedisConn interface.
type RedisMock struct {
	FD                map[string]json.RawMessage
	Min               int
	Max               int
	send              chan []byte
	ChangedKeysResult []string
}

// NewRedisMock initializes a RedisMock.
func NewRedisMock() *RedisMock {
	return &RedisMock{
		send: make(chan []byte, 1),
	}
}

// FullData returns the given values.
func (r *RedisMock) FullData() (data map[string]json.RawMessage, max int, min int, err error) {
	return r.FD, r.Max, r.Min, nil
}

// Update waits for Send() to be called and returned the send data.
func (r *RedisMock) Update(closing chan struct{}) ([]byte, error) {
	select {
	case <-closing:
		return nil, closingErr{}
	case v := <-r.send:
		return v, nil
	}
}

// ChangedKeys returnes ChangedKeysResult.
func (r *RedisMock) ChangedKeys(from, to int) ([]string, error) {
	return r.ChangedKeysResult, nil
}

// Data returns the keys from FD.
func (r *RedisMock) Data(keys []string) (map[string]json.RawMessage, error) {
	data := make(map[string]json.RawMessage, len(keys))
	for _, key := range keys {
		data[key] = r.FD[key]
	}
	return data, nil
}

// Send sends a value that can be received with Update.
func (r *RedisMock) Send(value []byte) {
	r.send <- value
}
