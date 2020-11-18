package agenda

import (
	"encoding/json"
	"fmt"
)

// RequiredSpeakers returns the user ids of a list of speaker objekt.
func RequiredSpeakers(data json.RawMessage) (map[int]bool, string, error) {
	var los struct {
		Speakers []struct {
			UserID int `json:"user_id"`
		} `json:"speakers"`
	}
	if err := json.Unmarshal(data, &los); err != nil {
		return nil, "", fmt.Errorf("unmarshal list of speaker: %w", err)
	}

	uids := make(map[int]bool, len(los.Speakers))
	for _, s := range los.Speakers {
		uids[s.UserID] = true
	}
	return uids, CanSeeListOfSpeakers, nil
}
