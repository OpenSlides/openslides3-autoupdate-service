package motion

import (
	"encoding/json"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

type motion struct {
	ID                      int                    `json:"id"`
	Title                   json.RawMessage        `json:"title"`
	Text                    json.RawMessage        `json:"text"`
	Reason                  json.RawMessage        `json:"reason"`
	Identifier              json.RawMessage        `json:"identifier"`
	ParentID                json.RawMessage        `json:"parent_id"`
	StatuteParagraphID      json.RawMessage        `json:"statute_paragraph_id"`
	ChangeRecommendationIDs []int                  `json:"change_recommendations_id"`
	AmendmentIDs            []int                  `json:"amendments_id"`
	AmendmentParagraphs     json.RawMessage        `json:"amendment_paragraphs"`
	ModifiedFinalVersion    json.RawMessage        `json:"modified_final_version"`
	RecommendationID        *projector.OptionalInt `json:"recommendation_id"`
	RecommendationExtension json.RawMessage        `json:"recommendation_extension"`
	Submitters              []struct {
		UserID int `json:"user_id"`
		Weight int `json:"weight"`
	} `json:"submitters"`
}

type amendment struct {
	Title            json.RawMessage        `json:"title"`
	Identifier       json.RawMessage        `json:"identifier"`
	Paragraphs       json.RawMessage        `json:"amendment_paragraphs"`
	StateID          *projector.OptionalInt `json:"state_id"`
	RecommendationID *projector.OptionalInt `json:"recommendation_id"`
}
