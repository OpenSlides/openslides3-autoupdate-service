package datastore

import (
	"encoding/json"
	"testing"
)

func TestRequiredUser(t *testing.T) {
	callables := map[string]func(json.RawMessage) (map[int]bool, string, error){
		"collection": func(json.RawMessage) (map[int]bool, string, error) {
			return map[int]bool{1: true}, "can.perm", nil
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

	if got := rq.UserRequired(1); len(got) != 1 || got[0] != "can.perm" {
		t.Errorf("UserRequired(1) returned %v, expected [can.perm]", got)
	}

	if got := rq.UserRequired(2); len(got) != 0 {
		t.Errorf("UserRequired(2) returned %v, expected []", got)
	}
}

func TestRequiredUserSecondUpdate(t *testing.T) {
	callables := map[string]func(json.RawMessage) (map[int]bool, string, error){
		"collection": func(json.RawMessage) (map[int]bool, string, error) {
			return map[int]bool{1: true}, "can.perm", nil
		},
	}

	data1 := map[string]json.RawMessage{
		"collection:1": []byte(`{}`),
	}

	data2 := map[string]json.RawMessage{
		"other:1": []byte(`{}`),
	}

	rq := requiredUser{callables: callables}

	if err := rq.update(data1); err != nil {
		t.Fatalf("updated failed: %v", err)
	}

	if err := rq.update(data2); err != nil {
		t.Fatalf("updated failed: %v", err)
	}

	if got := rq.UserRequired(1); len(got) != 1 || got[0] != "can.perm" {
		t.Errorf("UserRequired(1) returned %v, expected [can.perm]", got)
	}

	if got := rq.UserRequired(2); len(got) != 0 {
		t.Errorf("UserRequired(2) returned %v, expected []", got)
	}
}

func TestRequiredUserDelete(t *testing.T) {
	callables := map[string]func(json.RawMessage) (map[int]bool, string, error){
		"collection": func(json.RawMessage) (map[int]bool, string, error) {
			return map[int]bool{1: true}, "can.perm", nil
		},
	}

	data1 := map[string]json.RawMessage{
		"collection:1": []byte(`{}`),
	}

	data2 := map[string]json.RawMessage{
		"collection:1": nil,
	}

	rq := requiredUser{callables: callables}

	if err := rq.update(data1); err != nil {
		t.Fatalf("updated failed: %v", err)
	}

	if err := rq.update(data2); err != nil {
		t.Fatalf("updated failed: %v", err)
	}

	if got := rq.UserRequired(1); len(got) != 0 {
		t.Errorf("UserRequired(1) returned %v, expected []", got)
	}
}
