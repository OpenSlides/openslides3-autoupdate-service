package redis

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	// fullDataKey is the name of the redis key where the full_data are.
	fullDataKey = "full_data"

	// changeIDKey is the name of the redis key with the change ids.
	changeIDKey = "change_id"

	// lowestChangeIDField is the name of the field in the changeIDKey that
	// tells the lowerst change id in redis.
	lowestChangeIDField = "_config:lowest_change_id"

	// autoupdateKey is the keyname of the redis streams where the worker inform
	// about changed data.
	autoupdateKey = "autoupdate"

	// blockTimeout is the time in miliseconds, how long the xread command will
	// block.
	autoupdateBlockTimeout = "3600000" // One Hour
)

// Redis holds the connection to redis.
type Redis struct {
	pool   *redis.Pool
	lastID string
}

// New create a new Redis instance.
func New(addr string) *Redis {
	pool := &redis.Pool{
		MaxActive:   100,
		Wait:        true,
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}

	r := &Redis{
		pool: pool,
	}
	return r
}

// TestConn sends a ping command to redis. Does not return the response, but an
// error if there is no response.
func (r *Redis) TestConn() error {
	conn := r.pool.Get()
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		return fmt.Errorf("no connection to redis: %w", err)
	}
	return nil
}

// FullData gets all data from redis. It also gets the min and max change id in
// a atomic way.
func (r *Redis) FullData() (data map[string]json.RawMessage, max int, min int, err error) {
	conn := r.pool.Get()
	defer conn.Close()

	conn.Send("MULTI")
	conn.Send("HGETALL", fullDataKey)
	conn.Send("ZREVRANGEBYSCORE", changeIDKey, "+inf", "-inf", "WITHSCORES", "LIMIT", "0", "1")
	conn.Send("ZSCORE", changeIDKey, lowestChangeIDField)
	resp, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		return nil, 0, 0, fmt.Errorf("executing multi redis commands: %w", err)
	}

	if len(resp) != 3 {
		return nil, 0, 0, fmt.Errorf("invalid number of multi responce. Got %d, expected 3", len(resp))
	}

	rawData, err := redis.StringMap(resp[0], nil)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("get full_data from redis: %w", err)
	}

	data = make(map[string]json.RawMessage, len(rawData))
	for k, v := range rawData {
		data[k] = json.RawMessage(v)
	}

	maxChangeIDResp, err := redis.Strings(resp[1], nil)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("get max change id: %w", err)
	}

	if len(maxChangeIDResp) != 2 {
		return nil, 0, 0, fmt.Errorf("invalid values in max change id responce. Got %d, expected 2", len(maxChangeIDResp))
	}

	maxChangeID, err := strconv.Atoi(maxChangeIDResp[1])
	if err != nil {
		return nil, 0, 0, fmt.Errorf("invalid value in max change id responce, got %s, expected int", maxChangeIDResp[1])
	}

	minChangeID, err := redis.Int(resp[2], nil)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("get min change id: %w", err)
	}

	return data, maxChangeID, minChangeID, nil
}

// Update returns changed keys.
//
// Blocks until there is new data.
func (r *Redis) Update(closing chan struct{}) ([]byte, error) {
	conn := r.pool.Get()
	defer conn.Close()

	id := r.lastID
	if id == "" {
		id = "$"
	}

	var data []byte
	var err error
	received := make(chan struct{})

	go func() {
		id, data, err = stream(conn.Do("XREAD", "COUNT", 1, "BLOCK", autoupdateBlockTimeout, "STREAMS", autoupdateKey, id))
		close(received)
	}()

	select {
	case <-received:
	case <-closing:
		return nil, closingErr{}
	}

	if id != "" {
		r.lastID = id
	}

	if err != nil {
		if err == errNil {
			// No new data
			return nil, nil
		}

		return nil, fmt.Errorf("read autoupdate from redis: %w", err)
	}

	return data, nil
}

// ChangedKeys returns all keys in the changeidkey higher from and lower or
// equal to.
func (r *Redis) ChangedKeys(from, to int) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("ZRANGEBYSCORE", changeIDKey, "("+strconv.Itoa(from), strconv.Itoa(to)))
	if err != nil {
		return nil, fmt.Errorf("get changed keys from redis: %w", err)
	}
	return keys, nil
}

// Data returns the data from redis for specific keys from the full data hash
// key.
//
// If a key does not exist, the value in the returned dict is nil.
func (r *Redis) Data(keys []string) (map[string]json.RawMessage, error) {
	conn := r.pool.Get()
	defer conn.Close()

	args := make([]interface{}, len(keys)+1)
	args[0] = fullDataKey
	for i := range keys {
		args[i+1] = keys[i]
	}

	rawData, err := redis.ByteSlices(conn.Do("HMGET", args...))
	if err != nil {
		return nil, fmt.Errorf("hmget %s %v request: %w", fullDataKey, keys, err)
	}

	data := make(map[string]json.RawMessage, len(rawData))
	for i := range rawData {
		data[keys[i]] = json.RawMessage(rawData[i])
	}
	return data, nil
}

type closingErr struct{}

func (e closingErr) Closing()      {}
func (e closingErr) Error() string { return "closing" }
