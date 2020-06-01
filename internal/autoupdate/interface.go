package autoupdate

import (
	"encoding/json"
)

// Receiver gets the data for the autoupdate service.
type Receiver interface {
	// StartData reads the data that is needed for startup.
	StartData() (fullData map[string]json.RawMessage, minChangeID, maxChangeID int, err error)

	// Update blocks until there is new data. Then, it returns the data and the
	// new changeID.
	Update() (data map[string]json.RawMessage, changeID int, err error)

	// Receive is used to get missing data. It returns all data between higher
	// "from" and lower or equal "to".
	Receive(from, to int) (data map[string]json.RawMessage, err error)

	// Changed keys returns all keys that where changed between higher "from" and lower or equal "to".
	ChangedKeys(from, to int) (keys []string, err error)
}
