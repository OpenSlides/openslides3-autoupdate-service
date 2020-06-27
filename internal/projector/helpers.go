package projector

import (
	"encoding/json"
	"fmt"
)

// ModelFromElement returns a model from the element.id.
func ModelFromElement(ds Datastore, e json.RawMessage, collection string) (json.RawMessage, error) {
	var element struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal(e, &element); err != nil {
		return nil, fmt.Errorf("decoding element: %w", err)
	}

	m := ds.Get(fmt.Sprintf("%s:%d", collection, element.ID))
	if m == nil {
		return nil, fmt.Errorf("model %s:%d does not exist", collection, element.ID)
	}
	return m, nil
}
