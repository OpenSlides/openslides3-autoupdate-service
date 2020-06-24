package datastore_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/datastore"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

func TestDatastoreLowestID(t *testing.T) {
	r := &test.RedisMock{
		Min: 5,
	}
	closing := make(chan struct{})
	defer close(closing)
	ds, err := datastore.New("Addr", r, nil, closing)
	if err != nil {
		t.Fatalf("Can not initialize datastore: %v", err)
	}

	if ds.LowestID() != 5 {
		t.Errorf("LowestID() returned %d, expected 5", ds.LowestID())
	}
}

func TestKeysChangedNormal(t *testing.T) {
	data := []byte(`{
		"change_id": 6,
		"elements":  {
			"elements/element:1": {"id": 1, "value": "hello world"}
		}
	}`)
	r := test.NewRedisMock()
	r.Max = 5

	closing := make(chan struct{})
	defer close(closing)
	ds, err := datastore.New("Addr", r, nil, closing)
	if err != nil {
		t.Fatalf("Can not initialize datastore: %v", err)
	}

	r.Send(data)
	keys, chID, err := ds.KeysChanged(closing)

	if err != nil {
		t.Errorf("KeysChanged returned unexpected err: %v", err)
	}

	if chID != 6 {
		t.Errorf("KeysChanged returned change_id %d, expected 6", chID)
	}

	if !test.CmpStrSlice(keys, []string{"elements/element:1"}) {
		t.Errorf("KeysChanged returned keys %v, expected [elements/element:1]", keys)
	}
}

func TestKeysChangedDelete(t *testing.T) {
	data := []byte(`{
		"change_id": 6,
		"elements": {
			"elements/element:1": null
		}
	}`)
	r := test.NewRedisMock()
	r.FD = map[string]json.RawMessage{
		"elements/element:1": []byte(`{"id": 1, "value": "hello world"}`),
	}
	r.Max = 5

	closing := make(chan struct{})
	defer close(closing)
	ds, err := datastore.New("Addr", r, nil, closing)
	if err != nil {
		t.Fatalf("Can not initialize datastore: %v", err)
	}

	r.Send(data)
	keys, chID, err := ds.KeysChanged(closing)

	if err != nil {
		t.Errorf("KeysChanged returned unexpected err: %v", err)
	}

	if chID != 6 {
		t.Errorf("KeysChanged returned change_id %d, expected 6", chID)
	}

	if !test.CmpStrSlice(keys, []string{"elements/element:1"}) {
		t.Errorf("KeysChanged returned keys %v, expected nil", keys)
	}

	if got := ds.GetAll(); len(got) != 0 {
		t.Errorf("Expected no more data in ds, got %v", got)
	}
}

func TestKeysChangedSkippedChangeID(t *testing.T) {
	data := []byte(`{
		"change_id": 10,
		"elements":  {
			"elements/element:1": {"id": 1, "value": "hello world"}
		}
	}`)
	r := test.NewRedisMock()
	r.FD = map[string]json.RawMessage{
		"elements/element:2": []byte(`{"id": 2}`),
	}
	r.Max = 5
	r.ChangedKeysResult = []string{"elements/element:2"}

	closing := make(chan struct{})
	defer close(closing)
	ds, err := datastore.New("Addr", r, nil, closing)
	if err != nil {
		t.Fatalf("Can not initialize datastore: %v", err)
	}

	r.Send(data)
	keys, chID, err := ds.KeysChanged(closing)

	if err != nil {
		t.Errorf("KeysChanged returned unexpected err: %v", err)
	}

	if chID != 10 {
		t.Errorf("KeysChanged returned change_id %d, expected 6", chID)
	}

	if !test.CmpStrSlice(keys, []string{"elements/element:1", "elements/element:2"}) {
		t.Errorf("KeysChanged returned keys %v, expected [elements/element:1]", keys)
	}
}

func TestKeysChangedBlocking(t *testing.T) {
	data := []byte(`{
		"change_id": 6,
		"elements":  {
			"elements/element:1": {"id": 1, "value": "hello world"}
		}
	}`)
	r := test.NewRedisMock()
	r.Max = 5

	closing := make(chan struct{})
	defer close(closing)
	ds, err := datastore.New("Addr", r, nil, closing)
	if err != nil {
		t.Fatalf("Can not initialize datastore: %v", err)
	}

	unblocked := make(chan struct{})
	go func() {
		_, _, _ = ds.KeysChanged(closing)
		close(unblocked)
	}()

	// Check, that KeysChanged does not return for a millisecond.
	timer := time.NewTimer(time.Millisecond)
	select {
	case <-unblocked:
		t.Fatalf("KeysChanged was done before data was send")
	case <-timer.C:
	}

	r.Send(data)

	timer.Reset(time.Millisecond)
	select {
	case <-unblocked:
	case <-timer.C:
		t.Errorf("KeysChanged was did not return after data was send")
	}
}

func TestKeysChangedNilDoesNotUnblock(t *testing.T) {
	r := test.NewRedisMock()
	r.Max = 5

	closing := make(chan struct{})
	defer close(closing)
	ds, err := datastore.New("Addr", r, nil, closing)
	if err != nil {
		t.Fatalf("Can not initialize datastore: %v", err)
	}

	unblocked := make(chan struct{})
	go func() {
		_, _, _ = ds.KeysChanged(closing)
		close(unblocked)
	}()

	// Check, that KeysChanged does not return for a millisecond.
	timer := time.NewTimer(time.Millisecond)
	select {
	case <-unblocked:
		t.Fatalf("KeysChanged was done before data was send")
	case <-timer.C:
	}

	r.Send(nil)

	timer.Reset(time.Millisecond)
	select {
	case <-unblocked:
		t.Errorf("KeysChanged was done after sending nil")
	case <-timer.C:
	}
}

func TestKeysChangedLowIDDoesNotUnblock(t *testing.T) {
	data := []byte(`{
		"change_id": 5,
		"elements":  {
			"elements/element:1": {"id": 1, "value": "hello world"}
		}
	}`)
	r := test.NewRedisMock()
	r.Max = 5

	closing := make(chan struct{})
	defer close(closing)
	ds, err := datastore.New("Addr", r, nil, closing)
	if err != nil {
		t.Fatalf("Can not initialize datastore: %v", err)
	}

	unblocked := make(chan struct{})
	go func() {
		_, _, _ = ds.KeysChanged(closing)
		close(unblocked)
	}()

	// Check, that KeysChanged does not return for a millisecond.
	timer := time.NewTimer(time.Millisecond)
	select {
	case <-unblocked:
		t.Fatalf("KeysChanged was done before data was send")
	case <-timer.C:
	}

	r.Send(data)

	// Check, that KeysChanged does not return for a millisecond.
	timer.Reset(time.Millisecond)
	select {
	case <-unblocked:
		t.Errorf("KeysChanged was done after sending the same changeid.")
	case <-timer.C:
	}
}

func TestGetMany(t *testing.T) {
	r := test.NewRedisMock()
	r.Max = 5
	r.FD = map[string]json.RawMessage{
		"elements/element:1": []byte(`{"id": 1}`),
		"elements/element:2": []byte(`{"id": 2}`),
		"elements/element:3": []byte(`{"id": 3}`),
	}

	closing := make(chan struct{})
	defer close(closing)
	ds, err := datastore.New("Addr", r, nil, closing)
	if err != nil {
		t.Fatalf("Can not initialize datastore: %v", err)
	}

	got := ds.GetMany([]string{"elements/element:1", "elements/element:2"})

	if len(got) != 2 || string(got["elements/element:1"]) != `{"id": 1}` || string(got["elements/element:2"]) != `{"id": 2}` {
		t.Errorf("GetMany returned %v, expected values from elements 1 and 2", got)
	}
}

func TestGetAll(t *testing.T) {
	r := test.NewRedisMock()
	r.Max = 5
	r.FD = map[string]json.RawMessage{
		"elements/element:1": []byte(`{"id": 1}`),
		"elements/element:2": []byte(`{"id": 2}`),
		"elements/element:3": []byte(`{"id": 3}`),
	}

	closing := make(chan struct{})
	defer close(closing)
	ds, err := datastore.New("Addr", r, nil, closing)
	if err != nil {
		t.Fatalf("Can not initialize datastore: %v", err)
	}

	got := ds.GetAll()

	if len(got) != 3 || string(got["elements/element:1"]) != `{"id": 1}` || string(got["elements/element:2"]) != `{"id": 2}` || string(got["elements/element:3"]) != `{"id": 3}` {
		t.Errorf("GetAll returned %v, expected values from elements 1, 2 and 3", got)
	}
}
