package datastore

import (
	"encoding/json"
	"testing"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/test"
)

func TestHasPerm(t *testing.T) {
	for _, tt := range []struct {
		name      string
		userGroup []int
		groupPerm map[int]map[string]bool
		checkPerm string
		expect    bool
	}{
		{
			"normal",
			[]int{3},
			map[int]map[string]bool{3: {"my.perm": true}},
			"my.perm",
			true,
		},
		{
			"denied",
			[]int{3},
			map[int]map[string]bool{3: {"other.perm": true}},
			"my.perm",
			false,
		},
		{
			"isAdmin",
			[]int{2},
			map[int]map[string]bool{},
			"my.perm",
			true,
		},
		{
			"secondGroup",
			[]int{3, 4},
			map[int]map[string]bool{3: {}, 4: {"my.perm": true}},
			"my.perm",
			true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			hp := &hasPerm{
				userGroup: map[int][]int{1: tt.userGroup},
				groupPerm: tt.groupPerm,
			}

			if got := hp.HasPerm(1, tt.checkPerm); got != tt.expect {
				t.Errorf("HasPerm(1, %s) returned %t, expected %t", tt.checkPerm, got, tt.expect)
			}
		})
	}

	t.Run("anonymous", func(t *testing.T) {
		hp := &hasPerm{
			groupPerm: map[int]map[string]bool{
				1: {"my.perm": true},
				3: {"other.perm": true},
			},
		}

		if !hp.HasPerm(0, "my.perm") {
			t.Errorf("HasPerm(0, my.perm) returned false, expected true")
		}

		if hp.HasPerm(0, "other.perm") {
			t.Errorf("HasPerm(0, other.perm) returned true, expected false")
		}
	})
}

func TestInGroups(t *testing.T) {
	for _, tt := range []struct {
		name        string
		userGroup   []int
		checkGroups []int
		expect      bool
	}{
		{
			"check one",
			[]int{3, 4, 5},
			[]int{3},
			true,
		},
		{
			"check many",
			[]int{3, 4, 5},
			[]int{3, 6, 7},
			true,
		},
		{
			"not int groups",
			[]int{3, 4, 5},
			[]int{6, 7, 8},
			false,
		},
		{
			"admin group",
			[]int{2},
			[]int{3},
			true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			hp := &hasPerm{
				userGroup: map[int][]int{1: tt.userGroup},
			}

			if got := hp.InGroups(1, tt.checkGroups); got != tt.expect {
				t.Errorf("InGroup(1, %v) returned %t, expected %t", tt.checkGroups, got, tt.expect)
			}
		})
	}
}

func TestIsSuperadmin(t *testing.T) {
	hp := &hasPerm{
		userGroup: map[int][]int{
			1: {2, 3, 4},
			2: {3, 4, 5},
		},
	}

	if !hp.IsSuperadmin(1) {
		t.Errorf("IsSuperadmin(1) returned false, expected true")
	}

	if hp.IsSuperadmin(2) {
		t.Errorf("IsSuperadmin(2) returned true, expected false")
	}
}

func TestHasPermUpdate(t *testing.T) {
	hp := hasPerm{}

	err := hp.update(map[string]json.RawMessage{
		"users/user:1": []byte(`
			{
				"id": 1,
				"groups_id": [1, 2, 3],
				"other_field": "something"
			}
		`),
		"users/user:2": []byte(`
			{
				"id": 2,
				"groups_id": []
			}
		`),
		"users/group:1": []byte(`
			{
				"id": 1,
				"permissions": ["has_one", "has_two"]
			}
		`),
	})

	if err != nil {
		t.Fatalf("update returned unexpected err: %v", err)
	}

	if !test.CmpIntSlice(hp.userGroup[1], []int{1, 2, 3}) {
		t.Errorf("hp.userGroup[1] == %v, expected [1 2 3]", hp.userGroup[1])
	}

	if !test.CmpIntSlice(hp.userGroup[2], []int{1}) {
		t.Errorf("hp.userGroup[2] == %v, expected [1]", hp.userGroup[2])
	}

	if v := hp.groupPerm[1]; len(v) != 2 || !v["has_one"] || !v["has_two"] {
		t.Errorf("hp.groupPerm[1] == %v, expected map[has_one has_two]", hp.groupPerm[1])
	}
}

func TestHasPermUpdateDeletedUserAndGroup(t *testing.T) {
	hp := &hasPerm{
		userGroup: map[int][]int{
			1: {3},
		},
		groupPerm: map[int]map[string]bool{
			3: {"perm1": true},
		},
	}

	err := hp.update(map[string]json.RawMessage{
		"users/user:1":  []byte(`null`),
		"users/group:3": []byte(`null`),
	})

	if err != nil {
		t.Fatalf("update returned unexpected err: %v", err)
	}

	if hp.userGroup[1] != nil {
		t.Errorf("hp.userGroup[1] == %v, expected nil", hp.userGroup[1])
	}

	if hp.groupPerm[3] != nil {
		t.Errorf("hp.userGroup[3] == %v, expected nil", hp.userGroup[1])
	}
}
