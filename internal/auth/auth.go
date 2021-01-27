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
		return 0, NoAnonymousError{}
	}

	if respData.UserID == nil {
		return 0, nil
	}

	return *respData.UserID, nil
}

// Authenticate returns an context that can be used by auth.FromContext().
func (a *Auth) Authenticate(r *http.Request) (context.Context, error) {
	uid, err := a.whoami(r)
	if err != nil {
		return nil, fmt.Errorf("sending whoami: %w", err)
	}

	return context.WithValue(r.Context(), userIDType, uid), nil
}

// FromContext returnes the user id from a context that was called insinde the
// Middleware.
func FromContext(ctx context.Context) int {
	v := ctx.Value(userIDType)
	if v == nil {
		return 0
	}

	return v.(int)
}

type authString string

const userIDType authString = "user_id"
