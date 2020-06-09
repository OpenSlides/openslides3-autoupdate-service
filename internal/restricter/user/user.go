package user

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

// User handels restrictions of users/user elements.
type User struct {
	restricter.HasPermer
}

// NewUser creates a new User.
func NewUser(hasPermer restricter.HasPermer) *User {
	return &User{
		HasPermer: hasPermer,
	}
}

// Restrict returns the restricted version of the element.
func (u *User) Restrict(uid int, element json.RawMessage) (json.RawMessage, error) {
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

	if u.HasPerm(uid, "users.can_see_name") {
		if u.HasPerm(uid, "users.can_see_extra_data") {
			if u.HasPerm(uid, "users.can_manage") {
				return filter(element, allDataFields, 0, nil)
			}
			return filter(element, manyDataFields, 0, nil)
		}
		return filter(element, littleDataFields, uid, ownDataFields)
	}
	// TODO: build list of needed users
	return nil, nil
}

// Group handles users/group elements.
func Group(uid int, element json.RawMessage) (json.RawMessage, error) {
	return element, nil
}

// PersonalNote is the restricter for users/personal_note
func PersonalNote(uid int, data json.RawMessage) (json.RawMessage, error) {
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

func filter(value json.RawMessage, fields []string, uid int, ownfields []string) (json.RawMessage, error) {
	var allData map[string]json.RawMessage
	if err := json.Unmarshal(value, &allData); err != nil {
		return nil, fmt.Errorf("unmarshall data: %w", err)
	}

	filteredData := make(map[string]json.RawMessage, len(fields))

	id, err := strconv.Atoi(string(allData["id"]))
	if err != nil {
		return nil, fmt.Errorf("invalid id: %s", allData["id"])
	}

	useFields := fields
	if ownfields != nil && uid == id {
		useFields = ownfields
	}

	for _, k := range useFields {
		filteredData[k] = allData[k]
	}

	filtered, err := json.Marshal(filteredData)
	if err != nil {
		return nil, fmt.Errorf("remarshal data: %w", err)
	}
	return filtered, nil
}
