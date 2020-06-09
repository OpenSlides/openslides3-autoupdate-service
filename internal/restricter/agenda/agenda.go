package agenda

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/restricter"
)

// Item handels restrictions of agenda/item elements.
type Item struct {
	restricter.HasPermer
}

// NewItem creates a new Item.
func NewItem(hasPermer restricter.HasPermer) *Item {
	return &Item{
		HasPermer: hasPermer,
	}
}

// Restrict restricts the element.
func (i *Item) Restrict(uid int, element json.RawMessage) (json.RawMessage, error) {
	if !i.HasPerm(uid, "agenda.can_see") {
		return nil, nil
	}

	var agenda struct {
		IsHidden   bool `json:"is_hidden"`
		IsInternal bool `json:"is_internal"`
	}
	if err := json.Unmarshal(element, &agenda); err != nil {
		return nil, fmt.Errorf("decoding item: %w", err)
	}

	canManage := i.HasPerm(uid, "agenda.can_manage")
	canSeeInternal := i.HasPerm(uid, "agenda.can_see_internal_items")

	if !canManage && agenda.IsHidden {
		return nil, nil
	}

	if !canSeeInternal && agenda.IsInternal {
		return nil, nil
	}

	if canManage && canSeeInternal {
		return element, nil
	}

	var agendaData map[string]json.RawMessage
	if err := json.Unmarshal(element, &agendaData); err != nil {
		return nil, fmt.Errorf("decoding itemdata: %w", err)
	}

	if !canSeeInternal {
		delete(agendaData, "duration")
	}

	if !canManage {
		delete(agendaData, "comment")
	}

	element, err := json.Marshal(agendaData)
	if err != nil {
		return nil, fmt.Errorf("encoding itemdata: %w", err)
	}
	return element, nil
}
