package user_test

import (
	"fmt"
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/user"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

func TestSlide(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	db := test.NewDatastoreMock(0, closed)
	slide := user.Slide()

	for _, tt := range []struct {
		uid    int
		expect string
	}{
		{1, `{"user":"Administrator"}`},
		{2, `{"user":"candidate1"}`},
	} {
		t.Run(fmt.Sprintf("user %d", tt.uid), func(t *testing.T) {
			got, err := slide(db, []byte(fmt.Sprintf(`{"id":%d}`, tt.uid)), 0)
			if err != nil {
				t.Errorf("slide returned an unexpected error: %v", err)
			}

			if string(got) != tt.expect {
				t.Errorf("slide returned %s, expected %s", got, tt.expect)
			}
		})
	}
}
