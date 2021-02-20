package vote

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/auth"
)

// Datastore is the connection to the datastore package.
type Datastore interface {
	ConfigValue(string, interface{}) error
	RegisterUpdate(name string, f func(map[string]json.RawMessage) error)
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
func New(ds Datastore, backend Backend) *Vote {
	return &Vote{
		ds:      ds,
		backend: backend,
	}
}

// Assignment adds a vote to an assignment-poll.
func (v *Vote) Assignment(ctx context.Context, r io.Reader) error {
	return errors.New("TODO")
}

// Motion adds a vote to an motion-poll.
func (v *Vote) Motion(ctx context.Context, pid int, r io.Reader) error {
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
		Data       string `json:"data"`
		VoteUserID int    `json:"user_id"`
	}
	if err := json.NewDecoder(r).Decode(&payload); err != nil {
		return invalidInput("Body has invalid json.")
	}

	if payload.VoteUserID == 0 {
		payload.VoteUserID = requestUserID
	}

	voteUser, ok := v.users[payload.VoteUserID]
	if !ok {
		return invalidInput("Unknown vote user id: %d", payload.VoteUserID)
	}

	// TODO: Check voteUser in poll-group(?)

	d := payload.Data
	if d != "Y" && d != "N" && (d != "A" || poll.method == pollMethodYNA) {
		// Valid that given data matches poll method.
		return invalidInput("Data does not fit the poll method.")
	}

	if voteUser != requestUser && voteUser.delegetedTo != requestUserID {
		// Check request user is allowed to vote for vote-user.
		return invalidInput("You are not allowed to vote for user %d", payload.VoteUserID)
	}

	if voteUser == requestUser && voteUser.delegetedTo != 0 {
		return invalidInput("Your vote is delegated to someone else.")
	}

	// TODO: This has to be checked in a lua script at the same time then
	// writing the data!!!
	if voted, err := v.backend.HasVoted("motion", pid, payload.VoteUserID); voted || err != nil {
		if err != nil {
			return fmt.Errorf("ask for user has voted: %w", err)
		}
		return invalidInput("Your vote was already counted.")
	}

	// Valid vote. Save it.
	voteData := struct {
		RequestUser int    `json:"request_user_id,omitempty"`
		VoteUser    int    `json:"vote_user_id,omitempty"`
		Value       string `json:"value"`
		Weight      int    `json:"weight"`
	}{
		requestUserID,
		payload.VoteUserID,
		payload.Data,
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
