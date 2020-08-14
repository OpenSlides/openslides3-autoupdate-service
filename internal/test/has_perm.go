package test

// HasPermMock implements the restricter.HasPermer interface.
type HasPermMock struct {
	Perms []string
}

// HasPerm returns true, if the given perm is in the list of Perms.
func (h *HasPermMock) HasPerm(_ int, perm string) bool {
	for _, p := range h.Perms {
		if p == perm {
			return true
		}
	}
	return false
}
