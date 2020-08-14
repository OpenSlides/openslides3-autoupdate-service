package test

//go:generate  sh -c "go run gen_db/main.go > db.json.go && go fmt db.json.go"
import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// DatastoreMock implements the autoupdate.Datastore interface.
type DatastoreMock struct {
	FullData    map[string]json.RawMessage
	minChangeID int
	maxChangeID int
	Err         error

	closed  <-chan struct{}
	changes chan []string
}

// NewDatastoreMock initializes a DatastoreMock.
func NewDatastoreMock(startID int, closed <-chan struct{}) *DatastoreMock {
	changes := make(chan []string, 1)

	fdCopy := make(map[string]json.RawMessage, len(exampleData))
	for k, v := range exampleData {
		fdCopy[k] = v
	}

	d := &DatastoreMock{
		FullData:    fdCopy,
		changes:     changes,
		minChangeID: startID,
		maxChangeID: startID,
		closed:      closed,
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
func (d *DatastoreMock) KeysChanged() ([]string, int, error) {
	var changes []string
	select {
	case changes = <-d.changes:
	case <-d.closed:
		return nil, 0, closingErr{}
	}
	d.maxChangeID++
	return changes, d.maxChangeID, nil
}

// ChangedKeys does nothing...
func (d *DatastoreMock) ChangedKeys(from, to int) ([]string, error) {
	return nil, nil
}

// Get sets v to the decoded value of collection:id
func (d *DatastoreMock) Get(collection string, id int, v interface{}) error {
	e, ok := d.FullData[fmt.Sprintf("%s:%d", collection, id)]
	if !ok {
		return doesNotExist(fmt.Sprintf("%s:%d", collection, id))
	}
	return json.Unmarshal(e, v)
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

// GetCollection gets all elements of one collection.
func (d *DatastoreMock) GetCollection(collection string) []json.RawMessage {
	var elements []json.RawMessage
	prefix := collection + ":"
	for key, value := range d.FullData {
		if !strings.HasPrefix(key, prefix) {
			continue
		}
		elements = append(elements, value)
	}
	return elements
}

// GetModels returns each element from collection that is in the ids slide.
func (d *DatastoreMock) GetModels(collection string, ids []int) []json.RawMessage {
	set := make(map[int]bool)
	for _, id := range ids {
		set[id] = true
	}

	var elements []json.RawMessage
	for key, value := range d.FullData {
		parts := strings.Split(key, ":")
		if len(parts) != 2 || parts[0] != collection {
			continue
		}

		id, err := strconv.Atoi(parts[1])
		if err != nil || !set[id] {
			continue
		}

		elements = append(elements, value)
	}
	return elements
}

// Change sends an update event for KeysChanged.
func (d *DatastoreMock) Change(keys []string) {
	d.changes <- keys
}

// Projectors does currently nothing.
func (d *DatastoreMock) Projectors(ctx context.Context, tid uint64) (uint64, map[int]json.RawMessage, error) {
	return 0, nil, nil
}

// ConfigValue sets v to the value of config value key.
func (d *DatastoreMock) ConfigValue(key string, v interface{}) error {
	for k, v := range d.FullData {
		parts := strings.Split(k, ":")
		if len(parts) != 2 {
			return fmt.Errorf("key %s has wrong format. Expected one `:`", k)
		}

		if parts[0] != "core/config" {
			continue
		}

		id, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("key %s has wrong format. Expected id to be int not %s", k, parts[1])
		}
		var decoded struct {
			Key   string          `json:"key"`
			Value json.RawMessage `json:"value"`
		}
		if err := json.Unmarshal(v, &decoded); err != nil {
			return fmt.Errorf("decoding core/config:%d: %w", id, err)
		}

		if decoded.Key != key {
			continue
		}

		return json.Unmarshal(decoded.Value, &v)
	}
	return nil
}

type doesNotExist string

func (e doesNotExist) Error() string {
	return fmt.Sprintf("%s does not exist", string(e))
}

func (e doesNotExist) DoesNotExist() {}
