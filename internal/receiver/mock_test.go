package receiver_test

import "encoding/json"

type redisConnMock struct {
}

func (r *redisConnMock) FullData() (map[string]json.RawMessage, int, int, error) {
	data := map[string]json.RawMessage{
		"user:1": []byte("hello world"),
	}
	return data, 1, 2, nil
}

func (r *redisConnMock) Update() ([]byte, error) {
	return nil, nil
}
