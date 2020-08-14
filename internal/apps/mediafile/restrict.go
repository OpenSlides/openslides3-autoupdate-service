package mediafile

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

const (
	pCanSee = "mediafiles.can_see"
)

// Restrict restricts a mediafile object.
func Restrict(r restricter.HasPermer) restricter.ElementFunc {
	return func(uid int, data json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, pCanSee) {
			return nil, nil
		}

		if r.IsSuperadmin(uid) {
			return data, nil
		}

		var media struct {
			InheritedAccess boolOrIntSlice `json:"inherited_access_groups_id"`
		}

		if err := json.Unmarshal(data, &media); err != nil {
			return nil, fmt.Errorf("decoding mediafile: %w", err)
		}

		if accessGroups := media.InheritedAccess; accessGroups.isBool() {
			if accessGroups.b {
				return data, nil
			}
			return nil, nil
		}

		if r.InGroups(uid, media.InheritedAccess.iList) {
			return data, nil
		}
		return nil, nil
	}
}

type boolOrIntSlice struct {
	b     bool
	iList []int
}

func (e *boolOrIntSlice) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &e.b); err == nil {
		return nil
	}

	if err := json.Unmarshal(data, &e.iList); err != nil {
		return fmt.Errorf("unmarshal boolOrIntSlice: %w", err)
	}

	return nil
}

func (e *boolOrIntSlice) isBool() bool {
	return e.iList == nil
}
