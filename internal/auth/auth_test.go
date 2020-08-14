package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	whoami := new(WhoAmIMock)
	srv := httptest.NewServer(whoami)
	defer srv.Close()

	auth := New(srv.URL)

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
				r.AddCookie(&http.Cookie{Name: "test-auth-cookie", Value: "1"})
				if err != nil {
					t.Fatalf("Can not create request: %v", err)
				}
				return r
			},
			false,
			1,
			false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			whoami.anonymous = tt.anonymousEnabled

			uid, err := auth.whoami(tt.request())
			if tt.err {
				if err == nil {
					t.Errorf("Auth did not returned expected error")
				}
				return
			}

			if err != nil {
				t.Errorf("Auth returend unexpected err: %v", err)
			}

			if uid != tt.uid {
				t.Errorf("Auth returned uid %d, expected %d", uid, tt.uid)
			}
		})
	}
}

type WhoAmIMock struct {
	anonymous bool
}

func (wai *WhoAmIMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uid := "null"

	c, _ := r.Cookie("test-auth-cookie")

	if c != nil {
		uid = c.Value
	}

	anonymous := "false"
	if wai.anonymous {
		anonymous = "true"
	}

	fmt.Fprintf(w, `{"user_id":%s,"guest_enabled":%s,"user":null,"auth_type":"default","permissions":[]}`, uid, anonymous)
}
