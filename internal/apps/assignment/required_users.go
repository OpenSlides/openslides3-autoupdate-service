package assignment

import (
	"encoding/json"
	"fmt"
)

// RequiredPollOption returns the VoteID of the option.
func RequiredPollOption(data json.RawMessage) ([]int, string, error) {
	var option struct {
		UserID int `json:"user_id"`
	}
	if err := json.Unmarshal(data, &option); err != nil {
		return nil, "", fmt.Errorf("unmarshal assignment poll option: %w", err)
	}

	return []int{option.UserID}, CanSee, nil
}

// RequiredPoll returns the VoteID of the option.
func RequiredPoll(data json.RawMessage) ([]int, string, error) {
	var option struct {
		VoteID []int `json:"voted_id"`
	}
	if err := json.Unmarshal(data, &option); err != nil {
		return nil, "", fmt.Errorf("unmarshal assignment poll option: %w", err)
	}

	return option.VoteID, CanSee, nil
}
