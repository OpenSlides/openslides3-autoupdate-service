package test

import (
	"context"
	"net/http"
)

// AutherMock implements the autoupdate.Auther interface.
type AutherMock struct{}

// Authenticate returns the request context.
func (a *AutherMock) Authenticate(r *http.Request) (context.Context, error) {
	return r.Context(), nil
}
