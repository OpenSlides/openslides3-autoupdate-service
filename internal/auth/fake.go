package auth

import (
	"context"
	"net/http"
)

// Fake implements the http.Auther interface but always returns the same user
// id.
type Fake int

// Authenticate returns a context with the user id the Faker was initiaized
// with.
func (f Fake) Authenticate(r *http.Request) (context.Context, error) {
	return context.WithValue(r.Context(), userIDType, int(f)), nil
}
