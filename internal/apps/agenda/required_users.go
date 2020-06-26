package agenda

import (
	"encoding/json"
	"fmt"
)

// RequiredSpeakers returns the user ids of a list of speaker objekt.
func RequiredSpeakers(data json.RawMessage) ([]int, string, error) {
	var los struct {
		Speakers []struct {
			UserID int `json:"user_id"`
		} `json:"speakers"`
	}
	if err := json.Unmarshal(data, &los); err != nil {
		return nil, "", fmt.Errorf("unmarshal list of speaker: %w", err)
	}

	uids := make([]int, len(los.Speakers))
	for i, s := range los.Speakers {
		uids[i] = s.UserID
	}
	return uids, pCanSeeInternal, nil
}
