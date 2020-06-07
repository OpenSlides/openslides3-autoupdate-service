package autoupdate_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

func TestAutoupdateReceiveNewData(t *testing.T) {
	datastore := test.NewDatastoreMock(1)
	datastore.FullData = map[string]json.RawMessage{
		"user:1": []byte("hello world1"),
		"user:2": []byte("hello world2"),
	}
	datastore.Change([]string{"user:1"})

	a, err := autoupdate.New(datastore)
	if err != nil {
		t.Fatalf("autoupdate startup failed: %v", err)
	}
	defer a.Close()

	all, data, id, err := a.Receive(context.Background(), 1, 1)

	if err != nil {
		t.Errorf("Receive returned an unexpected error: %v", err)
	}

	if all {
		t.Error("Receive returned all == true, expected false")
	}

	if id != 2 {
		t.Errorf("Receive returnd changeID %d, expected %d", id, 2)
	}

	if len(data) != 1 {
		t.Errorf("Receive returned %d elements, expected 1. Got %v", len(data), data)
	}

	if string(data["user:1"]) != "hello world1" {
		t.Errorf("Receice returned for user:1: `%s`, expected `hello world1`", data["user:1"])
	}
}

func TestAutoupdateReceiveFirstData(t *testing.T) {
	datastore := test.NewDatastoreMock(2)
	datastore.FullData = map[string]json.RawMessage{
		"user:1": []byte("hello world1"),
		"user:2": []byte("hello world2"),
	}

	a, err := autoupdate.New(datastore)
	if err != nil {
		t.Fatalf("autoupdate startup failed: %v", err)
	}
	defer a.Close()

	all, data, id, err := a.Receive(context.Background(), 1, 0)

	if err != nil {
		t.Errorf("Receive returned an unexpected error: %v", err)
	}

	if !all {
		t.Error("Receive returned all == false, expected true")
	}

	if id != 2 {
		t.Errorf("Receive returnd changeID %d, expected %d", id, 2)
	}

	if len(data) != 2 {
		t.Errorf("Receive returned %d values, expected 2", len(data))
	}

	if string(data["user:2"]) != "hello world2" {
		t.Errorf("Receice returned user:2 = `%s`, expected `hello world2`", data["user:2"])
	}
}
