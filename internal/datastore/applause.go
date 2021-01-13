package datastore

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type applause struct {
	c            *config
	presentUsers map[int]bool

	mu          sync.RWMutex
	waitSeconds int
	base        int
}

func (a *applause) update(data map[string]json.RawMessage) error {
	if a.presentUsers == nil {
		a.presentUsers = make(map[int]bool)
	}

	var applauseTimeout int
	if err := a.c.ConfigValue("general_system_stream_applause_timeout", &applauseTimeout); err != nil {
		var d doesNotExistError
		if !errors.As(err, &d) {
			return fmt.Errorf("getting applause timeout: %w", err)
		}
		applauseTimeout = 5
	}

	for elementID, v := range data {
		parts := strings.Split(elementID, ":")
		if len(parts) != 2 {
			return fmt.Errorf("key %s has wrong format. Expected one `:`", elementID)
		}

		if parts[0] != "users/user" {
			continue
		}

		userID, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("getting userID for %s: %w", parts[1], err)
		}

		var user struct {
			Present bool `json:"is_present"`
		}
		if err := json.Unmarshal(v, &user); err != nil {
			return fmt.Errorf("decoding user %d: %w", userID, err)
		}

		a.presentUsers[userID] = user.Present
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.waitSeconds = applauseTimeout

	a.base = 0
	for _, present := range a.presentUsers {
		if present {
			a.base++
		}
	}
	return nil
}

func (a *applause) ApplauseConfig() (waitTime int, base int) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.waitSeconds, a.base
}
