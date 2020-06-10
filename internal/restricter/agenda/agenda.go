package agenda

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

// Restrict handels restrictions of agenda/item elements.
func Restrict(r restricter.HasPermer) restricter.ElementFunc {
	return func(uid int, element json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, "agenda.can_see") {
			return nil, nil
		}

		var agenda struct {
			IsHidden   bool `json:"is_hidden"`
			IsInternal bool `json:"is_internal"`
		}
		if err := json.Unmarshal(element, &agenda); err != nil {
			return nil, fmt.Errorf("decoding item: %w", err)
		}

		canManage := r.HasPerm(uid, "agenda.can_manage")
		canSeeInternal := r.HasPerm(uid, "agenda.can_see_internal_items")

		if !canManage && agenda.IsHidden {
			return nil, nil
		}

		if !canSeeInternal && agenda.IsInternal {
			return nil, nil
		}

		if canManage && canSeeInternal {
			return element, nil
		}

		var agendaData map[string]json.RawMessage
		if err := json.Unmarshal(element, &agendaData); err != nil {
			return nil, fmt.Errorf("decoding itemdata: %w", err)
		}

		if !canSeeInternal {
			delete(agendaData, "duration")
		}

		if !canManage {
			delete(agendaData, "comment")
		}

		element, err := json.Marshal(agendaData)
		if err != nil {
			return nil, fmt.Errorf("encoding itemdata: %w", err)
		}
		return element, nil
	}
}

// RequiredSpeakers returns the user ids of a list of speaker objekt.
func RequiredSpeakers(data json.RawMessage) ([]int, error) {
	var los struct {
		Speakers []struct {
			UserID int `json:"user_id"`
		} `json:"speakers"`
	}
	if err := json.Unmarshal(data, &los); err != nil {
		return nil, fmt.Errorf("unmarshal list of speaker: %w", err)
	}

	var uids []int
	for _, s := range los.Speakers {
		uids = append(uids, s.UserID)
	}
	return uids, nil
}
