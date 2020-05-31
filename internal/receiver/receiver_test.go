package receiver_test

import (
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/receiver"
)

func TestStartData(t *testing.T) {
	r := receiver.New(":8000", &redisConnMock{})

	fd, minID, maxID, err := r.StartData()

	if minID != 1 {
		t.Errorf("Got minID %d, expected 1", minID)
	}

	if maxID != 2 {
		t.Errorf("Got minID %d, expected 2", maxID)
	}

	if err != nil {
		t.Errorf("Got error %v, expected none", err)
	}

	if _, ok := fd["user:1"]; !ok {
		t.Errorf("Got map without expected key `user:1`")
	}

	if got := fd["user:1"]; string(got) != "hello world" {
		t.Errorf("Got value for key user:1 `%s`, expected `hello world`", got)
	}
}
