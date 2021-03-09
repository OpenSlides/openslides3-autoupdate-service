package agenda

import (
	"encoding/json"
	"errors"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

type agendaItem struct {
	ID               int                               `json:"id"`
	ParentID         int                               `json:"parent_id"`
	Weight           int                               `json:"weight"`
	Type             int                               `json:"type"`
	TitleInformation map[string]*projector.OptionalStr `json:"title_information"`
	ItemNumber       string                            `json:"item_number"`
	ContentObject    struct {
		Collection string `json:"collection"`
	} `json:"content_object"`
}

type listOfSpeakers struct {
	TitleInformation map[string]*projector.OptionalStr `json:"title_information"`
	ContentObject    struct {
		ID         int    `json:"id"`
		Collection string `json:"collection"`
	} `json:"content_object"`
	Speakers []struct {
		UserID       int                    `json:"user_id"`
		Marked       json.RawMessage        `json:"marked"`
		PointOfOrder bool                   `json:"point_of_order"`
		Weight       projector.OptionalInt `json:"weight"`
		BeginTime    json.RawMessage        `json:"begin_time"`
		EndTime      json.RawMessage        `json:"end_time"`
	} `json:"speakers"`
	Closed json.RawMessage `json:"closed"`
}

var errNoListOfSpeakers = errors.New("list of speakers does not exist")
