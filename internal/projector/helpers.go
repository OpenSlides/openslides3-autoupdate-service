package projector

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ModelFromElement returns a model from the element.id.
func ModelFromElement(ds Datastore, e json.RawMessage, collection string, v interface{}) error {
	var element struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal(e, &element); err != nil {
		return fmt.Errorf("decoding element: %w", err)
	}

	if err := ds.Get(collection, element.ID, v); err != nil {
		return fmt.Errorf("get model: %w", err)
	}
	return nil
}

// OptionalInt is a type that can be null or an int.
type OptionalInt struct {
	value int
	exist bool
}

// Value returns the value of the type. Returns 0 if it does not exist.
func (o *OptionalInt) Value() int {
	return o.value
}

// Null returns true, if, the value does not exist.
func (o *OptionalInt) Null() bool {
	return !o.exist
}

// UnmarshalJSON builds this type from json.
func (o *OptionalInt) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte(`null`)) {
		o.exist = false
		return nil
	}

	o.exist = true
	return json.Unmarshal(b, &o.value)
}

// MarshalJSON decodes the type to json.
func (o *OptionalInt) MarshalJSON() ([]byte, error) {
	if o.Null() {
		return []byte(`null`), nil
	}
	return json.Marshal(o.value)
}
