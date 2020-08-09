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

	req.Method = http.MethodGet
	req.Body = nil
	req.ContentLength = 0

	// This is inspired by httputil.ReverseProxy
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

// Middleware adds the user id into the request.Context.
func (a *Auth) Middleware(next func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		uid, err := a.Auth(r)
		if err != nil {
			return fmt.Errorf("authenticate request")
		}

		ctx := context.WithValue(r.Context(), UserIDKey, uid)
		r = r.WithContext(ctx)

		return next(w, r)
	}
}
