package test

import "encoding/json"

// ReceiverMock implements the Receiver interface.
type ReceiverMock struct {
	FullData        map[string]json.RawMessage
	MinChangeID     int
	MaxChangeID     int
	Err             error
	StartDataCalled bool
	Changed         []string
}

// StartData returns the frist data.
func (r *ReceiverMock) StartData() (map[string]json.RawMessage, int, int, error) {
	r.StartDataCalled = true
	return r.FullData, r.MinChangeID, r.MaxChangeID, r.Err
}

// Update blocks forever.
func (r *ReceiverMock) Update() (map[string]json.RawMessage, int, error) {
	select {}
}

// Receive returns nothing.
func (r *ReceiverMock) Receive(from, to int) (map[string]json.RawMessage, error) {
	return nil, nil
}

// ChangedKeys returns the value in r.Changed.
func (r *ReceiverMock) ChangedKeys(from, to int) (keys []string, all bool, err error) {
	if from == 0 || r.Changed == nil {
		return nil, true, nil
	}
	return r.Changed, false, nil
}
