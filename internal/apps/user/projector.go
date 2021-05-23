package user

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

// Slide renders "users/user".
func Slide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		var element struct {
			ID int `json:"id"`
		}
		if err := json.Unmarshal(e, &element); err != nil {
			var element struct {
				ID interface{} `json:"id"`
			}

			// Fallback for better error messages
			if err := json.Unmarshal(e, &element); err == nil {
				return nil, projector.NewClientError("users/user with id %s does not exist", element.ID)
			}
			return nil, fmt.Errorf("decoding element: %w", err)
		}

		if element.ID == 0 {
			return nil, projector.NewClientError("id is required for users/user slide")
		}

		name, err := GetUserName(ds, element.ID)
		if err != nil {
			return nil, fmt.Errorf("get username: %w", err)
		}

		return []byte(fmt.Sprintf(`{"user":%s}`, name)), nil
	}
}

// GetUserName returns the display name for an user id as valid json.
func GetUserName(ds projector.Datastore, uid int) ([]byte, error) {
	var user struct {
		Username  string `json:"username"`
		Title     string `json:"title"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Level     string `json:"structure_level"`
	}

	if err := ds.Get("users/user", uid, &user); err != nil {
		return nil, projector.NewClientError("users/user with id %d does not exist", uid)
	}

	parts := func(sp ...string) []string {
		var full []string
		for _, s := range sp {
			if s == "" {
				continue
			}
			full = append(full, s)
		}
		return full
	}(user.Title, user.FirstName, user.LastName)

	if len(parts) == 0 {
		return json.Marshal(user.Username)
	}

	if user.Level != "" {
		parts = append(parts, fmt.Sprintf("(%s)", user.Level))
	}

	return json.Marshal(strings.Join(parts, " "))
}
