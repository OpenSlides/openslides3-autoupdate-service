package vote

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/auth"
)

// Datastore is the connection to the datastore package.
type Datastore interface {
	ConfigValue(string, interface{}) error
	RegisterUpdate(name string, f func(map[string]json.RawMessage) error) error
	InGroups(uid int, groups []int) bool
}

// Backend to store the votes to.
type Backend interface {
	HasVoted(collection string, id int, userID int) (bool, error)
	Save(collection string, id int, voteUserID int, data []byte) error
}

// Vote handles the vote cache.
type Vote struct {
	mu sync.RWMutex

	ds      Datastore
	backend Backend

	users           map[int]*user
	pollsMotion     map[int]*poll
	pollsAssignment map[int]*poll

	optionToPoll map[int]int
}

// New creates a new Vote-Service.
func New(ds Datastore, backend Backend) (*Vote, error) {
	v := &Vote{
		ds:      ds,
		backend: backend,
	}
	if err := ds.RegisterUpdate("voteCache", v.update); err != nil {
		return nil, fmt.Errorf("register updater: %w", err)
	}
	return v, nil
}

func (v *Vote) assignmentDataValidate(pid int, d pollData) error {
	p := v.pollsAssignment[pid]

	if p.method == pollMethodY || p.method == pollMethodN {
		if d.Type() == pollDataString {
			if d.str == "Y" && p.globalYes || d.str == "N" && p.globalNo || d.str == "A" && p.globalAbstain {
				// Data is Y|N|A and global_X is true
				return nil
			}
			return invalidInput("Your answer was not allowed")
		}

		if d.Type() == pollDataOptionAmount {
			var sumAmount int
			for optionID, amount := range d.optionAmount {
				if amount < 0 {
					return invalidInput("Your answer for option %d has to be >= 0", optionID)
				}

				if !p.multipleVotes && amount > 1 {
					return invalidInput("Your answer for option %d has to be 0 or 1", optionID)
				}

				if pollID, ok := v.optionToPoll[optionID]; pollID != pid || !ok {
					return invalidInput("Option_id %d does not belong to the poll", optionID)
				}

				sumAmount += amount
			}

			if sumAmount < p.minAmount || sumAmount > p.maxAmount {
				return invalidInput("The sum of your answers have to be between %d and %d", p.minAmount, p.maxAmount)
			}
			return nil
		}

		return invalidInput("Invalid data")
	}

	// Pollmethod == YN | YNA

	if d.Type() != pollDataOptionString {
		return invalidInput("Invalid data")
	}

	for optionID, yna := range d.optionYNA {
		if pollID, ok := v.optionToPoll[optionID]; pollID != pid || !ok {
			return invalidInput("Option_id %d does not belong to the poll", optionID)
		}

		if yna != "Y" && yna != "N" && (yna != "A" || p.method == pollMethodYNA) {
			// Valid that given data matches poll method.
			return invalidInput("Data for option %d does not fit the poll method.", optionID)
		}
	}
	return nil
}

// Assignment adds a vote to an assignment-poll.
func (v *Vote) Assignment(ctx context.Context, pid int, r io.Reader) error {
	return v.save(ctx, pid, r, v.pollsAssignment, v.assignmentDataValidate)
}

// Motion adds a vote to an motion-poll.
func (v *Vote) Motion(ctx context.Context, pid int, r io.Reader) error {
	payloadValidate := func(pid int, d pollData) error {

		if d.Type() != pollDataString {
			return invalidInput("Data has to be a string")
		}

		if d.str != "Y" && d.str != "N" && (d.str != "A" || v.pollsMotion[pid].method == pollMethodYNA) {
			// Valid that given data matches poll method.
			return invalidInput("Data does not fit the poll method.")
		}
		return nil
	}

	return v.save(ctx, pid, r, v.pollsMotion, payloadValidate)
}

func (v *Vote) save(ctx context.Context, pid int, r io.Reader, polls map[int]*poll, payloadValidate func(pid int, d pollData) error) error {
	v.mu.RLock()
	defer v.mu.RUnlock()

	poll, ok := v.pollsMotion[pid]
	if !ok {
		// Poll is not in the cache if it does not exist, is analog or is in the wrong state.
		return invalidInput("Poll %d does not exist does not recieve votes.", pid)
	}

	requestUserID := auth.FromContext(ctx)
	if requestUserID == 0 {
		return fmt.Errorf("anonymous vote, this should be checked in the http package")
	}

	requestUser, ok := v.users[requestUserID]
	if !ok {
		return invalidInput("Unknown request user id: %d", requestUserID)
	}

	if !requestUser.isPresent {
		return invalidInput("You have to be present")
	}

	var payload struct {
		Data       pollData `json:"data"`
		VoteUserID int      `json:"user_id"`
	}
	if err := json.NewDecoder(r).Decode(&payload); err != nil {
		return invalidInput("Body has invalid json.")
	}

	if err := payloadValidate(pid, payload.Data); err != nil {
		return invalidInput("Invalid vote data")
	}

	if payload.VoteUserID == 0 {
		payload.VoteUserID = requestUserID
	}

	voteUser, ok := v.users[payload.VoteUserID]
	if !ok {
		return invalidInput("Unknown vote user id: %d", payload.VoteUserID)
	}

	if !v.ds.InGroups(payload.VoteUserID, poll.groups) {
		return invalidInput("The Poll is not for you")
	}

	if voteUser != requestUser && voteUser.delegetedTo != requestUserID {
		// Check request user is allowed to vote for vote-user.
		return invalidInput("You are not allowed to vote for user %d", payload.VoteUserID)
	}

	// TODO: This has to be checked in a lua script at the same time as
	// writing the data!!!
	if voted, err := v.backend.HasVoted("motion", pid, payload.VoteUserID); voted || err != nil {
		if err != nil {
			return fmt.Errorf("ask for user has voted: %w", err)
		}
		return invalidInput("Your vote was already counted.")
	}

	// Valid vote. Save it.
	voteData := struct {
		RequestUser int             `json:"request_user_id,omitempty"`
		VoteUser    int             `json:"vote_user_id,omitempty"`
		Value       json.RawMessage `json:"value"`
		Weight      int             `json:"weight"`
	}{
		requestUserID,
		payload.VoteUserID,
		payload.Data.original,
		voteUser.voteWeight,
	}

	if poll.anonymous {
		voteData.RequestUser = 0
		voteData.VoteUser = 0
	}

	var voteWeightConfig bool
	if err := v.ds.ConfigValue("users_activate_vote_weight", &voteWeightConfig); err != nil {
		return fmt.Errorf("getting vote weight value from config: %w", err)
	}

	if !voteWeightConfig {
		voteData.Weight = 1
	}

	bs, err := json.Marshal(voteData)
	if err != nil {
		return fmt.Errorf("encoding vote data: %w", err)
	}

	if err := v.backend.Save("motion", pid, payload.VoteUserID, bs); err != nil {
		return fmt.Errorf("saving vote: %w", err)
	}

	// TODO: inform backend

	return nil
}

type pollData struct {
	str          string
	optionAmount map[int]int
	optionYNA    map[int]string

	original json.RawMessage
}

func (p *pollData) UnmarshalJSON(b []byte) error {
	p.original = b

	if err := json.Unmarshal(b, &p.str); err == nil {
		// pollData is a string
		return nil
	}

	if err := json.Unmarshal(b, &p.optionAmount); err == nil {
		// pollData is option_id to amount
		return nil
	}

	if err := json.Unmarshal(b, &p.optionYNA); err == nil {
		// pollData  is option_id to string
		return nil
	}

	return fmt.Errorf("unknown poll data: %s", b)
}

const (
	pollDataUnknown = iota
	pollDataString
	pollDataOptionAmount
	pollDataOptionString
)

func (p *pollData) Type() int {
	if p.str != "" {
		return pollDataOptionString
	}

	if p.optionAmount != nil {
		return pollDataOptionAmount
	}

	if p.optionYNA != nil {
		return pollDataOptionString
	}

	return pollDataUnknown
}
