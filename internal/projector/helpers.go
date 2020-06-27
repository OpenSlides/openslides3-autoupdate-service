package projector

import (
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
