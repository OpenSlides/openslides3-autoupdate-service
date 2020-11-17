package motion

import (
	"encoding/json"
	"fmt"
)

// RequiredMotions returns the user ids of a motion element.
func RequiredMotions(data json.RawMessage) (map[int]bool, string, error) {
	var motion struct {
		Submitters []struct {
			ID int `json:"user_id"`
		} `json:"submitters"`
		Supporters []int `json:"supporters_id"`
	}
	if err := json.Unmarshal(data, &motion); err != nil {
		return nil, "", fmt.Errorf("unmarshal motion: %w", err)
	}

	uids := map[int]bool{}
	for _, s := range motion.Submitters {
		uids[s.ID] = true
	}
	for _, id := range motion.Supporters {
		uids[id] = true
	}

	return uids, CanSee, nil
}

// RequiredPollOption returns the VoteID of the option.
func RequiredPollOption(data json.RawMessage) (map[int]bool, string, error) {
	var option struct {
		VoteID int `json:"voted_id"`
	}
	if err := json.Unmarshal(data, &option); err != nil {
		return nil, "", fmt.Errorf("unmarshal motion poll option: %w", err)
	}

	return map[int]bool{option.VoteID: true}, CanSee, nil
}
