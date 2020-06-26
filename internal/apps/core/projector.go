package core

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/projector"
)

// CountdownSlide renders "core/countdown".
func CountdownSlide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		c, err := projector.ModelFromElement(ds, e, "core/countdown")
		if err != nil {
			return nil, fmt.Errorf("getting core/countdown: %w", err)
		}

		var countdown struct {
			Description   json.RawMessage `json:"description"`
			Running       json.RawMessage `json:"running"`
			CountdownTime json.RawMessage `json:"countdown_time"`
			WarningTime   json.RawMessage `json:"warning_time"`
		}
		if err := json.Unmarshal(c, &countdown); err != nil {
			return nil, fmt.Errorf("decoding countdown: %w", err)
		}

		countdown.WarningTime = ds.ConfigValue("agenda_countdown_warning_time")

		out, err := json.Marshal(countdown)
		if err != nil {
			return nil, fmt.Errorf("decoding countdown: %w", err)
		}
		return out, nil
	}
}

// MessageSlide renders "core/projector-message".
func MessageSlide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		m, err := projector.ModelFromElement(ds, e, "core/projector-message")
		if err != nil {
			return nil, fmt.Errorf("getting core/projector-message: %w", err)
		}

		return m, nil
	}
}

// ClockSlide renders "core/clock".
func ClockSlide() projector.CallableFunc {
	return func(ds projector.Datastore, e json.RawMessage, pid int) (json.RawMessage, error) {
		return []byte(`{}`), nil
	}
}
