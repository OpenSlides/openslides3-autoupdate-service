package http_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/autoupdate"
	ahttp "github.com/OpenSlides/openslides3-autoupdate-service/internal/http"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

func TestAutoupdateFirstData(t *testing.T) {
	auther := new(test.AutherMock)
	datastore := test.NewDatastoreMock(1)
	datastore.FullData = map[string]json.RawMessage{
		"user:1": []byte("hello world1"),
		"user:2": []byte("hello world2"),
	}

	restricter := new(test.RestricterMock)

	a, err := autoupdate.New(datastore, restricter)
	if err != nil {
		t.Fatalf("autoupdate startup failed: %v", err)
	}
	defer a.Close()

	srv := httptest.NewServer(ahttp.New(a, auther, 0))
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL+"/system/autoupdate", nil)
	if err != nil {
		t.Fatalf("Can not create request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Can not send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Handler returned status %s, expected 200, %s", resp.Status, http.StatusText(200))
	}
}

func TestAuth(t *testing.T) {
	auther := new(test.AutherMock)
	datastore := test.NewDatastoreMock(1)

	restricter := new(test.RestricterMock)

	a, err := autoupdate.New(datastore, restricter)
	if err != nil {
		t.Fatalf("autoupdate startup failed: %v", err)
	}
	defer a.Close()

	srv := httptest.NewServer(ahttp.New(a, auther, 0))
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL+"/system/autoupdate", nil)
	if err != nil {
		t.Fatalf("Can not create request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Can not send request: %v", err)
	}
	defer resp.Body.Close()
	cancel()

	if !auther.AuthCalled {
		t.Errorf("Auther was not called.")
	}
}
