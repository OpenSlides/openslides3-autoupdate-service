package agenda

import (
	"encoding/json"
	"errors"
)

type agendaItem struct {
	ID               int               `json:"id"`
	ParentID         int               `json:"parent_id"`
	Weight           int               `json:"weight"`
	Type             int               `json:"type"`
	TitleInformation map[string]string `json:"title_information"`
	ItemNumber       string            `json:"item_number"`
	ContentObject    struct {
		Collection string `json:"collection"`
	} `json:"content_object"`
}

type listOfSpeakers struct {
	TitleInformation map[string]string `json:"title_information"`
	ContentObject    struct {
		ID         int    `json:"id"`
		Collection string `json:"collection"`
	} `json:"content_object"`
	Speakers []struct {
		UserID    int             `json:"user_id"`
		Marked    json.RawMessage `json:"marked"`
		Weight    int             `json:"weight"`
		BeginTime json.RawMessage `json:"begin_time"`
		EndTime   json.RawMessage `json:"end_time"`
	} `json:"speakers"`
	Closed json.RawMessage `json:"closed"`
}

var errNoListOfSpeakers = errors.New("list of speakers does not exist")
