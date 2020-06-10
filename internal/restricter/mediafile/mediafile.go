package mediafile

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

type required interface {
	restricter.HasPermer
	IsSuperadmin(uid int) bool
	InGroups(uid int, groups []int) bool
}

// Restrict restricts a mediafile object.
func Restrict(r required) restricter.ElementFunc {
	return func(uid int, data json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, "mediafiles.can_see") {
			return nil, nil
		}

		if r.IsSuperadmin(uid) {
			return data, nil
		}

		var media struct {
			InheritedAccess interface{} `json:"inherited_access_groups_id"`
		}

		if err := json.Unmarshal(data, &media); err != nil {
			return nil, fmt.Errorf("decoding mediafile: %w", err)
		}

		switch accessGroups := media.InheritedAccess.(type) {
		case bool:
			if accessGroups {
				return data, nil
			}
		case []int:
			if r.InGroups(uid, accessGroups) {
				return data, nil
			}
		default:
			return nil, fmt.Errorf("invalid type %T for mediafile.inherited_access_groups_id", accessGroups)
		}
		return nil, nil
	}
}
