package test

import (
	"context"
	"net/http"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/auth"
)

// AutherMock implements the autoupdate.Auther interface.
type AutherMock struct {
	UID        int
	AuthCalled bool
}

// Middleware returns a mock auth errorHandleFunc.
func (a *AutherMock) Middleware(next func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		a.AuthCalled = true
		ctx := context.WithValue(r.Context(), auth.UserIDKey, a.UID)
		r = r.WithContext(ctx)
		return next(w, r)
	}
}
