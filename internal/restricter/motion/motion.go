package motion

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

type required interface {
	restricter.HasPermer
	InGroups(uid int, groups []int) bool
}

// Restrict restricts motions/motion.
func Restrict(r required) restricter.ElementFunc {
	return func(uid int, data json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, "motions.can_see") {
			return nil, nil
		}

		var motion struct {
			Submitters  []int    `json:"submitters"`
			Restriction []string `json:"state_restriction"`
			Comments    []struct {
				ReadGroups []int `json:"read_groups_id"`
			} `json:"comments"`
		}

		if err := json.Unmarshal(data, &motion); err != nil {
			return nil, fmt.Errorf("decode motion: %w", err)
		}

		var isSumitter bool
		for _, sid := range motion.Submitters {
			if sid == uid {
				isSumitter = true
				break
			}
		}

		permission := r.HasPerm(uid, "motions.can_manage") || len(motion.Restriction) == 0

		if !permission {
			for _, value := range motion.Restriction {
				if (value == "motions.can_see_internal" || value == "motions.can_manage_metadata" || value == "motions.can_manage") && r.HasPerm(uid, value) {
					permission = true
					break
				}
				if value == "is_submitter" && isSumitter {
					permission = true
					break
				}
			}
		}

		if !permission {
			return nil, nil
		}

		var motionData map[string]json.RawMessage
		if err := json.Unmarshal(data, &motionData); err != nil {
			return nil, fmt.Errorf("decode motion data: %w", err)
		}

		var comments []json.RawMessage
		if err := json.Unmarshal(motionData["comments"], &comments); err != nil {
			return nil, fmt.Errorf("decode motion comments: %w", err)
		}

		newComments := make([]json.RawMessage, 0)

		for i, c := range comments {
			if r.InGroups(uid, motion.Comments[i].ReadGroups) {
				newComments = append(newComments, c)
			}
		}

		newCommentsEncoded, err := json.Marshal(newComments)
		if err != nil {
			return nil, fmt.Errorf("encode comments: %w", err)
		}

		motionData["comments"] = newCommentsEncoded

		data, err = json.Marshal(motionData)
		if err != nil {
			return nil, fmt.Errorf("encode motion: %w", err)
		}

		return data, nil
	}
}

// BlockRestrict restricts motions/motion-block.
func BlockRestrict(r restricter.HasPermer) restricter.ElementFunc {
	return func(uid int, data json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, "motions.can_see") {
			return nil, nil
		}

		if r.HasPerm(uid, "motions.can_manage") {
			return data, nil
		}

		var block struct {
			Internal bool `json:"internal"`
		}
		if err := json.Unmarshal(data, &block); err != nil {
			return nil, fmt.Errorf("decode block: %w", err)
		}

		if !block.Internal {
			return data, nil
		}

		return nil, nil
	}
}

// CommentSectionRestrict restricts motions/motion-comment-section.
func CommentSectionRestrict(r required) restricter.ElementFunc {
	return func(uid int, data json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, "motions.can_see") {
			return nil, nil
		}

		if r.HasPerm(uid, "motions.can_manage") {
			return data, nil
		}

		var cs struct {
			ReadGroups []int `json:"read_groups_id"`
		}
		if err := json.Unmarshal(data, &cs); err != nil {
			return nil, fmt.Errorf("decode comment section: %w", err)
		}

		if r.InGroups(uid, cs.ReadGroups) {
			return data, nil
		}

		return nil, nil
	}
}

// ChangeRecommendationRestrict restricts motions/motion-change-recommendation.
func ChangeRecommendationRestrict(r restricter.HasPermer) restricter.ElementFunc {
	return func(uid int, data json.RawMessage) (json.RawMessage, error) {
		if !r.HasPerm(uid, "motions.can_see") {
			return nil, nil
		}

		if r.HasPerm(uid, "motions.can_manage") {
			return data, nil
		}

		var cr struct {
			Internal bool `json:"internal"`
		}
		if err := json.Unmarshal(data, &cr); err != nil {
			return nil, fmt.Errorf("decode change recommendation: %w", err)
		}

		if !cr.Internal {
			return data, nil
		}

		return nil, nil
	}
}
