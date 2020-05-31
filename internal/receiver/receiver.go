package receiver

import (
	"encoding/json"
	"fmt"
)

// Receiver holds the connection to OpenSlides and Redis.
type Receiver struct {
	osAddr    string
	redisConn RedisConn
}

// New returns an initialized Receiver.
func New(osAddr string, redisConn RedisConn) *Receiver {
	r := &Receiver{
		osAddr:    osAddr,
		redisConn: redisConn,
	}
	return r
}

// StartData reads the data that is needed for startup.
func (r *Receiver) StartData() (fullData map[string]json.RawMessage, minChangeID, maxChangeID int, err error) {
	fd, min, max, err := r.redisConn.FullData()
	if err != nil {
		return nil, 0, 0, fmt.Errorf("get startdata from redis: %w", err)
	}

	// TODO: Handle case, that data are not ready.

	return fd, min, max, nil
}

// Update blocks until there is new data. Then, it returns the data and the
// new changeID.
func (r *Receiver) Update() (data map[string]json.RawMessage, changeID int, err error) {
	var rawData []byte
	for rawData == nil {
		// Update() blocks until there is new data. But when there is no new
		// data for an hour, then it returns with nil.
		rawData, err = r.redisConn.Update()
		if err != nil {
			return nil, 0, fmt.Errorf("get autoupdate data: %w", err)
		}
	}

	var sData struct {
		Elements map[string]json.RawMessage `json:"elements"`
		ChangeID int                        `json:"change_id"`
	}

	if err := json.Unmarshal(rawData, &sData); err != nil {
		return nil, 0, fmt.Errorf("parse data from redis: %w", err)
	}

	return sData.Elements, sData.ChangeID, nil

}

// Receive is used to get missing data. It returns all data between higher
// "from" and lower or equal "to".
func (r *Receiver) Receive(from, to int) (data map[string]json.RawMessage, err error) {
	keys, err := r.redisConn.ChangedKeys(from, to)
	if err != nil {
		return nil, fmt.Errorf("get changed keys: %w", err)
	}

	data, err = r.redisConn.Data(keys)
	if err != nil {
		return nil, fmt.Errorf("get data: %w", err)
	}
	return data, nil
}

// ChangedKeys returns all keys that where changed between higher "from" and
// lower or equal "to".
func (r *Receiver) ChangedKeys(from, to int) (keys []string, err error) {
	return r.redisConn.ChangedKeys(from, to)
}
