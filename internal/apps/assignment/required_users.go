package assignment

import (
	"encoding/json"
	"fmt"
)

// RequiredAssignments returns the user ids of a assignment element.
func RequiredAssignments(data json.RawMessage) (map[int]bool, string, error) {
	var assignment struct {
		Users []struct {
			ID int `json:"user_id"`
		} `json:"assignment_related_users"`
	}
	if err := json.Unmarshal(data, &assignment); err != nil {
		return nil, "", fmt.Errorf("unmarshal assignment: %w", err)
	}

	uids := make(map[int]bool, len(assignment.Users))
	for _, u := range assignment.Users {
		uids[u.ID] = true
	}

	return uids, CanSee, nil
}

// RequiredPollOption returns the VoteID of the option.
func RequiredPollOption(data json.RawMessage) (map[int]bool, string, error) {
	var option struct {
		UserID int `json:"user_id"`
	}
	if err := json.Unmarshal(data, &option); err != nil {
		return nil, "", fmt.Errorf("unmarshal assignment poll option: %w", err)
	}

	return map[int]bool{option.UserID: true}, CanSee, nil
}
