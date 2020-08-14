package mediafile_test

import (
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/mediafile"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

func TestRestrict(t *testing.T) {
	for _, tt := range test.ExampleRestrictedData("mediafiles/mediafile") {
		t.Run(tt.Name, func(t *testing.T) {
			r := mediafile.Restrict(tt.Permer)

			got, err := r.Restrict(tt.UID, tt.Element)
			if err != nil {
				t.Errorf("Restrict returned unexpected error: %v ", err)
			}

			// Full restriction.
			if tt.Expected == nil {
				if got != nil {
					t.Errorf("Restrict returned `%s`, expected nil", got)
				}
				return
			}

			if got == nil {
				t.Errorf("Restrict() returned nil, expected %s", tt.Expected)
				return
			}

			test.ExpectEqualJSON(t, got, tt.Expected)
		})
	}
}
