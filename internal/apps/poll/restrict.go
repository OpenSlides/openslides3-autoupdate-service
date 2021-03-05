package poll

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

// StatePublished is the state of a poll, when it is published.
const StatePublished = 4

// RestrictPoll restricts an element for an assignment or motion poll.
func RestrictPoll(r restricter.HasPermer, canSee, canManage string, restrictedFiels []string) restricter.ElementFunc {
	return func(uid int, element json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, canSee) {
			return nil, nil
		}

		var poll map[string]json.RawMessage
		if err := json.Unmarshal(element, &poll); err != nil {
			return nil, fmt.Errorf("unmarshal poll: %w", err)
		}

		if uid != 0 {
			// The following code adds the fields `user_has_voted` and
			// `user_has_voted_for_delegations` to the poll. This can be skipped
			// for anonymous users.
			var votedID []int
			if err := json.Unmarshal(poll["voted_id"], &votedID); err != nil {
				return nil, fmt.Errorf("unmarshal voted_id: %w", err)
			}

			hasVoted := []byte("false")
			for _, id := range votedID {
				if id == uid {
					hasVoted = []byte("true")
					break
				}
			}
			poll["user_has_voted"] = hasVoted

			// Get the users `vote_delegated_from_users_id`.
			var user struct {
				VoteDelegationIds []int `json:"vote_delegated_from_users_id"`
			}
			if err := r.Get("users/user", uid, &user); err != nil {
				return nil, fmt.Errorf("unmarshal user: %w", err)
			}
			// Calc the intersection of voteDelegationIds and votedID.
			ids := []int{}
			for _, delegationID := range user.VoteDelegationIds {
				for _, votedID := range votedID {
					if delegationID == votedID {
						ids = append(ids, delegationID)
						break
					}
				}
			}
			m, err := json.Marshal(ids)
			if err != nil {
				return nil, fmt.Errorf("marshal user_has_voted_for_delegations: %w", err)
			}
			poll["user_has_voted_for_delegations"] = m
		}

		var state int
		if err := json.Unmarshal(poll["state"], &state); err != nil {
			return nil, fmt.Errorf("unmarshal poll state: %w", err)
		}

		// Delete some fields for no managers and unpublished polls.
		if !(r.HasPerm(uid, canManage) || state == StatePublished) {
			delete(poll, "votesvalid")
			delete(poll, "votesinvalid")
			delete(poll, "votescast")
			delete(poll, "voted_id")
			for _, field := range restrictedFiels {
				delete(poll, field)
			}
		}

		data, err := json.Marshal(poll)
		if err != nil {
			return nil, fmt.Errorf("marshal poll: %w", err)
		}
		return data, nil
	}
}

// RestrictOption restricts an element for a poll option.
func RestrictOption(r restricter.HasPermer, canSee, canManage string) restricter.ElementFunc {
	return func(uid int, element json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, canSee) {
			return nil, nil
		}

		if r.HasPerm(uid, canManage) {
			return element, nil
		}

		var option map[string]json.RawMessage
		if err := json.Unmarshal(element, &option); err != nil {
			return nil, fmt.Errorf("unmarshal poll: %w", err)
		}

		var state int
		if err := json.Unmarshal(option["pollstate"], &state); err != nil {
			return nil, fmt.Errorf("unmarshal pollstate: %w", err)
		}

		// Delete some fields for unpublished options.
		if state != StatePublished {
			delete(option, "yes")
			delete(option, "no")
			delete(option, "abstain")
		}

		data, err := json.Marshal(option)
		if err != nil {
			return nil, fmt.Errorf("marshal option: %w", err)
		}
		return data, nil
	}
}

// RestrictVote restricts an element for a poll vote.
func RestrictVote(r restricter.HasPermer, canSee, canManage string) restricter.ElementFunc {
	return func(uid int, element json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, canSee) {
			return nil, nil
		}

		if r.HasPerm(uid, canManage) {
			return element, nil
		}

		var vote map[string]json.RawMessage
		if err := json.Unmarshal(element, &vote); err != nil {
			return nil, fmt.Errorf("unmarshal vote: %w", err)
		}

		var userID int
		if err := json.Unmarshal(vote["user_id"], &userID); err != nil {
			return nil, fmt.Errorf("unmarshal user_id: %w", err)
		}

		if userID == uid {
			return element, nil
		}

		var delegatedUserID int
		if err := json.Unmarshal(vote["delegated_user_id"], &delegatedUserID); err != nil {
			return nil, fmt.Errorf("unmarshal delegated_user_id: %w", err)
		}

		if delegatedUserID == uid {
			return element, nil
		}

		var state int
		if err := json.Unmarshal(vote["pollstate"], &state); err != nil {
			return nil, fmt.Errorf("unmarshal pollstate: %w", err)
		}

		if state == StatePublished {
			return element, nil
		}

		return nil, nil
	}
}
