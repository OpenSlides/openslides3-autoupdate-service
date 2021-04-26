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
		var element struct {
			ID interface{} `json:"id"`
		}

		// Fallback for better error messages
		if err := json.Unmarshal(e, &element); err == nil {
			return NewClientError("%s with id %s does not exist", collection, element.ID)
		}
		return fmt.Errorf("decoding element: %w", err)
	}

	if element.ID == 0 {
		return NewClientError("id is required for %s slide", collection)
	}

	if err := ds.Get(collection, element.ID, v); err != nil {
		return NewClientError("%s with id %d does not exist", collection, element.ID)
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
	if o == nil {
		return 0
	}
	return o.value
}

// Null returns true, if, the value does not exist.
func (o *OptionalInt) Null() bool {
	if o == nil {
		return true
	}
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

// OptionalStr is a type that can be null or an string.
type OptionalStr struct {
	value string
	exist bool
}

// NewOptionalStr returns an OptionalStr from a golang string.
func NewOptionalStr(s string) *OptionalStr {
	return &OptionalStr{value: s, exist: true}
}

// Value returns the value of the type. Returns "" if it does not exist.
func (o *OptionalStr) Value() string {
	if o == nil {
		return ""
	}
	return o.value
}

// Null returns true, if, the value does not exist.
func (o *OptionalStr) Null() bool {
	if o == nil {
		return true
	}
	return !o.exist
}

// UnmarshalJSON builds this type from json.
func (o *OptionalStr) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte(`null`)) {
		o.exist = false
		return nil
	}

	o.exist = true
	return json.Unmarshal(b, &o.value)
}

// MarshalJSON decodes the type to json.
func (o *OptionalStr) MarshalJSON() ([]byte, error) {
	if o.Null() {
		return []byte(`null`), nil
	}
	return json.Marshal(o.value)
}

// OptionalBool is a type that can be null or an bool.
type OptionalBool struct {
	value bool
	exist bool
}

// Value returns the value of the type. Returns false if it does not exist.
func (o *OptionalBool) Value() bool {
	if o == nil {
		return false
	}
	return o.value
}

// Null returns true, if, the value does not exist.
func (o *OptionalBool) Null() bool {
	if o == nil {
		return true
	}
	return !o.exist
}

// UnmarshalJSON builds this type from json.
func (o *OptionalBool) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte(`null`)) {
		o.exist = false
		return nil
	}

	o.exist = true
	return json.Unmarshal(b, &o.value)
}

// MarshalJSON decodes the type to json.
func (o *OptionalBool) MarshalJSON() ([]byte, error) {
	if o.Null() {
		return []byte(`null`), nil
	}
	return json.Marshal(o.value)
}
