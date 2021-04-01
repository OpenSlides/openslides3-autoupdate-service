package poll

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

// StateFinished and StatePublished are state ids of polls.
const (
	StateCreated   = 1
	StateStarted   = 2
	StateFinished  = 3
	StatePublished = 4
)

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

		var votedIDs []int
		if err := json.Unmarshal(poll["voted_id"], &votedIDs); err != nil {
			return nil, fmt.Errorf("unmarshal voted_id: %w", err)
		}

		if uid != 0 {
			// The following code adds the fields `user_has_voted` and
			// `user_has_voted_for_delegations` to the poll. This can be skipped
			// for anonymous users.
			hasVoted := []byte("false")
			for _, id := range votedIDs {
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
				for _, votedID := range votedIDs {
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

		// Remove values that should not be published to anyone as long as the
		// vote is open.
		if state != StatePublished && state != StateFinished {
			poll["votescast"] = []byte(strconv.Itoa(len(votedIDs)))
			delete(poll, "voted_id")
			delete(poll, "votesvalid")
			delete(poll, "votesinvalid")
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

		// Make sure that the user ids are sorted.
		if _, ok := poll["voted_id"]; ok {
			sort.Ints(votedIDs)
			bs, err := json.Marshal(votedIDs)
			if err != nil {
				return nil, fmt.Errorf("marshal voted_id: %w", err)
			}
			poll["voted_id"] = bs
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
//
// 1. You need the poll-can-see permission to get any vote.
// 2. If the poll state is published, everyone can see the full vote.
// 3. If the poll state is finished, managers can see the "user_token".
// 4. Otherwise, no one can see the field "user_token" and only managers, the
//    vote user and the delegated_vote_user can see the vote.
func RestrictVote(r restricter.HasPermer, canSee, canManage string) restricter.ElementFunc {
	return func(uid int, element json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, canSee) {
			return nil, nil
		}

		var vote map[string]json.RawMessage
		if err := json.Unmarshal(element, &vote); err != nil {
			return nil, fmt.Errorf("unmarshal vote: %w", err)
		}

		var state int
		if err := json.Unmarshal(vote["pollstate"], &state); err != nil {
			return nil, fmt.Errorf("unmarshal pollstate: %w", err)
		}

		if state == StatePublished {
			return element, nil
		}

		// The user_token is only visible in the published state, and finished state
		// for managers
		if !(r.HasPerm(uid, canManage) && state == StateFinished) {
			delete(vote, "user_token")
		}

		data, err := json.Marshal(vote)
		if err != nil {
			return nil, fmt.Errorf("marshal vote: %w", err)
		}

		if r.HasPerm(uid, canManage) {
			return data, nil
		}

		var userID int
		if err := json.Unmarshal(vote["user_id"], &userID); err != nil {
			return nil, fmt.Errorf("unmarshal user_id: %w", err)
		}

		if userID == uid {
			return data, nil
		}

		var delegatedUserID int
		if err := json.Unmarshal(vote["delegated_user_id"], &delegatedUserID); err != nil {
			return nil, fmt.Errorf("unmarshal delegated_user_id: %w", err)
		}

		if delegatedUserID == uid {
			return data, nil
		}

		return nil, nil
	}
}
