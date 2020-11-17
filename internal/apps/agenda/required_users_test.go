package agenda_test

import (
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/apps/agenda"
)

func TestRequiredUsersListOfSpeakers(t *testing.T) {
	element := `{
        "id": 8,
        "title_information": {
            "title": "w1"
        },
        "speakers": [
            {
                "id": 3,
                "user_id": 1,
                "begin_time": null,
                "end_time": null,
                "weight": 1,
                "marked": false
            },
            {
                "id": 4,
                "user_id": 1,
                "begin_time": "2020-08-11T12:28:30.894574+02:00",
                "end_time": null,
                "weight": 1,
                "marked": false
            }
        ],
        "closed": false,
        "content_object": {
            "collection": "assignments/assignment",
            "id": 2
        }
	}`

	ids, perm, err := agenda.RequiredSpeakers([]byte(element))
	if err != nil {
		t.Errorf("RequiredSpeakers returned unexpected error: %v", err)
	}
	if len(ids) != 1 || !ids[1] {
		t.Errorf("RequiredSpeakers retunred %v, expected {1: true}", ids)
	}
	if perm != "agenda.can_see_list_of_speakers" {
		t.Errorf("RequiredSpeakers retuned perm %s, expected agenda.can_see_list_of_speakers", perm)
	}
}
