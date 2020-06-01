package test

import "net/http"

// AutherMock implements the autoupdate.Auther interface.
type AutherMock struct {
	UID        int
	AuthCalled bool
}

// Auth returns the given UID.
func (a *AutherMock) Auth(r *http.Request) (int, error) {
	a.AuthCalled = true
	return a.UID, nil
}
