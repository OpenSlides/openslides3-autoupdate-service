package test

import (
	"net/http"
)

// AutherMock implements the autoupdate.Auther interface.
type AutherMock struct{}

// Middleware returns a mock auth errorHandleFunc.
func (a *AutherMock) Middleware(next func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		return next(w, r)
	}
}
