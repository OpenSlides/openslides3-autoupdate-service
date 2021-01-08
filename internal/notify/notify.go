package notify

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/ostcar/topic"
)

// Notify is a service to send messages between clients.
type Notify struct {
	backend Backend
	topic   *topic.Topic
	closed  <-chan struct{}
	cIDGen  cIDGen
}

// New returns an initializes Notify object.
func New(backend Backend, closed <-chan struct{}) *Notify {
	n := &Notify{
		backend: backend,
		topic:   topic.New(topic.WithClosed(closed)),
		closed:  closed,
	}

	go n.listen()
	// TODO prune topic

	go n.applause()

	return n
}

// listen waits for notify messages from the backend and saves them into the
// topic.
func (n *Notify) listen() {
	for {
		m, err := n.backend.ReceiveNotify(n.closed)
		if err != nil {
			var closing interface {
				Closing()
			}
			if errors.As(err, &closing) {
				return
			}

			log.Printf("Notify: Can not receive data from backend: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		n.topic.Publish(m)
	}
}

func (n *Notify) applause() {
	for {
		// TODO: get since time from config
		d := time.Now().Add(-5 * time.Second)
		a, err := n.backend.GetApplause(d.Unix())
		if err != nil {
			log.Printf("Notify: Can not receice applause: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		b, err := json.Marshal(struct {
			Count int `json:"count"`
			Base  int `json:"base"`
		}{
			a,
			100, // TODO: get base
		})
		if err != nil {
			log.Printf("Notify: Can not encode applause: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		m, err := json.Marshal(mail{
			ToAll:   true,
			Name:    "applause",
			Message: b,
		})
		if err != nil {
			log.Printf("Notify: Can not encode applause message: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		n.topic.Publish(string(m))
		time.Sleep(time.Second)
	}
}

type mail struct {
	From       channelID       `json:"channel_id"`
	ToAll      bool            `json:"to_all"`
	ToUsers    []int           `json:"to_users"`
	ToChannels []string        `json:"to_channels"`
	Name       string          `json:"name"`
	Message    json.RawMessage `json:"message"`
}

func (m mail) forMe(uid int, cID channelID) bool {
	if m.ToAll {
		return true
	}

	for _, toUID := range m.ToUsers {
		if toUID == uid {
			return true
		}
	}

	for _, toCID := range m.ToChannels {
		if toCID == cID.String() {
			return true
		}
	}
	return false
}
