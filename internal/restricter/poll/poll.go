package poll

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

const published = 4

// RestrictPoll restricts an element for an assignment or motion poll.
func RestrictPoll(r restricter.HasPermer, canSee, canManage string) restricter.ElementFunc {
	return func(uid int, element json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, canSee) {
			return nil, nil
		}

		var poll map[string]json.RawMessage
		if err := json.Unmarshal(element, &poll); err != nil {
			return nil, fmt.Errorf("unmarshal poll: %w", err)
		}

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

		var state int
		if err := json.Unmarshal(poll["state"], &state); err != nil {
			return nil, fmt.Errorf("unmarshal poll state: %w", err)
		}

		// delete some fields for no managers and unpublished polls.
		if !(r.HasPerm(uid, canManage) || state == published) {
			delete(poll, "votesvalid")
			delete(poll, "votesinvalid")
			delete(poll, "votescast")
			delete(poll, "voted_id")
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

		// delete some fields for unpublished options.
		if state != published {
			delete(option, "votesvalid")
			delete(option, "votesinvalid")
			delete(option, "votescast")
			delete(option, "voted_id")
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

		var state int
		if err := json.Unmarshal(vote["pollstate"], &state); err != nil {
			return nil, fmt.Errorf("unmarshal pollstate: %w", err)
		}

		if state == published {
			return element, nil
		}

		return nil, nil
	}
}
