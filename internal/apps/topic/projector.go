package topic

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

// Slide renders "topics/topic".
func Slide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		var topic struct {
			Title        string `json:"title"`
			Text         string `json:"text"`
			AgendaItemID int    `json:"agenda_item_id"`
		}

		if err := projector.ModelFromElement(ds, e, "topics/topic", &topic); err != nil {
			return nil, fmt.Errorf("getting topics/topic: %w", err)
		}

		var item struct {
			Number string `json:"item_number"`
		}
		if err := ds.Get("agenda/item", topic.AgendaItemID, &item); err != nil {
			return nil, fmt.Errorf("getting agenda/item: %w", err)
		}

		out := struct {
			Title      string `json:"title"`
			Text       string `json:"text"`
			ItemNumber string `json:"item_number"`
		}{
			topic.Title,
			topic.Text,
			item.Number,
		}
		data, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("decoding topic slide: %w", err)
		}
		return data, nil
	}
}
