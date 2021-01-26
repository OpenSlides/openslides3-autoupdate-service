package datastore

import (
	"encoding/json"
	"testing"
)

func TestApplause(t *testing.T) {
	a := &applause{c: new(config)}

	t.Run("create user", func(t *testing.T) {
		data := map[string]json.RawMessage{
			"users/user:1": []byte(`{"is_present":true}`),
		}
		if err := a.update(data); err != nil {
			t.Fatalf("Got unexpected error: %v", err)
		}

		if !a.presentUsers[1] {
			t.Error("User is not present")
		}
	})

	t.Run("delete user", func(t *testing.T) {
		data := map[string]json.RawMessage{
			"users/user:1": nil,
		}
		if err := a.update(data); err != nil {
			t.Fatalf("Got unexpected error: %v", err)
		}

		if a.presentUsers[1] {
			t.Error("User is present")
		}
	})

}
