// Package test contains helper functions to test the other packages.
package test

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
