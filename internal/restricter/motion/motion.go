package motion

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

// Motion restricts motions/motion.
type Motion struct {
	required
}

// NewMotion returns a new Motion.
func NewMotion(required required) *Motion {
	return &Motion{required}
}

// Restrict restricts the motion.
func (m *Motion) Restrict(uid int, data json.RawMessage) (json.RawMessage, error) {
	if !m.HasPerm(uid, "motions.can_see") {
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

	permission := m.HasPerm(uid, "motions.can_manage") || len(motion.Restriction) == 0

	if !permission {
		for _, value := range motion.Restriction {
			if (value == "motions.can_see_internal" || value == "motions.can_manage_metadata" || value == "motions.can_manage") && m.HasPerm(uid, value) {
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
		if m.InGroups(uid, motion.Comments[i].ReadGroups) {
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

// Block restricts motions/motion-block.
type Block struct {
	restricter.HasPermer
}

// NewBlock initializes a new Block.
func NewBlock(h restricter.HasPermer) *Block {
	return &Block{h}
}

// Restrict restricts the block.
func (b *Block) Restrict(uid int, data json.RawMessage) (json.RawMessage, error) {
	if !b.HasPerm(uid, "motions.can_see") {
		return nil, nil
	}

	if b.HasPerm(uid, "motions.can_manage") {
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

// CommentSection restricts motions/motion-comment-section.
type CommentSection struct {
	required
}

// NewCommentSection initializes a new CommentSection.
func NewCommentSection(r required) *CommentSection {
	return &CommentSection{r}
}

// Restrict restricts the block.
func (c *CommentSection) Restrict(uid int, data json.RawMessage) (json.RawMessage, error) {
	if !c.HasPerm(uid, "motions.can_see") {
		return nil, nil
	}

	if c.HasPerm(uid, "motions.can_manage") {
		return data, nil
	}

	var cs struct {
		ReadGroups []int `json:"read_groups_id"`
	}
	if err := json.Unmarshal(data, &cs); err != nil {
		return nil, fmt.Errorf("decode comment section: %w", err)
	}

	if c.InGroups(uid, cs.ReadGroups) {
		return data, nil
	}

	return nil, nil
}

// ChangeRecommendation restricts motions/motion-change-recommendation.
type ChangeRecommendation struct {
	restricter.HasPermer
}

// NewChangeRecommendation initializes a new ChangeRecommendation.
func NewChangeRecommendation(r restricter.HasPermer) *ChangeRecommendation {
	return &ChangeRecommendation{r}
}

// Restrict restricts the ChangeRecommendation.
func (c *ChangeRecommendation) Restrict(uid int, data json.RawMessage) (json.RawMessage, error) {
	if !c.HasPerm(uid, "motions.can_see") {
		return nil, nil
	}

	if c.HasPerm(uid, "motions.can_manage") {
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
