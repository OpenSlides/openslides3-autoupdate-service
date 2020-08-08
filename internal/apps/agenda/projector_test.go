package agenda_test

import (
	"encoding/json"
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/agenda"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

const (
	i1 = `{
        "id": 1,
        "item_number": "",
        "title_information": {
            "title": "w1"
        },
        "comment": null,
        "closed": false,
        "type": 1,
        "is_internal": true,
        "is_hidden": false,
        "duration": null,
        "content_object": {
            "collection": "assignments/assignment",
            "id": 2
        },
        "weight": 10000,
        "parent_id": null,
        "level": 0,
        "tags_id": []
	}`
	i2 = `{
        "id": 2,
        "item_number": "",
        "title_information": {
            "title": "Antrag1",
            "identifier": "1"
        },
        "comment": null,
        "closed": false,
        "type": 1,
        "is_internal": true,
        "is_hidden": false,
        "duration": null,
        "content_object": {
            "collection": "motions/motion",
            "id": 1
        },
        "weight": 10000,
        "parent_id": null,
        "level": 0,
        "tags_id": []
	}`
	i3 = `{
        "id": 3,
        "item_number": "2",
        "title_information": {
            "title": "Neues3"
        },
        "comment": null,
        "closed": true,
        "type": 1,
        "is_internal": false,
        "is_hidden": false,
        "duration": null,
        "content_object": {
            "collection": "topics/topic",
            "id": 5
        },
        "weight": 10000,
        "parent_id": null,
        "level": 0,
        "tags_id": []
    }`
)

func TestProjectorItemListSlide(t *testing.T) {
	ds := new(test.DatastoreMock)
	ds.FullData = map[string]json.RawMessage{
		"agenda/item:1": []byte(i1),
		"agenda/item:2": []byte(i2),
		"agenda/item:3": []byte(i3),
	}

	element := []byte(`{"name": "agenda/item-list","only_main_items": false}`)

	got, err := agenda.ItemListSlide().Build(ds, element, 1)
	if err != nil {
		t.Errorf("Build returned unexpected error: %v", err)
	}

	expected := `{
		"items": [
		  {
			"title_information": {
			  "_agenda_item_number": "",
			  "title": "w1"
			},
			"collection": "assignments/assignment",
			"depth": 0
		  },
		  {
			"title_information": {
			  "_agenda_item_number": "",
			  "identifier": "1",
			  "title": "Antrag1"
			},
			"collection": "motions/motion",
			"depth": 0
		  },
		  {
			"title_information": {
			  "_agenda_item_number": "2",
			  "title": "Neues3"
			},
			"collection": "topics/topic",
			"depth": 0
		  }
		]
	  }`

	test.ExpectEqualJSON(t, got, []byte(expected))
}
