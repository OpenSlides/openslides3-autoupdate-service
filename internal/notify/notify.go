package notify

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/ostcar/topic"
)

// Notify is a service to send messages between clients.
type Notify struct {
	mux       *http.ServeMux
	backend   Backend
	auther    Auther
	topic     *topic.Topic
	closed    <-chan struct{}
	cIDGen    cIDGen
	keepAlive int
}

// New returns an initializes Notify object.
func New(backend Backend, auth Auther, keepAlive int, closed <-chan struct{}) *Notify {
	n := &Notify{
		mux:       http.NewServeMux(),
		backend:   backend,
		topic:     topic.New(topic.WithClosed(closed)),
		closed:    closed,
		keepAlive: keepAlive,
	}
	n.mux.Handle("/system/notify", auth.Middleware(errHandleFunc(n.handleNotify)))
	n.mux.Handle("/system/notify/send", auth.Middleware(errHandleFunc(n.handleSend)))

	go n.listen()
	// TODO prune topic

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
			continue
		}

		n.topic.Publish(m)
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
