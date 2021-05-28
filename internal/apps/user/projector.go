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

// GetUserShortName returns the short display name for an user id (name without structure_level)
func GetUserShortName(ds projector.Datastore, uid int) ([]byte, error) {
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

	return json.Marshal(strings.Join(parts, " "))
}

// GetUserLevel returns the structre level for user id as json.
func GetUserLevel(ds projector.Datastore, uid int) ([]byte, error) {
	var user struct {
		Level json.RawMessage `json:"structure_level"`
	}

	if err := ds.Get("users/user", uid, &user); err != nil {
		return nil, projector.NewClientError("users/user with id %d does not exist", uid)
	}

	return user.Level, nil
}

// GetUserName returns the display name for an user id.
func GetUserName(ds projector.Datastore, uid int) ([]byte, error) {
	shortName, err := GetUserShortName(ds, uid)
	if err != nil {
		return nil, fmt.Errorf("getting short name: %w", err)
	}

	structureLevel, err := GetUserLevel(ds, uid)
	if err != nil {
		return nil, fmt.Errorf("get user level: %w", err)
	}

	if structureLevel == nil {
		return shortName, nil
	}

	var decoded string
	if err := json.Unmarshal(shortName, &decoded); err != nil {
		return nil, fmt.Errorf("decoding short name: %w", err)
	}
	decoded = fmt.Sprintf("%s (%s)", decoded, structureLevel)

	shortName, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("encoding short name: %w", err)
	}

	return shortName, nil
}
