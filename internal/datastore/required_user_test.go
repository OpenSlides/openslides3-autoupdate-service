package datastore

import (
	"encoding/json"
	"testing"
)

func TestRequiredUser(t *testing.T) {
	callables := map[string]func(json.RawMessage) ([]int, error){
		"collection": func(json.RawMessage) ([]int, error) {
			return []int{1}, nil
		},
	}

	data := map[string]json.RawMessage{
		"collection:1": []byte(`{}`),
		"other:1":      []byte(`{}`),
	}

	rq := requiredUser{callables: callables}

	if err := rq.update(data); err != nil {
		t.Fatalf("updated failed: %v", err)
	}

	if !rq.UserRequired(1) {
		t.Errorf("UserRequired(1) returned false, expected true")
	}

	if rq.UserRequired(2) {
		t.Errorf("UserRequired(2) returned true, expected false")
	}
}
