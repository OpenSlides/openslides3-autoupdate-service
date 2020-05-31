package autoupdate_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

func TestAutoupdateConstructor(t *testing.T) {
	receiver := new(test.ReceiverMock)

	a, err := autoupdate.New(receiver)
	if err != nil {
		t.Fatalf("autoupdate startup failed: %v", err)
	}
	defer a.Close()

	if !receiver.StartDataCalled {
		t.Errorf("Method StartData() on receiver not called")
	}
}

func TestAutoupdateReceiveNewData(t *testing.T) {
	receiver := new(test.ReceiverMock)
	receiver.MaxChangeID = 2
	receiver.FullData = map[string]json.RawMessage{
		"user:1": []byte("hello world1"),
		"user:2": []byte("hello world2"),
	}
	receiver.Changed = []string{"user:1"}

	a, err := autoupdate.New(receiver)
	if err != nil {
		t.Fatalf("autoupdate startup failed: %v", err)
	}
	defer a.Close()

	data, id, err := a.Receive(context.Background(), 1)

	if err != nil {
		t.Errorf("Receive returned an unexpected error: %v", err)
	}

	if id != 2 {
		t.Errorf("Receive returnd changeID %d, expected %d", id, 2)
	}

	if len(data) != 1 || string(data["user:1"]) != "hello world1" {
		t.Errorf("Receice returned user:1 = %s, expected `hello world1`", data["user:1"])
	}
}

func TestAutoupdateReceiveFirstData(t *testing.T) {
	receiver := new(test.ReceiverMock)
	receiver.MaxChangeID = 2
	receiver.FullData = map[string]json.RawMessage{
		"user:1": []byte("hello world1"),
		"user:2": []byte("hello world2"),
	}
	receiver.Changed = []string{"user:1"}

	a, err := autoupdate.New(receiver)
	if err != nil {
		t.Fatalf("autoupdate startup failed: %v", err)
	}
	defer a.Close()

	data, id, err := a.Receive(context.Background(), 0)

	if err != nil {
		t.Errorf("Receive returned an unexpected error: %v", err)
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
