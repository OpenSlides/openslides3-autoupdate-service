package assignment

import (
	"encoding/json"
	"fmt"
)

const (
	// CanSee is the can see permission string of assignments.
	CanSee = "assignments.can_see"

	// CanManage is the manage permission string of assignments.
	CanManage = "assignments.can_manage"
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

	uids := make([]int, len(assignment.Users))
	for i, u := range assignment.Users {
		uids[i] = u.ID
	}

	return uids, CanSee, nil
}

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
