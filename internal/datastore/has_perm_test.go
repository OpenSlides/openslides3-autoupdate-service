package datastore

import "testing"

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
