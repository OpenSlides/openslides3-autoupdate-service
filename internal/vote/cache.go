package vote

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type user struct {
	isPresent   bool
	delegetedTo int
	voteWeight  int
}

type poll struct {
	// On motion poll and assignment poll.
	anonymous bool
	method    int

	// Only on assignment poll.
	globalYes     bool
	globalNo      bool
	globalAbstain bool
	multipleVotes bool
	minAmount     int
	maxAmount     int
}

func (v *Vote) update(data map[string]json.RawMessage) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	if v.users == nil {
		v.users = make(map[int]*user)
		v.pollsMotion = make(map[int]*poll)
		v.pollsAssignment = make(map[int]*poll)
		v.optionToPoll = make(map[int]int)
	}

	for elementID, value := range data {
		parts := strings.Split(elementID, ":")
		if len(parts) != 2 {
			return fmt.Errorf("key %s has wrong format. Expected one `:`, got: ", elementID)
		}

		id, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("invalid element id `%s`, second part has to be an int not %s", elementID, parts[1])
		}

		switch parts[0] {
		case "users/user":
			if err := v.updateUser(id, value); err != nil {
				return fmt.Errorf("updating user %s: %w", parts[1], err)
			}
		case "motions/motion-poll":
			if err := v.updateMotionPoll(id, value); err != nil {
				return fmt.Errorf("updating motion-poll %s: %w", parts[1], err)
			}
		case "assignments/assignment-poll":
			if err := v.updateAssignmentPoll(id, value); err != nil {
				return fmt.Errorf("updating assignment-poll %s: %w", parts[1], err)
			}
		case "assignments/assignment-option":
			if err := v.updateAssignmntOption(id, value); err != nil {
				return fmt.Errorf("updating assignment-option %s: %w", parts[1], err)
			}
		}
	}
	return nil
}

func (v *Vote) updateUser(id int, value json.RawMessage) error {
	if value == nil {
		// Delete user.
		delete(v.users, id)
		return nil
	}

	var decoded struct {
		Present   bool   `json:"is_present"`
		Delegated int    `json:"vote_delegated_to_id"`
		Weight    string `json:"vote_weight"`
	}
	if err := json.Unmarshal(value, &decoded); err != nil {
		return fmt.Errorf("decoding user: %w", err)
	}

	weight, err := decodeDecimal(decoded.Weight)
	if err != nil {
		return fmt.Errorf("decoding decimal: %w", err)
	}

	v.users[id] = &user{
		isPresent:   decoded.Present,
		delegetedTo: decoded.Delegated,
		voteWeight:  weight,
	}
	return nil
}

func (v *Vote) updateMotionPoll(id int, value json.RawMessage) error {
	if value == nil {
		// Delete poll.
		delete(v.pollsMotion, id)
		return nil
	}

	var decoded struct {
		State  int    `json:"state"`
		Type   string `json:"type"`
		Method string `json:"pollmethod"`
	}
	if err := json.Unmarshal(value, &decoded); err != nil {
		return fmt.Errorf("decoding motion-poll: %w", err)
	}

	if decoded.State != 2 {
		// Only polls in state 2 (STATE_STARTED) are relevant. If they are in
		// the cache, delete them.
		delete(v.pollsMotion, id)
		return nil
	}

	var anonymous bool
	switch decoded.Type {
	case "analog":
		// Anonlog polls are ignored. Delete, if exist.
		delete(v.pollsMotion, id)
		return nil
	case "named":
	case "pseudoanonymous":
		anonymous = true
	default:
		return fmt.Errorf("invalid vote type in poll %d: %s", id, decoded.Type)
	}

	var method int
	switch decoded.Method {
	case "YN":
		method = pollMethodYN
	case "YNA":
		method = pollMethodYNA
	default:
		return fmt.Errorf("invalid poll method in poll %d: %s", id, decoded.Method)
	}

	v.pollsMotion[id] = &poll{
		anonymous: anonymous,
		method:    method,
	}

	return nil
}

func (v *Vote) updateAssignmentPoll(id int, value json.RawMessage) error {
	if value == nil {
		// Delete poll.
		delete(v.pollsAssignment, id)
		return nil
	}

	var decoded struct {
		State         int    `json:"state"`
		Type          string `json:"type"`
		Method        string `json:"pollmethod"`
		GlobalYes     bool   `json:"global_yes"`
		GlobalNo      bool   `json:"global_no"`
		GlobalAbstain bool   `json:"global_abstain"`
		Multi         bool   `json:"allow_multiple_votes_per_candidate"`
		Min           int    `json:"min_votes_amount"`
		Max           int    `json:"max_votes_amount"`
	}
	if err := json.Unmarshal(value, &decoded); err != nil {
		return fmt.Errorf("decoding motion-poll: %w", err)
	}

	if decoded.State != 2 {
		// Only polls in state 2 (STATE_STARTED) are relevant. If they are in
		// the cache, delete them.
		delete(v.pollsMotion, id)
		return nil
	}

	var anonymous bool
	switch decoded.Type {
	case "analog":
		// Anonlog polls are ignored. Delete, if exist.
		delete(v.pollsMotion, id)
		return nil
	case "named":
	case "pseudoanonymous":
		anonymous = true
	default:
		return fmt.Errorf("invalid vote type in poll %d: %s", id, decoded.Type)
	}

	var method int
	switch decoded.Method {
	case "YN":
		method = pollMethodYN
	case "YNA":
		method = pollMethodYNA
	case "Y":
		method = pollMethodY
	case "N":
		method = pollMethodN
	default:
		return fmt.Errorf("invalid poll method in poll %d: %s", id, decoded.Method)
	}

	v.pollsMotion[id] = &poll{
		anonymous:     anonymous,
		method:        method,
		globalYes:     decoded.GlobalYes,
		globalNo:      decoded.GlobalNo,
		globalAbstain: decoded.GlobalAbstain,
		multipleVotes: decoded.Multi,
		minAmount:     decoded.Min,
		maxAmount:     decoded.Max,
	}

	return nil
}

func (v *Vote) updateAssignmntOption(id int, value json.RawMessage) error {
	if value == nil {
		// Delete option from poll.
		delete(v.optionToPoll, id)
		return nil
	}

	var decoded struct {
		PollID int `json:"poll_id"`
	}
	if err := json.Unmarshal(value, &decoded); err != nil {
		return fmt.Errorf("decoding assignments/assignment-poll: %w", err)
	}

	// This code expects, that a option.poll_id can not change.
	v.optionToPoll[id] = decoded.PollID

	return nil
}

func decodeDecimal(n string) (int, error) {
	parts := strings.Split(n, ".")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid decimal value %s", n)
	}
	i1, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid decimal value %s", n)
	}
	i2, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid decimal value %s", n)
	}
	return i1*1_000_000 + i2, nil
}

const (
	pollMethodYN = iota
	pollMethodYNA
	pollMethodY
	pollMethodN
)
