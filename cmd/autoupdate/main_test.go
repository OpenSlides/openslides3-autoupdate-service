package main

import (
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

func TestRestrictedData(t *testing.T) {
	for _, tt := range test.ExampleRestrictedData() {
		t.Run(tt.Name, func(t *testing.T) {
			restricters := openslidesRestricters(tt.Permer)
			r := restricters[tt.Collection]

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

func TestRequiredUser(t *testing.T) {
	required := openslidesRequiredUsers()

	for _, tt := range test.ExampleRequiredUser() {
		t.Run(tt.Name, func(t *testing.T) {
			r, ok := required[tt.Collection]
			if !ok {
				// No required users for element.
				return
			}

			ids, perm, err := r(tt.Element)
			if err != nil {
				t.Errorf("RequiredUser returned unexpected error: %v", err)
			}

			if perm != tt.ExpectPerm {
				t.Errorf("RequiredUser returned perm %s, expected %s", perm, tt.ExpectPerm)
			}

			if !test.CmpIntSlice(ids, tt.ExpectIDs) {
				t.Errorf("RequiredUser returned ids %v, expected %v", ids, tt.ExpectIDs)
			}
		})
	}
}

func TestProjector(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	callabes := openslidesProjectorCallables()
	ds := test.NewDatastoreMock(0, closed)

	todoList := map[string]bool{
		"motions/motion-poll": true,
	}

	for _, tt := range test.ExampleProjector() {
		t.Run(tt.Name, func(t *testing.T) {
			if todoList[tt.ElementName] {
				t.Skip()
			}

			c, ok := callabes[tt.ElementName]
			if !ok {
				t.Fatalf("No callable for Element `%s`", tt.ElementName)
			}

			got, err := c.Build(ds, tt.Element, 1)
			if err != nil {
				t.Errorf("ProjectorCallable returned unexpected error: %v", err)
			}

			test.ExpectEqualJSON(t, got, tt.Expected)
		})
	}
}
