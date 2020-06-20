package redis

import (
	"errors"
	"fmt"
)

var errNil = errors.New("nil returned")

// stream parses a redis stream object.
func stream(reply interface{}, err error) (string, []byte, error) {
	if err != nil {
		return "", nil, err
	}
	if reply == nil {
		return "", nil, errNil
	}
	streams, ok := reply.([]interface{})
	if !ok {
		return "", nil, fmt.Errorf("invalid input. Data has to be a list, not %T", reply)
	}
	if len(streams) == 0 {
		return "", nil, fmt.Errorf("invalid input. No stream in data")
	}
	stream1, ok := streams[0].([]interface{})
	if !ok {
		return "", nil, fmt.Errorf("invalid input. Stream has to be a two-tuple, not %T", streams[0])
	}
	if len(stream1) != 2 {
		return "", nil, fmt.Errorf("invalid input. Stream has to be a two-tuple, got %d elements", len(stream1))
	}
	data, ok := stream1[1].([]interface{})
	if !ok {
		return "", nil, fmt.Errorf("invalid input. Stream data has to be a list, got %T", stream1[1])
	}

	if len(data) != 1 {
		return "", nil, fmt.Errorf("invalid input. Expected got %d stream data, expected 1", len(data))
	}

	v := data[0]
	element, ok := v.([]interface{})
	if !ok {
		return "", nil, fmt.Errorf("invalid input. Stream element has to be a two-tuple, got %T", v)
	}
	if len(element) != 2 {
		return "", nil, fmt.Errorf("invalid input. Stream element has to be a two-tuple, got %d elements", len(element))
	}
	id, ok := element[0].([]byte)
	if !ok {
		return "", nil, fmt.Errorf("invalid input. Stream ID has to be a string, got %T", element[0])
	}
	kv, ok := element[1].([]interface{})
	if !ok {
		return "", nil, fmt.Errorf("invalid input. Key values has to be a list of strings, got %T", element[1])
	}
	if len(kv)%2 != 0 {
		return "", nil, fmt.Errorf("invalid input. Odd number of key value pairs")
	}

	for i := 0; i < len(kv)-1; i += 2 {
		key, ok := kv[i].([]byte)
		if !ok {
			return "", nil, fmt.Errorf("invalid input. Key has to be a string, got %T", kv[i])
		}
		value, ok := kv[i+1].([]byte)
		if !ok {
			return "", nil, fmt.Errorf("invalid input. Values has to be a []byte, got %T", kv[i+1])
		}
		switch string(key) {
		case "content":
			return string(id), value, nil
		default:
			return "", nil, fmt.Errorf("invalid input. Unknown key \"%s\"", key)
		}
	}
	return "", nil, fmt.Errorf("invalid input. `content` not in response")
}
