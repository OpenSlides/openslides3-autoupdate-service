// Package datastore receives and caches the data from redis and the worker.
package datastore

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
)

const (
	groupDefaultPK = 1
	groupAdminPK   = 2
)

// Datastore holds the connection to OpenSlides and Redis.
type Datastore struct {
	osAddr      string
	redisConn   RedisConn
	cache       *cache
	minChangeID int

	mu          sync.RWMutex
	maxChangeID int
}

// New returns an initialized Datastore instance.
func New(osAddr string, redisConn RedisConn) (*Datastore, error) {
	// TODO: Handle case, that data is not ready.
	fd, min, max, err := redisConn.FullData()
	if err != nil {
		return nil, fmt.Errorf("get startdata from redis: %w", err)
	}

	d := &Datastore{
		osAddr:      osAddr,
		redisConn:   redisConn,
		cache:       &cache{data: fd},
		minChangeID: min,
		maxChangeID: max,
	}

	return d, nil
}

// LowestID returns the highest known change id.
func (d *Datastore) LowestID() int {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.maxChangeID
}

// KeysChanged blocks until there is new data. It updates the internal cache and
// returns the changed keys and the new change id.
func (d *Datastore) KeysChanged() ([]string, int, error) {
	var rawData []byte
	var err error
	for rawData == nil {
		// Update() blocks until there is new data. But when there is no new
		// data for an hour, then it returns with nil.
		rawData, err = d.redisConn.Update()
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

	changeID := sData.ChangeID

	if changeID > d.maxChangeID+1 {
		// Data is to new. Get the data in between.
		data, err := d.receive(d.maxChangeID, changeID-1)
		if err != nil {
			return nil, 0, fmt.Errorf("receive missing data from %d to %d: %w", d.maxChangeID, changeID-1, err)
		}
		d.updateCache(data, changeID-1)
	}

	if changeID < d.maxChangeID+1 {
		// Data already known. Try the next.
		return d.KeysChanged()
	}

	d.updateCache(sData.Elements, changeID)

	keys := make([]string, 0, len(sData.Elements))
	for k := range sData.Elements {
		keys = append(keys, k)
	}
	return keys, changeID, nil
}

// GetMany returns the values for the given keys.
func (d *Datastore) GetMany(keys []string) map[string]json.RawMessage {
	return d.cache.forKeys(keys...)
}

// GetAll returns all data.
func (d *Datastore) GetAll() map[string]json.RawMessage {
	return d.cache.all()
}

// HasPerm returns, if the user has the perm.
func (d *Datastore) HasPerm(uid int, perm string) bool {
	var group struct {
		Permissions []string `json:"permissions"`
	}

	if uid == 0 {
		// Check if the perm is in the default group.
		defaultGroup := d.cache.element("users/group:" + strconv.Itoa(groupDefaultPK))

		if err := json.Unmarshal(defaultGroup, &group); err != nil {
			return false
		}

		for _, p := range group.Permissions {
			if p == perm {
				return true
			}
		}
		return false
	}

	userData := d.cache.element("users/user" + strconv.Itoa(uid))

	if userData == nil {
		return false
	}

	var user struct {
		GroupsID []int `json:"groups_id"`
	}

	if err := json.Unmarshal(userData, &user); err != nil {
		return false
	}

	if len(user.GroupsID) == 0 {
		user.GroupsID = []int{groupDefaultPK}
	}

	for _, id := range user.GroupsID {
		if id == groupAdminPK {
			// User is in admin group.
			return true
		}

		group.Permissions = nil

		groupData := d.cache.element("users/group:" + strconv.Itoa(id))

		if groupData == nil {
			// User is in group that does not exist.
			return false
		}

		if err := json.Unmarshal(groupData, &group); err != nil {
			return false
		}

		for _, p := range group.Permissions {
			if p == perm {
				return true
			}
		}
	}
	return false
}

// update updates the cache. It is not save for concourent use.
func (d *Datastore) updateCache(data map[string]json.RawMessage, changeID int) {
	d.cache.update(data)

	d.mu.Lock()
	d.maxChangeID = changeID
	d.mu.Unlock()
}

// receive is used to get missing data. It returns all data between higher
// "from" and lower or equal "to".
func (d *Datastore) receive(from, to int) (data map[string]json.RawMessage, err error) {
	keys, err := d.redisConn.ChangedKeys(from, to)
	if err != nil {
		return nil, fmt.Errorf("get changed keys: %w", err)
	}

	data, err = d.redisConn.Data(keys)
	if err != nil {
		return nil, fmt.Errorf("get data: %w", err)
	}
	return data, nil
}
