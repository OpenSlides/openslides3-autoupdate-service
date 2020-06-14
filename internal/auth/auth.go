package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	whoAmIPath    = "/apps/users/whoami/"
	whoamITimeout = 2 * time.Second
)

// Auth authentivates a request using the whoami view.
type Auth struct {
	addr string
}

// New creates a new Auth instance.
func New(workerAddr string) *Auth {
	return &Auth{
		addr: workerAddr,
	}
}

// Auth returns a user id for a given request.
func (a *Auth) Auth(r *http.Request) (int, error) {
	ctx, close := context.WithTimeout(r.Context(), whoamITimeout)
	defer close()
	req := r.Clone(ctx)
	req.Close = false

	var err error
	req.URL, err = url.Parse(a.addr + whoAmIPath)
	if err != nil {
		return 0, fmt.Errorf("build whoami url: %w", err)
	}

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return 0, fmt.Errorf("send whoami request: %w", err)
	}
	defer resp.Body.Close()

	var respData struct {
		UserID       *int `json:"user_id"`
		GuestEnabled bool `json:"guest_enabled"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return 0, fmt.Errorf("decoding response: %w", err)
	}

	if !respData.GuestEnabled && respData.UserID == nil {
		return 0, new(errorNoAnonymous)
	}

	if respData.UserID == nil {
		return 0, nil
	}

	return *respData.UserID, nil
}
