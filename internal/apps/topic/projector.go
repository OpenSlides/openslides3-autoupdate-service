package topic

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

// Slide renders "topics/topic".
func Slide(ds projector.Datastore) projector.CallableFunc {
	return func(e json.RawMessage, pid int) (json.RawMessage, error) {
		var element struct {
			ID int `json:"id"`
		}
		if err := json.Unmarshal(e, &element); err != nil {
			return nil, fmt.Errorf("decoding element")
		}

		rTopic := ds.Get("topics/topic:" + strconv.Itoa(element.ID))
		if rTopic == nil {
			return nil, fmt.Errorf("topics/topic:%d does not exist", element.ID)
		}

		var topic struct {
			Title        string `json:"title"`
			Text         string `json:"text"`
			AgendaItemID int    `json:"agenda_item_id"`
		}
		if err := json.Unmarshal(rTopic, &topic); err != nil {
			return nil, fmt.Errorf("decoding topic")
		}

		rItem := ds.Get("agenda/item:" + strconv.Itoa(topic.AgendaItemID))
		if rItem == nil {
			return nil, fmt.Errorf("agenda/item:%d does not exist", topic.AgendaItemID)
		}

		var item struct {
			Number string `json:"item_number"`
		}
		if err := json.Unmarshal(rItem, &item); err != nil {
			return nil, fmt.Errorf("decoding item: %w", err)
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
