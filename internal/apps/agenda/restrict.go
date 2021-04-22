package agenda

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

const (
	pCanSee         = "agenda.can_see"
	pCanManage      = "agenda.can_manage"
	pCanSeeInternal = "agenda.can_see_internal_items"

	// CanSeeListOfSpeakers is the permission string if a user can see the list
	// of speakers.
	CanSeeListOfSpeakers     = "agenda.can_see_list_of_speakers"
	pCanManageListOfSpeakers = "agenda.can_manage_list_of_speakers"
)

// RestrictItem handels restrictions of agenda/item elements.
func RestrictItem(r restricter.HasPermer) restricter.ElementFunc {
	return func(uid int, element json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, pCanSee) {
			return nil, nil
		}

		var agenda struct {
			IsHidden   bool `json:"is_hidden"`
			IsInternal bool `json:"is_internal"`
		}
		if err := json.Unmarshal(element, &agenda); err != nil {
			return nil, fmt.Errorf("decoding item: %w", err)
		}

		canManage := r.HasPerm(uid, pCanManage)
		canSeeInternal := r.HasPerm(uid, pCanSeeInternal)

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

// RestrictListOfSpeakers restricts the list of speakers: The field "note" is
// removed for every speaker, if the user is not a manager, the speaker user or
// the config agenda_list_of_speakers_speaker_note_for_everyone is set.
func RestrictListOfSpeakers(r restricter.HasPermer) restricter.ElementFunc {
	return func(uid int, element json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, CanSeeListOfSpeakers) {
			return nil, nil
		}

		var notesForEveryone bool
		if err := r.ConfigValue("agenda_list_of_speakers_speaker_note_for_everyone", &notesForEveryone); err != nil {
			return nil, fmt.Errorf("getting agenda_list_of_speakers_speaker_note_for_everyone: %w", err)
		}

		if notesForEveryone || r.HasPerm(uid, pCanManageListOfSpeakers) {
			return element, nil
		}

		// Delete each speaker's note for speakers which user is not the one to restrict for.
		var los map[string]json.RawMessage
		if err := json.Unmarshal(element, &los); err != nil {
			return nil, fmt.Errorf("decoding list of speakers: %w", err)
		}
		var speakers []map[string]json.RawMessage
		if err := json.Unmarshal(los["speakers"], &speakers); err != nil {
			return nil, fmt.Errorf("decoding speakers: %w", err)
		}

		for _, speaker := range speakers {
			var userId int
			if err := json.Unmarshal(speaker["user_id"], &userId); err != nil {
				return nil, fmt.Errorf("decoding speaker: %w", err)
			}
			if uid == userId {
				continue
			}
			delete(speaker, "note")
		}

		speakersEncoded, err := json.Marshal(speakers)
		if err != nil {
			return nil, fmt.Errorf("encoding speakers: %w", err)
		}

		los["speakers"] = speakersEncoded
		element, err = json.Marshal(los)
		if err != nil {
			return nil, fmt.Errorf("encoding list of speakers: %w", err)
		}
		return element, nil
	}
}
