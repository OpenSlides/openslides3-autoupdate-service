package chat

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

type required interface {
	restricter.HasPermer
	InGroups(uid int, groups []int) bool
}

// Restrict restricts chat/chat-group and chat/chat-message.
func Restrict(r required) restricter.ElementFunc {
	return func(uid int, data json.RawMessage) (json.RawMessage, error) {
		if r.HasPerm(uid, "chat.can_manage") {
			return data, nil
		}

		var chatobject struct {
			AccessGroupsId []int `json:"access_groups_id"`
		}
		if err := json.Unmarshal(data, &chatobject); err != nil {
			return nil, fmt.Errorf("decode chat-group/chat-message: %w", err)
		}

		if r.InGroups(uid, chatobject.AccessGroupsId) {
			return data, nil
		}

		return nil, nil
	}
}
