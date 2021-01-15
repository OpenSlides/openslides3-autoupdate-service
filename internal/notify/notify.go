package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ostcar/topic"
)

// Notify is a service to send messages between clients.
type Notify struct {
	backend          Backend
	topic            *topic.Topic
	closed           <-chan struct{}
	cIDGen           cIDGen
	applauser        Applauser
	applauseInterval int
}

// New returns an initializes Notify object.
func New(backend Backend, applauser Applauser, applauseInterval int, closed <-chan struct{}) *Notify {
	n := &Notify{
		backend:          backend,
		applauser:        applauser,
		applauseInterval: applauseInterval,
		topic:            topic.New(topic.WithClosed(closed)),
		closed:           closed,
	}

	go n.listen()
	// TODO prune topic

	go n.applauseLoop()

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

func (n *Notify) applauseLoop() {
	var last int
	for {
		time.Sleep(time.Duration(n.applauseInterval) * time.Millisecond)
		select {
		case <-n.closed:
			return
		default:
		}

		a, base, err := n.receiceApplause()
		if err != nil {
			log.Printf("Notify: Can not receive applause: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if a == 0 && last == 0 {
			continue
		}
		last = a

		b, err := json.Marshal(struct {
			Level        int `json:"level"`
			PresentUsers int `json:"presentUsers"`
		}{
			a,
			base,
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
			From:    "Server",
		})
		if err != nil {
			log.Printf("Notify: Can not encode applause message: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		n.topic.Publish(string(m))
	}
}

func (n *Notify) receiceApplause() (int, int, error) {
	countTime, base := n.applauser.ApplauseConfig()

	d := time.Now().Add(-time.Duration(countTime) * time.Second)
	a, err := n.backend.GetApplause(d.Unix())
	if err != nil {
		return 0, 0, fmt.Errorf("getting applause: %w", err)
	}
	return a, base, nil
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
