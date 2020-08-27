package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	whoAmIPath    = "/apps/users/whoami/"
	whoamITimeout = 2 * time.Second
)

type userIDKey string

// UserIDKey is the key of the request.Context where the user id is saved to.
const UserIDKey userIDKey = "user_id"

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

func (a *Auth) whoami(r *http.Request) (int, error) {
	ctx, close := context.WithTimeout(r.Context(), whoamITimeout)
	defer close()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.addr+whoAmIPath, nil)
	if err != nil {
		return 0, fmt.Errorf("creating whoami request: %w", err)
	}

	// Write cookies from the original request to the whoami request.
	for _, c := range r.Cookies() {
		req.AddCookie(c)
	}

	resp, err := http.DefaultClient.Do(req)
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

// Middleware adds the user id into the request.Context.
func (a *Auth) Middleware(next func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		uid, err := a.whoami(r)
		if err != nil {
			return fmt.Errorf("authenticate request: %w", err)
		}

		ctx := context.WithValue(r.Context(), UserIDKey, uid)
		r = r.WithContext(ctx)

		return next(w, r)
	}
}
