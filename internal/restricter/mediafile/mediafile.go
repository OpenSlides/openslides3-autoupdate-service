package mediafile

import (
	"encoding/json"
	"fmt"
)

// Mediafile restricts a mediafile object.
type Mediafile struct {
	required
}

// New creates a new Mediafile instance.
func New(r required) *Mediafile {
	return &Mediafile{r}
}

// Restrict restricts one mediafile.
func (m *Mediafile) Restrict(uid int, data json.RawMessage) (json.RawMessage, error) {
	if !m.HasPerm(uid, "mediafiles.can_see") {
		return nil, nil
	}

	if m.IsSuperadmin(uid) {
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
		if m.InGroups(uid, accessGroups) {
			return data, nil
		}
	default:
		return nil, fmt.Errorf("invalid type %T for mediafile.inherited_access_groups_id", accessGroups)
	}
	return nil, nil
}
