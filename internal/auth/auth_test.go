package auth_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/auth"
)

func TestAuth(t *testing.T) {
	whoami := new(WhoAmIMock)
	srv := httptest.NewServer(whoami)
	defer srv.Close()

	auth := auth.New(srv.URL)

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
				r.Header.Set("x-test-user", "1")
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

			uid, err := auth.Auth(tt.request())
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
	uid := r.Header.Get("x-test-user")
	if uid == "" {
		uid = "null"
	}

	anonymous := "false"
	if wai.anonymous {
		anonymous = "true"
	}

	fmt.Fprintf(w, `{"user_id":%s,"guest_enabled":%s,"user":null,"auth_type":"default","permissions":[]}`, uid, anonymous)
}
