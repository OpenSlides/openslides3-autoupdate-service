package redis

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
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

	// cacheReayKey is the name of the redis key that tells, that the cache is
	// ready to use.
	cacheReadyKey = "cache_ready"

	// readyWait is the duration to wait to check again if the cache is ready.
	readyWait = 3 * time.Second

	// notifyKey is the name of the notify stream name.
	notifyKey = "notify"

	// applauseKey is the name of the redis key for applause.
	applauseKey = "applause"
)

// Redis holds the connection to redis.
type Redis struct {
	readPool         *redis.Pool
	writePool        *redis.Pool
	lastAutoupdateID string
	lastNotifyID     string
}

// New create a new Redis instance.
func New(readAddr, writeAddr string) *Redis {
	readPool := &redis.Pool{
		MaxActive:   100,
		Wait:        true,
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", readAddr) },
	}

	writePool := readPool
	if readAddr != writeAddr {
		writePool = &redis.Pool{
			MaxActive:   100,
			Wait:        true,
			MaxIdle:     10,
			IdleTimeout: 240 * time.Second,
			Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", writeAddr) },
		}
	}

	r := &Redis{
		readPool:  readPool,
		writePool: writePool,
	}
	return r
}

// TestReadConn sends a ping command to redis. Does not return the response, but an
// error if there is no response.
func (r *Redis) TestReadConn() error {
	return r.testConn(r.readPool)
}

// TestWriteConn sends a ping command to redis. Does not return the response, but an
// error if there is no response.
func (r *Redis) TestWriteConn() error {
	return r.testConn(r.writePool)
}

func (r *Redis) testConn(pool *redis.Pool) error {
	conn := pool.Get()
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		return fmt.Errorf("no connection to redis: %w", err)
	}
	return nil
}

// FullData gets all data from redis. It also gets the min and max change id in
// a atomic way.
func (r *Redis) FullData() (data map[string]json.RawMessage, max int, min int, err error) {
	conn := r.readPool.Get()
	defer conn.Close()

	var ready string
	for {
		ready, err = redis.String(conn.Do("GET", cacheReadyKey))
		if err != nil && err != redis.ErrNil {
			return nil, 0, 0, fmt.Errorf("get ready: %w", err)
		}

		if ready == "" {
			log.Printf("Cache not ready. Try again in %d seconds", readyWait/time.Second)
			time.Sleep(readyWait)
			continue
		}

		if ready == "ok" {
			break
		}

		return nil, 0, 0, fmt.Errorf("redis command get ready returned `%s`", ready)
	}

	if err := conn.Send("MULTI"); err != nil {
		return nil, 0, 0, fmt.Errorf("send MULTI to redis: %w", err)
	}

	if err := conn.Send("HGETALL", fullDataKey); err != nil {
		return nil, 0, 0, fmt.Errorf("send HGETALL to redis: %w", err)
	}

	if err := conn.Send("ZREVRANGEBYSCORE", changeIDKey, "+inf", "-inf", "WITHSCORES", "LIMIT", "0", "1"); err != nil {
		return nil, 0, 0, fmt.Errorf("send ZREVRANGEBYSCORE to redis: %w", err)
	}

	if err := conn.Send("ZSCORE", changeIDKey, lowestChangeIDField); err != nil {
		return nil, 0, 0, fmt.Errorf("send ZSCORE to redis: %w", err)
	}

	resp, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		return nil, 0, 0, fmt.Errorf("executing multi redis commands: %w", err)
	}

	if len(resp) != 3 {
		return nil, 0, 0, fmt.Errorf("invalid number of multi response. Got %d, expected 3", len(resp))
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
		return nil, 0, 0, fmt.Errorf("invalid values in max change id response, got %d, expected 2", len(maxChangeIDResp))
	}

	maxChangeID, err := strconv.Atoi(maxChangeIDResp[1])
	if err != nil {
		return nil, 0, 0, fmt.Errorf("invalid value in max change id response, got %s, expected int", maxChangeIDResp[1])
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
func (r *Redis) Update(closing <-chan struct{}) ([]byte, error) {
	conn := r.readPool.Get()
	defer conn.Close()

	id := r.lastAutoupdateID
	if id == "" {
		id = "$"
	}

	var data []byte
	var err error
	received := make(chan struct{})

	go func() {
		id, data, err = stream(conn.Do("XREAD", "COUNT", 1, "BLOCK", "0", "STREAMS", autoupdateKey, id))
		close(received)
	}()

	select {
	case <-received:
	case <-closing:
		return nil, closingError{}
	}

	if id != "" {
		r.lastAutoupdateID = id
	}

	if err != nil {
		return nil, fmt.Errorf("read autoupdate from redis: %w", err)
	}

	return data, nil
}

// ChangedKeys returns all keys in the changeidkey higher from and lower or
// equal to.
func (r *Redis) ChangedKeys(from, to int) ([]string, error) {
	conn := r.readPool.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("ZRANGEBYSCORE", changeIDKey, "("+strconv.Itoa(from), strconv.Itoa(to)))
	if err != nil {
		return nil, fmt.Errorf("get changed keys from redis: %w", err)
	}

	var cleanedKeys []string
	for _, key := range keys {
		if !strings.HasPrefix(key, "_config:") {
			cleanedKeys = append(cleanedKeys, key)
		}
	}
	return cleanedKeys, nil
}

// Data returns the data from redis for specific keys from the full data hash
// key.
//
// If a key does not exist, the value in the returned dict is nil.
func (r *Redis) Data(keys []string) (map[string]json.RawMessage, error) {
	conn := r.readPool.Get()
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

// AddApplause adds a user to the applause set.
//
// Also deletes applause that is older then a minute
func (r *Redis) AddApplause(userID int) error {
	conn := r.writePool.Get()
	defer conn.Close()

	if _, err := conn.Do("ZADD", applauseKey, time.Now().Unix(), userID); err != nil {
		return fmt.Errorf("adding applause in redis: %w", err)
	}

	return nil
}

// GetApplause returns the amount of applause sind a time.
func (r *Redis) GetApplause(since int64) (int, error) {
	conn := r.writePool.Get()
	defer conn.Close()

	// Delete old applause.
	before := time.Now().Add(-time.Minute)
	if _, err := conn.Do("ZREMRANGEBYSCORE", applauseKey, 0, before.Unix()); err != nil {
		return 0, fmt.Errorf("removing old applause from redis: %w", err)
	}

	c, err := redis.Int(conn.Do("ZCOUNT", applauseKey, since, "+inf"))
	if err != nil {
		return 0, fmt.Errorf("getting applause from redis: %w", err)
	}

	return c, nil
}

// SendNotify publishes a notify message in redis.
func (r *Redis) SendNotify(message []byte) error {
	conn := r.writePool.Get()
	defer conn.Close()

	_, err := conn.Do("XADD", notifyKey, "*", "content", message)
	if err != nil {
		return fmt.Errorf("xadd: %w", err)
	}
	return nil
}

// ReceiveNotify returns the next notify message.
func (r *Redis) ReceiveNotify(closing <-chan struct{}) (message string, err error) {
	conn := r.readPool.Get()
	defer conn.Close()

	id := r.lastNotifyID
	if id == "" {
		id = "$"
	}

	var data []byte
	received := make(chan struct{})

	go func() {
		id, data, err = stream(conn.Do("XREAD", "COUNT", 1, "BLOCK", "0", "STREAMS", notifyKey, id))
		close(received)
	}()

	select {
	case <-received:
	case <-closing:
		return "", closingError{}
	}

	if id != "" {
		r.lastNotifyID = id
	}

	if err != nil {
		if err == errNil {
			// No new data
			return "", nil
		}

		return "", fmt.Errorf("read notify from redis: %w", err)
	}

	return string(data), nil
}
