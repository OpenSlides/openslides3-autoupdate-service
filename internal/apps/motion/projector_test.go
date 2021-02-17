package motion_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/motion"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

func TestConditionError(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	ds := test.NewDatastoreMock(1, done)
	ds.FullData = map[string]json.RawMessage{
		"motions/motion:1": []byte(`{"statute_paragraph_id":1}`),
	}

	element := json.RawMessage(`{"id":1}`)

	_, err := motion.Slide().Build(ds, element, 1)
	if err == nil {
		t.Fatalf("Expected an error")
	}

	var conErr interface {
		Conditions() []string
		ConditionError() error
	}
	if !errors.As(err, &conErr) {
		t.Fatalf("Expected an conditio error, got: %v", err)
	}

	consMap := make(map[string]bool)
	for _, c := range conErr.Conditions() {
		consMap[c] = true
	}

	expect := "motion.statute_paragraph_id == `1`"
	if !consMap[expect] {
		t.Errorf("expected condition %s", expect)
	}
}
