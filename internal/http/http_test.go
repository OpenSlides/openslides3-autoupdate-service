package http_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/autoupdate"
	ahttp "github.com/OpenSlides/openslides3-autoupdate-service/internal/http"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

func TestAutoupdateFirstData(t *testing.T) {
	auther := new(test.AutherMock)
	receiver := new(test.ReceiverMock)
	receiver.MaxChangeID = 2
	receiver.FullData = map[string]json.RawMessage{
		"user:1": []byte("hello world1"),
		"user:2": []byte("hello world2"),
	}

	a, err := autoupdate.New(receiver)
	if err != nil {
		t.Fatalf("autoupdate startup failed: %v", err)
	}
	defer a.Close()

	srv := httptest.NewServer(ahttp.New(a, auther))
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

func TestAutoupdateWithChangeID(t *testing.T) {
	t.Skip("This does not work :(")
	auther := new(test.AutherMock)
	receiver := new(test.ReceiverMock)
	receiver.MaxChangeID = 2
	receiver.FullData = map[string]json.RawMessage{
		"user:1": []byte("hello world1"),
		"user:2": []byte("hello world2"),
	}

	a, err := autoupdate.New(receiver)
	if err != nil {
		t.Fatalf("autoupdate startup failed: %v", err)
	}
	defer a.Close()

	srv := httptest.NewServer(ahttp.New(a, auther))
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	req, err := http.NewRequest(http.MethodGet, srv.URL+"/system/autoupdate?change_id=3", nil)
	if err != nil {
		t.Fatalf("Can not create request: %v", err)
	}

	transport := &http.Transport{}
	client := &http.Client{Transport: transport}

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		t.Fatalf("Can not send request: %v", err)
	}
	defer resp.Body.Close()
	transport.CancelRequest(req)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Handler returned status %s, expected 200, %s", resp.Status, http.StatusText(200))
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatalf("Can not read body: %v", err)
	}

	if len(body) != 0 {
		t.Errorf("Got content in body, expected no data")
	}
}

func TestAuth(t *testing.T) {
	auther := new(test.AutherMock)
	receiver := new(test.ReceiverMock)

	a, err := autoupdate.New(receiver)
	if err != nil {
		t.Fatalf("autoupdate startup failed: %v", err)
	}
	defer a.Close()

	srv := httptest.NewServer(ahttp.New(a, auther))
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
