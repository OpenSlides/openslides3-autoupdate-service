package http_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/autoupdate"
	ahttp "github.com/OpenSlides/openslides3-autoupdate-service/internal/http"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

func TestAuth(t *testing.T) {
	auther := new(test.AutherMock)
	datastore := test.NewDatastoreMock(1)

	restricter := new(test.RestricterMock)

	closed := make(chan struct{})
	defer close(closed)
	a, err := autoupdate.New(datastore, restricter, closed)
	if err != nil {
		t.Fatalf("autoupdate startup failed: %v", err)
	}

	srv := httptest.NewUnstartedServer(ahttp.New(auther, ahttp.WithAutoupdate(a)))
	srv.EnableHTTP2 = true
	srv.StartTLS()
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL+"/system/autoupdate", nil)
	if err != nil {
		t.Fatalf("Can not create request: %v", err)
	}
	resp, err := srv.Client().Do(req)
	if err != nil {
		t.Fatalf("Can not send request: %v", err)
	}
	defer resp.Body.Close()
	cancel()

	if !auther.AuthCalled {
		t.Errorf("Auther was not called.")
	}
}
