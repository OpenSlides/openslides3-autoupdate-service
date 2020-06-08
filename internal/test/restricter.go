package test

import (
	"encoding/json"
)

// RestricterMock implements the autoupdate.Restrict interface.
type RestricterMock struct {
	Called bool
}

// Restrict does nothing.
func (r *RestricterMock) Restrict(uid int, data map[string]json.RawMessage) {
	r.Called = true
}
