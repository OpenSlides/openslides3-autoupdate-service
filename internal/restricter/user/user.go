package user

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

type required interface {
	restricter.HasPermer
	UserRequired(uid int) bool
}

// Restrict handels restrictions of users/user elements.
func Restrict(r required) restricter.ElementFunc {
	return func(uid int, element json.RawMessage) (json.RawMessage, error) {
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
		}

		manyDataFields := append(littleDataFields, "gender", "email", "last_email_send", "comment", "is_active", "auth_type")
		allDataFields := append(manyDataFields, "default_password")
		ownDataFields := append(littleDataFields, "email", "gender")

		var user struct {
			ID int `json:"id"`
		}
		if err := json.Unmarshal(element, &user); err != nil {
			return nil, fmt.Errorf("unmarshal user: %w", err)
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

		if !r.UserRequired(user.ID) {
			return nil, nil
		}

		// TODO: check permission of required user
		return filter(element, littleDataFields)
	}
}

// PersonalNoteRestrict is the restricter for users/personal_note
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
