package assignment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/poll"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/user"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

// Slide renders a an assignment.
func Slide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
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
		if err := projector.ModelFromElement(ds, e, "assignments/assignment", &assignment); err != nil {
			return nil, fmt.Errorf("getting assignments/assignment: %w", err)
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

			ru = append(ru, []byte(fmt.Sprintf(`{"user":%s}`, username)))
		}
		dru, err := json.Marshal(ru)
		if err != nil {
			return nil, fmt.Errorf("encoding related users: %w", err)
		}

		out := map[string]json.RawMessage{
			"title":                    assignment.Title,
			"phase":                    assignment.Phase,
			"open_posts":               assignment.OpenPosts,
			"description":              assignment.Description,
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

// PollSlide renders a an assignment poll.
func PollSlide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		var pollInfo map[string]json.RawMessage
		if err := projector.ModelFromElement(ds, e, "assignments/assignment-poll", &pollInfo); err != nil {
			return nil, fmt.Errorf("getting assignments/assignment-poll: %w", err)
		}

		assignmentID, err := strconv.Atoi(string(pollInfo["assignment_id"]))
		if err != nil {
			return nil, fmt.Errorf("getting assignment id: %w", err)
		}

		var pollState int
		if err := json.Unmarshal(pollInfo["state"], &pollState); err != nil {
			return nil, fmt.Errorf("decoding poll state: %w", err)
		}

		var assignment struct {
			Title string `json:"title"`
		}
		if err := ds.Get("assignments/assignment", assignmentID, &assignment); err != nil {
			return nil, fmt.Errorf("getting assignment: %w", err)
		}

		pollData := make(map[string]json.RawMessage)
		keys := []string{
			"title",
			"type",
			"pollmethod",
			"min_votes_amount",
			"max_votes_amount",
			"description",
			"state",
			"onehundred_percent_base",
			"majority_method",
			"entitled_users_at_stop",
		}
		for _, key := range keys {
			pollData[key] = pollInfo[key]
		}

		var optionIDs []int
		if err := json.Unmarshal(pollInfo["options_id"], &optionIDs); err != nil {
			return nil, fmt.Errorf("decoding poll.options_id: %w", err)
		}

		os := ds.GetModels("assignments/assignment-option", optionIDs)

		optionData := make([]openDataType, len(os))
		for i, o := range os {
			var option struct {
				UserID  int    `json:"user_id"`
				Yes     string `json:"yes"`
				No      string `json:"no"`
				Abstain string `json:"abstain"`
				Weight  int    `json:"weight"`
			}
			if err := json.Unmarshal(o, &option); err != nil {
				return nil, fmt.Errorf("decoding assingment option: %w", err)
			}

			username, err := user.GetUserName(ds, option.UserID)
			if err != nil {
				return nil, fmt.Errorf("getting username: %w", err)
			}

			data := map[string]interface{}{
				"user": map[string]json.RawMessage{
					"short_name": username,
				},
			}
			if pollState == poll.StatePublished {
				yes, err := strconv.ParseFloat(option.Yes, 32)
				if err != nil {
					return nil, fmt.Errorf("decoding option.yes: %w", err)
				}
				data["yes"] = yes

				no, err := strconv.ParseFloat(option.No, 32)
				if err != nil {
					return nil, fmt.Errorf("decoding option.no: %w", err)
				}
				data["no"] = no

				abstain, err := strconv.ParseFloat(option.Abstain, 32)
				if err != nil {
					return nil, fmt.Errorf("decoding option.abstain: %w", err)
				}
				data["abstain"] = abstain
			}
			optionData[i].data = data
			optionData[i].weight = option.Weight
			optionData[i].userID = option.UserID
		}

		sort.Slice(optionData, func(i, j int) bool {
			if optionData[i].weight != optionData[j].weight {
				return optionData[i].weight < optionData[j].weight
			}
			return optionData[i].userID > optionData[j].userID
		})

		b, err := json.Marshal(optionData)
		if err != nil {
			return nil, fmt.Errorf("encoding option data: %w", err)
		}
		pollData["options"] = b

		if pollState == poll.StatePublished {
			pollData["amount_global_yes"] = []byte("null")
			if !bytes.Equal(pollInfo["amount_global_yes"], []byte("0")) {
				pollData["amount_global_yes"] = pollInfo["amount_global_yes"]
			}

			pollData["amount_global_no"] = []byte("null")
			if !bytes.Equal(pollInfo["amount_global_no"], []byte("0")) {
				pollData["amount_global_no"] = pollInfo["amount_global_no"]
			}

			pollData["amount_global_abstain"] = []byte("null")
			if !bytes.Equal(pollInfo["amount_global_abstain"], []byte("0")) {
				pollData["amount_global_abstain"] = pollInfo["amount_global_abstain"]
			}

			pollData["votesvalid"] = pollInfo["votesvalid"]
			pollData["votesinvalid"] = pollInfo["votesinvalid"]
			pollData["votescast"] = pollInfo["votescast"]
		}

		out := struct {
			Assignment struct {
				Title string `json:"title"`
			} `json:"assignment"`
			Poll map[string]json.RawMessage `json:"poll"`
		}{
			assignment,
			pollData,
		}

		b, err = json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("decoding slide: %w", err)
		}
		return b, nil
	}
}

type openDataType struct {
	weight int
	userID int
	data   map[string]interface{}
}

func (o *openDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.data)
}
