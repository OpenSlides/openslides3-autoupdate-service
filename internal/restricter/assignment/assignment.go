package assignment

import (
	"encoding/json"
	"fmt"
)

const (
	pCanSee = "assignment.can_see"
)

// RequiredAssignments returns the user ids of a assignment element.
func RequiredAssignments(data json.RawMessage) ([]int, string, error) {
	var assignment struct {
		Users []struct {
			ID int `json:"user_id"`
		} `json:"assignment_related_users"`
	}
	if err := json.Unmarshal(data, &assignment); err != nil {
		return nil, "", fmt.Errorf("unmarshal assignment: %w", err)
	}

	var uids []int
	for _, u := range assignment.Users {
		uids = append(uids, u.ID)
	}

	return uids, pCanSee, nil
}

// RequiredPollOption returns the VoteID of the option.
func RequiredPollOption(data json.RawMessage) ([]int, string, error) {
	var option struct {
		UserID int `json:"user_id"`
	}
	if err := json.Unmarshal(data, &option); err != nil {
		return nil, "", fmt.Errorf("unmarshal assignment poll option: %w", err)
	}

	return []int{option.UserID}, pCanSee, nil

}

// RequiredPoll returns the VoteID of the option.
func RequiredPoll(data json.RawMessage) ([]int, string, error) {
	var option struct {
		VoteID int `json:"voted_id"`
	}
	if err := json.Unmarshal(data, &option); err != nil {
		return nil, "", fmt.Errorf("unmarshal assignment poll option: %w", err)
	}

	return []int{option.VoteID}, pCanSee, nil

}
