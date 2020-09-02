package autoupdate_test

import (
	"bytes"
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
	closed := make(chan struct{})
	defer close(closed)

	auther := new(test.AutherMock)
	datastore := test.NewDatastoreMock(1, closed)
	datastore.FullData = map[string]json.RawMessage{
		"user:1": []byte(`"hello world1"`),
		"user:2": []byte(`"hello world2"`),
	}

	restricter := new(test.RestricterMock)

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

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			body = []byte("no body")
		}
		body = bytes.TrimSpace(body)
		t.Errorf("Handler returned status %s: `%s`, expected 200, %s", resp.Status, body, http.StatusText(200))
	}
}
