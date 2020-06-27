package assignment

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/user"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

// Slide renders a an assignment.
func Slide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		a, err := projector.ModelFromElement(ds, e, "assignments/assignment")
		if err != nil {
			return nil, fmt.Errorf("getting assignments/assignment: %w", err)
		}

		var assignment struct {
			Title          json.RawMessage `json:"title"`
			Phase          json.RawMessage `json:"phase"`
			OpenPosts      json.RawMessage `json:"open_posts"`
			Description    json.RawMessage `json:"description"`
			PullCandidates json.RawMessage `json:"number_poll_candidates"`
			RelatedUsers   []struct {
				ID     int `json:"user_id"`
				Weight int `json:"weight"`
			} `json:"assignment_related_users"`
		}
		if err := json.Unmarshal(a, &assignment); err != nil {
			return nil, fmt.Errorf("decoding assignment: %w", err)
		}

		sort.Slice(assignment.RelatedUsers, func(i, j int) bool {
			return assignment.RelatedUsers[i].Weight < assignment.RelatedUsers[j].Weight
		})

		ru := make([]json.RawMessage, 0)
		for _, u := range assignment.RelatedUsers {
			username, err := user.GetUserName(ds, u.ID)
			if err != nil {
				return nil, fmt.Errorf("getting username: %w", err)
			}

			ru = append(ru, []byte(fmt.Sprintf(`{"user":"%s"}`, username)))
		}
		dru, err := json.Marshal(ru)
		if err != nil {
			return nil, fmt.Errorf("encoding related users: %w", err)
		}

		out := map[string]json.RawMessage{
			"title":                    assignment.Title,
			"phase":                    assignment.Phase,
			"open_posts":               assignment.OpenPosts,
			"desciption":               assignment.Description,
			"assignment_related_users": dru,
			"number_poll_candidates":   assignment.PullCandidates,
		}
		b, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("encoding outgoing data: %w", err)
		}
		return b, nil
	}
}
