package test

import "encoding/json"

// DatastoreMock implements the autoupdate.Datastore interface.
type DatastoreMock struct {
	FullData    map[string]json.RawMessage
	minChangeID int
	maxChangeID int
	Err         error

	changes chan []string
}

// NewDatastoreMock initializes a DatastoreMock
func NewDatastoreMock(startID int) *DatastoreMock {
	changes := make(chan []string, 1)
	d := &DatastoreMock{
		changes:     changes,
		minChangeID: startID,
		maxChangeID: startID,
	}
	return d
}

// LowestID returns the max id if the mock.
func (d *DatastoreMock) LowestID() int {
	return d.minChangeID
}

// CurrentID returns the max id if the mock.
func (d *DatastoreMock) CurrentID() int {
	return d.maxChangeID
}

// KeysChanged waits for Changes to be called.
func (d *DatastoreMock) KeysChanged(closing chan struct{}) ([]string, int, error) {
	var changes []string
	select {
	case changes = <-d.changes:
	case <-closing:
		return nil, 0, closingErr{}
	}
	d.maxChangeID++
	return changes, d.maxChangeID, nil
}

// ChangedKeys does nothing...
func (d *DatastoreMock) ChangedKeys(from, to int) ([]string, error) {
	return nil, nil
}

// GetMany returns the given keys from FullData.
func (d *DatastoreMock) GetMany(keys []string) map[string]json.RawMessage {
	data := make(map[string]json.RawMessage, len(keys))
	for _, k := range keys {
		data[k] = d.FullData[k]
	}
	return data
}

// GetAll returns FullData.
func (d *DatastoreMock) GetAll() map[string]json.RawMessage {
	return d.FullData
}

// Change sends an update event for KeysChanged.
func (d *DatastoreMock) Change(keys []string) {
	d.changes <- keys
}
