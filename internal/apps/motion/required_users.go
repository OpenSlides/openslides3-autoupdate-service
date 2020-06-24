package motion

import (
	"encoding/json"
	"fmt"
)

// RequiredMotions returns the user ids of a motion element.
func RequiredMotions(data json.RawMessage) ([]int, string, error) {
	var motion struct {
		Submitters []struct {
			ID int `json:"user_id"`
		} `json:"submitters"`
		Supporters []int `json:"supporters_id"`
	}
	if err := json.Unmarshal(data, &motion); err != nil {
		return nil, "", fmt.Errorf("unmarshal motion: %w", err)
	}

	uidSet := make(map[int]bool)
	for _, s := range motion.Submitters {
		uidSet[s.ID] = true
	}
	for _, id := range motion.Supporters {
		uidSet[id] = true
	}

	uids := make([]int, 0, len(uidSet))
	for id := range uidSet {
		uids = append(uids, id)
	}

	return uids, CanSee, nil
}

// RequiredPollOption returns the VoteID of the option.
func RequiredPollOption(data json.RawMessage) ([]int, string, error) {
	var option struct {
		VoteID int `json:"voted_id"`
	}
	if err := json.Unmarshal(data, &option); err != nil {
		return nil, "", fmt.Errorf("unmarshal motion poll option: %w", err)
	}

	return []int{option.VoteID}, CanSee, nil
}
