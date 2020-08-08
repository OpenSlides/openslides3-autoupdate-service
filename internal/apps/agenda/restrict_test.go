package agenda_test

import (
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/agenda"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

const (
	normalItem = `{
		"id": 1,
		"item_number": "",
		"is_hidden": false,
		"duration": null,
		"content_object": {
			"collection": "assignments/assignment",
			"id": 2
		},
		"tags_id": [],
		"is_internal": false,
		"closed": false,
		"weight": 10000,
		"level": 0,
		"type": 2,
		"title_information": {
			"title": "w1"
		},
		"parent_id": null
	}`
	normalItemOnlyReadPerm = `{
		"id":1,
		"item_number":"",
		"is_hidden":false,
		"content_object":{
			"collection":"assignments/assignment",
			"id":2
		},
		"tags_id":[],
		"is_internal":false,
		"closed":false,
		"weight":10000,
		"level":0,
		"type":2,
		"title_information":{
			"title":"w1"
		},
		"parent_id":null
	}`
)

func TestRestrict(t *testing.T) {
	permer := new(test.HasPermMock)
	r := agenda.Restrict(permer)

	for _, tt := range []struct {
		name     string
		perms    []string
		fullData string
		expected string
	}{
		{
			"No read permission",
			[]string{},
			normalItem,
			"",
		},
		{
			"Read permission",
			[]string{"agenda.can_see"},
			normalItem,
			normalItemOnlyReadPerm,
		},
		{
			"Manager",
			[]string{"agenda.can_see", "agenda.can_see_internal_items", "agenda.can_manage"},
			normalItem,
			normalItem,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			permer.Perms = tt.perms

			got, err := r.Restrict(1, []byte(tt.fullData))
			if err != nil {
				t.Errorf("Restrict returned unexpected error: %v ", err)
			}

			// Full restriction.
			if tt.expected == "" {
				if got != nil {
					t.Errorf("Restrict returned `%s`, expected nil", got)
				}
				return
			}

			if got == nil {
				t.Errorf("Restrict() returned nil, expected %s", tt.expected)
				return
			}

			test.ExpectEqualJSON(t, got, []byte(tt.expected))
		})
	}
}
