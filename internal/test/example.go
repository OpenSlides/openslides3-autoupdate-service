package test

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ExampleData returns a full database with example data.
func ExampleData() map[string]json.RawMessage {
	fdCopy := make(map[string]json.RawMessage, len(exampleData))
	for k, v := range exampleData {
		fdCopy[k] = v
	}
	return fdCopy
}

// RestrictedDataset is a used for a test.
type RestrictedDataset struct {
	Name       string
	UID        int
	Permer     *HasPermMock
	Collection string
	Element    json.RawMessage
	Expected   json.RawMessage
}

// ExampleRestrictedData returns a list of Data for restricted data tests for an user.
func ExampleRestrictedData() []RestrictedDataset {
	permers := permsForUser()
	var dsl []RestrictedDataset
	for uid, restrictedData := range exampleRestrictedData {
		for elementID, elment := range exampleData {
			collection := elementID[:strings.Index(elementID, ":")]
			ds := RestrictedDataset{
				Name:       fmt.Sprintf("User%d-Element-%s", uid, elementID),
				UID:        uid,
				Collection: collection,
				Permer:     permers[uid],
				Element:    elment,
				Expected:   restrictedData[elementID],
			}
			dsl = append(dsl, ds)
		}
	}
	return dsl
}

// RequiredUserDataset is a testcase.
type RequiredUserDataset struct {
	Name       string
	Collection string
	Element    json.RawMessage
	ExpectPerm string
	ExpectIDs  []int
}

// ExampleRequiredUser returns testcases to test required users.
func ExampleRequiredUser() []RequiredUserDataset {
	var rus []RequiredUserDataset
	for elementID, data := range exampleRequiredUser {
		collection := elementID[:strings.Index(elementID, ":")]
		ru := RequiredUserDataset{
			Name:       elementID,
			Collection: collection,
			Element:    exampleData[elementID],
			ExpectPerm: data.perm,
			ExpectIDs:  data.ids,
		}
		rus = append(rus, ru)
	}
	return rus
}

// ProjectorDataset is a test case.
type ProjectorDataset struct {
	Name        string
	ElementName string
	Overwrite   map[string]json.RawMessage
	Expected    json.RawMessage
}

// ExampleProjector returns test cases.
func ExampleProjector() []ProjectorDataset {
	var pds []ProjectorDataset
	for i, projector := range exampleProjector {
		var data struct {
			Element struct {
				Name string `json:"name"`
			} `json:"element"`
		}
		if err := json.Unmarshal(projector.Data, &data); err != nil {
			panic(err)
		}

		name := data.Element.Name
		if name == "" {
			name = "NONE"
		}

		pd := ProjectorDataset{
			Name:        fmt.Sprintf("Dataset%d-%s", i, name),
			ElementName: name,
			Overwrite:   projector.Overwrite,
			Expected:    projector.Data,
		}
		pds = append(pds, pd)
	}
	return pds
}

func permsForUser() map[int]*HasPermMock {
	r := make(map[int]*HasPermMock)
Outer:
	for uid := range exampleRestrictedData {
		var user struct {
			GroupsID []int `json:"groups_id"`
		}
		if err := json.Unmarshal(exampleData[fmt.Sprintf("users/user:%d", uid)], &user); err != nil {
			panic(err)
		}

		if len(user.GroupsID) == 0 {
			user.GroupsID = append(user.GroupsID, 1)
		}

		var perms []string
		groupSet := make(map[int]bool, len(user.GroupsID))
		for _, gid := range user.GroupsID {
			if gid == 2 {
				r[uid] = &HasPermMock{IsSuperuser: true}
				continue Outer
			}

			var group struct {
				Permissions []string `json:"permissions"`
			}
			if err := json.Unmarshal(exampleData[fmt.Sprintf("users/group:%d", gid)], &group); err != nil {
				panic(err)
			}
			perms = append(perms, group.Permissions...)
			groupSet[gid] = true
		}
		r[uid] = &HasPermMock{Perms: perms, Groups: groupSet}
	}
	return r
}
