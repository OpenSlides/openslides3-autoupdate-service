package mediafile

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

const (
	pCanSee = "mediafiles.can_see"
)

type required interface {
	restricter.HasPermer
	IsSuperadmin(uid int) bool
	InGroups(uid int, groups []int) bool
}

// Restrict restricts a mediafile object.
func Restrict(r required) restricter.ElementFunc {
	return func(uid int, data json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, pCanSee) {
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
		case []interface{}:
			ints := make([]int, len(accessGroups))
			for i, g := range accessGroups {
				j, ok := g.(int)
				if !ok {
					return nil, fmt.Errorf("mediafile.inherited_access_groups_id[%d] has type %T, expected int", i, g)
				}

				ints[i] = j
			}
			if r.InGroups(uid, ints) {
				return data, nil
			}
		default:
			return nil, fmt.Errorf("mediafile.inherited_access_groups_id has invalid type %T", accessGroups)
		}
		return nil, nil
	}
}
