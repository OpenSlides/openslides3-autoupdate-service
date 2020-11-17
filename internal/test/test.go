// Package test contains helper functions to test the other packages.
package test

import (
	"encoding/json"
	"reflect"
	"testing"
)

// CmpIntSlice returns true, if both slices have the same values.
func CmpIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 && len(b) == 0 {
		// handle nil case
		return true
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// CmpIntSet returns true, if both sets have the smae entries
func CmpIntSet(a, b map[int]bool) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 && len(b) == 0 {
		// handle nil case
		return true
	}

	for key := range a {
		if !b[key] {
			return false
		}
	}

	for key := range b {
		if !a[key] {
			return false
		}
	}

	return true
}

// CmpStrSlice returns true, if both slices have the same values.
func CmpStrSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 && len(b) == 0 {
		// handle nil case
		return true
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type closingErr struct{}

func (e closingErr) Closing()      {}
func (e closingErr) Error() string { return "closing" }

// ExpectEqualJSON tests, that a and b are the same json elements.
func ExpectEqualJSON(t *testing.T, a, b []byte) {
	var aenc interface{}
	var benc interface{}
	if err := json.Unmarshal(a, &aenc); err != nil {
		t.Fatalf("a is invalid json: %v", err)
	}
	if err := json.Unmarshal(b, &benc); err != nil {
		t.Fatalf("b is invalid json: %v", err)
	}
	if !reflect.DeepEqual(aenc, benc) {
		ap, err := json.MarshalIndent(aenc, "", "  ")
		if err != nil {
			t.Fatalf("remarshalling a: %v", err)
		}
		bp, err := json.MarshalIndent(benc, "", "  ")
		if err != nil {
			t.Fatalf("remarshalling b: %v", err)
		}
		t.Errorf("json not equal: %s\n\n\n%s", ap, bp)
	}
}
