package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/auth"
	ahttp "github.com/OpenSlides/openslides3-autoupdate-service/internal/http"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

func TestAuth(t *testing.T) {
	auther := new(test.AutherMock)
	auther.UID = 99

	srv := httptest.NewUnstartedServer(ahttp.New(auther, ahttp.WithErrHandleFunc("/test", testHandler)))
	srv.EnableHTTP2 = true
	srv.StartTLS()
	defer srv.Close()

	resp, err := srv.Client().Get(srv.URL + "/test")
	if err != nil {
		t.Fatalf("Can not send request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code %s, expected 200 OK", resp.Status)
	}

	if !auther.AuthCalled {
		t.Errorf("Auther was not called.")
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) error {
	uid := r.Context().Value(auth.UserIDKey).(int)
	if uid != 99 {
		return fmt.Errorf("got uid %d, expected 99", uid)
	}
	return nil
}
