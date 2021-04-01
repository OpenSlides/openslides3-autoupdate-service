package poll

import (
	"encoding/json"
	"fmt"
)

type RequiredUsersFunc func(json.RawMessage) (map[int]bool, string, error)

// RequiredPoll returns the users which voted for this poll if it is not anonymized.
func RequiredPoll(canSee string) RequiredUsersFunc {
	return func(data json.RawMessage) (map[int]bool, string, error) {
		var poll struct {
			VotedID            []int `json:"voted_id"`
			State              int   `json:"state"`
			IsPseudoanonymized bool  `json:"is_pseudoanonymized"`
		}
		if err := json.Unmarshal(data, &poll); err != nil {
			return nil, "", fmt.Errorf("unmarshal assignment poll option: %w", err)
		}

		// Pseudoanonymous and unpublished polls do not need any required users.
		if poll.State != 4 || poll.IsPseudoanonymized {
			return nil, canSee, nil
		}

		uids := make(map[int]bool, len(poll.VotedID))
		for _, id := range poll.VotedID {
			uids[id] = true
		}

		return uids, canSee, nil
	}
}
