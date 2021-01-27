package notify

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ChannelID is an id for a notify channel.
type ChannelID string

// uid returnes the user id that was used to create the channel id. Returns 0
// for an invalid channel id.
func (c ChannelID) uid() int {
	parts := strings.Split(string(c), ":")
	if len(parts) != 3 {
		return 0
	}

	uid, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0
	}
	return uid
}

func (c ChannelID) String() string {
	return string(c)
}

type cIDGen struct {
	host    string
	hostGen sync.Once

	mu    sync.Mutex
	count uint64
}

func (c *cIDGen) generate(uid int) ChannelID {
	c.hostGen.Do(c.buildHostID)

	c.mu.Lock()
	count := c.count
	c.count++
	c.mu.Unlock()

	cid := fmt.Sprintf("%s:%d:%d", c.host, uid, count)
	return ChannelID(cid)
}

func (c *cIDGen) buildHostID() {
	rand.Seed(time.Now().UnixNano())

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 8

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	c.host = string(b)
}
