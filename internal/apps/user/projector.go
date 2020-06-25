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
			return nil, fmt.Errorf("decoding element: %w", err)
		}

		name, err := getUserName(ds, element.ID)
		if err != nil {
			return nil, fmt.Errorf("get username: %w", err)
		}

		return []byte(fmt.Sprintf(`{"user":"%s"}`, name)), nil
	}
}

func getUserName(ds projector.Datastore, uid int) (string, error) {
	u := ds.Get(fmt.Sprintf("users/user:%d", uid))

	var user struct {
		Username  string `json:"username"`
		Title     string `json:"title"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Level     string `json:"structure_level"`
	}
	if err := json.Unmarshal(u, &user); err != nil {
		return "", fmt.Errorf("decoding user: %w", err)
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
		return user.Username, nil
	}

	if user.Level != "" {
		parts = append(parts, fmt.Sprintf("(%s)", user.Level))
	}

	return strings.Join(parts, " "), nil
}
