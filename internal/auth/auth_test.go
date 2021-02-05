package auth_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/auth"
)

const testSecred = "@$kze04dxf37t=&1z$5bdem-6nf1wpo3)rcntc5a@n60pllc6m"

func TestAuth(t *testing.T) {
	anonymous := new(anonymousMock)
	a := auth.New("test-auth-cookie", testSecred, new(backendMock), anonymous)

	for _, tt := range []struct {
		name             string
		request          func() *http.Request
		anonymousEnabled bool
		uid              int
		err              bool
	}{
		{
			"anonymous allowed",
			func() *http.Request {
				r, err := http.NewRequest("GET", "openslides.com/service/autoupdate", nil)
				if err != nil {
					t.Fatalf("Can not create request: %v", err)
				}
				return r
			},
			true,
			0,
			false,
		},

		{
			"anonymous not allowed",
			func() *http.Request {
				r, err := http.NewRequest("GET", "openslides.com/service/autoupdate", nil)
				if err != nil {
					t.Fatalf("Can not create request: %v", err)
				}
				return r
			},
			false,
			0,
			true,
		},

		{
			"specific user",
			func() *http.Request {
				r, err := http.NewRequest("GET", "openslides.com/service/autoupdate", nil)
				r.AddCookie(&http.Cookie{Name: "test-auth-cookie", Value: "sessionID"})
				if err != nil {
					t.Fatalf("Can not create request: %v", err)
				}
				return r
			},
			true,
			1,
			false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			anonymous.enabled = tt.anonymousEnabled

			ctx, err := a.Authenticate(tt.request())
			if tt.err {
				if err == nil {
					t.Errorf("Auth did not returned expected error")
				}
				return
			}

			if err != nil {
				t.Fatalf("Auth returend unexpected err: %v", err)
			}

			if uid := auth.FromContext(ctx); uid != tt.uid {
				t.Errorf("Auth returned uid %d, expected %d", uid, tt.uid)
			}
		})
	}
}

type backendMock struct{}

func (b *backendMock) GetSession(sessionID string) ([]byte, error) {
	return []byte("MTQxZDQxNzUwMWRjMGVjN2NlNWE4NjUwZDZjMTBiMzljODk4MTNiYjp7Il9hdXRoX3VzZXJfaWQiOiIxIiwiX2F1dGhfdXNlcl9iYWNrZW5kIjoiZGphbmdvLmNvbnRyaWIuYXV0aC5iYWNrZW5kcy5Nb2RlbEJhY2tlbmQiLCJfYXV0aF91c2VyX2hhc2giOiJhYTdmNmFlNzhiOGM2YThlNDMyMzIwMmVlYTY0OWFiOTNlNWUxM2MwIn0="), nil
}

type anonymousMock struct {
	enabled bool
}

func (a *anonymousMock) ConfigValue(key string, v interface{}) error {
	rv := reflect.ValueOf(v)
	rv.Elem().SetBool(a.enabled)
	return nil
}
