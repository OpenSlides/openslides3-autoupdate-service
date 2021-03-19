package user

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

// Restrict handels restrictions of users/user elements.
func Restrict(r restricter.HasPermer) restricter.ElementFunc {
	littleDataFields := []string{
		"id",
		"username",
		"title",
		"first_name",
		"last_name",
		"structure_level",
		"number",
		"about_me",
		"groups_id",
		"is_present",
		"is_committee",
		"vote_weight",
		"gender",
	}
	manyDataFields := append(littleDataFields, "email", "last_email_send", "comment", "is_active", "auth_type", "vote_delegated_to_id", "vote_delegated_from_users_id")
	allDataFields := append(manyDataFields, "default_password")
	ownDataFields := append(littleDataFields, "email", "gender", "vote_delegated_to_id", "vote_delegated_from_users_id")

	return func(uid int, element json.RawMessage) (json.RawMessage, error) {
		var user struct {
			ID int `json:"id"`
		}
		if err := json.Unmarshal(element, &user); err != nil {
			return nil, fmt.Errorf("unmarshal user id: %w", err)
		}
		if user.ID == uid && !r.HasPerm(uid, "users.can_see_extra_data") {
			return filter(element, ownDataFields)
		}

		if r.HasPerm(uid, "users.can_see_name") {
			if r.HasPerm(uid, "users.can_see_extra_data") {
				if r.HasPerm(uid, "users.can_manage") {
					return filter(element, allDataFields)
				}
				return filter(element, manyDataFields)
			}
			return filter(element, littleDataFields)
		}

		// Registered users can see there delegated users.
		if uid != 0 {
			// Get the users `vote_delegated_from_users_id`.
			var requestUser struct {
				VoteDelegationIds []int `json:"vote_delegated_from_users_id"`
			}
			if err := r.Get("users/user", uid, &requestUser); err != nil {
				return nil, fmt.Errorf("fetching and unmarshalling user: %w", err)
			}
			// The user.ID is required, if it is in VoteDelegationIds.
			for _, id := range requestUser.VoteDelegationIds {
				if id == user.ID {
					return filter(element, littleDataFields)
				}
			}
		}

		for _, perm := range r.UserRequired(user.ID) {
			if !r.HasPerm(uid, perm) {
				continue
			}

			return filter(element, littleDataFields)
		}

		return nil, nil
	}
}

// PersonalNoteRestrict is the restricter for users/personal_note.
func PersonalNoteRestrict(uid int, data json.RawMessage) (json.RawMessage, error) {
	if uid == 0 {
		return nil, nil
	}

	var element struct {
		UserID int `json:"user_id"`
	}
	if err := json.Unmarshal(data, &element); err != nil {
		return nil, fmt.Errorf("decoding personal element: %w", err)
	}

	if element.UserID != uid {
		return nil, nil
	}

	return data, nil
}

func filter(value json.RawMessage, fields []string) (json.RawMessage, error) {
	var allData map[string]json.RawMessage
	if err := json.Unmarshal(value, &allData); err != nil {
		return nil, fmt.Errorf("unmarshall data: %w", err)
	}

	filteredData := make(map[string]json.RawMessage, len(fields))

	for _, k := range fields {
		filteredData[k] = allData[k]
	}

	filtered, err := json.Marshal(filteredData)
	if err != nil {
		return nil, fmt.Errorf("remarshal data: %w", err)
	}
	return filtered, nil
}
