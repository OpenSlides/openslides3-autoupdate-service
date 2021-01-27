package notify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

// Receive recieces a notify message and sends it to the writer.
func (n *Notify) Receive(ctx context.Context, w io.Writer, tid uint64, uid int, cid ChannelID, encoder *json.Encoder) (uint64, error) {
	var rMails []string
	var err error
	tid, rMails, err = n.topic.Receive(ctx, tid)
	if err != nil {
		return 0, fmt.Errorf("receiving message: %w", err)
	}

	for _, rMail := range rMails {
		var m mail
		if err := json.Unmarshal([]byte(rMail), &m); err != nil {
			return 0, fmt.Errorf("decoding message: %w", err)
		}

		if !m.forMe(uid, cid) {
			continue
		}

		out := struct {
			SenderUserID    int             `json:"sender_user_id"`
			SenderChannelID string          `json:"sender_channel_id"`
			Name            string          `json:"name"`
			Message         json.RawMessage `json:"message"`
		}{
			m.From.uid(),
			m.From.String(),
			m.Name,
			m.Message,
		}

		if err := encoder.Encode(out); err != nil {
			return 0, fmt.Errorf("sending message: %w", err)
		}
	}

	return tid, nil
}

// GenerateChannelID creates a new channel id for a user.
func (n *Notify) GenerateChannelID(userID int) ChannelID {
	return n.cIDGen.generate(userID)
}

// LastID returns the last id from the topic.
func (n *Notify) LastID() uint64 {
	return n.topic.LastID()
}

// AddApplause adds the applause of a user.
func (n *Notify) AddApplause(userID int) error {
	return n.backend.AddApplause(userID)
}

// ValidateRequest tells, if the given data is a valid notify request.
//
// All errors from this function can be send to the clietn.
func ValidateRequest(m []byte, userID int) error {
	var from mail

	if err := json.Unmarshal(m, &from); err != nil {
		return fmt.Errorf("invalid json: %v", err)
	}

	if from.From.uid() != userID {
		return fmt.Errorf("invalid channel id")
	}

	if from.Name == "" {
		return fmt.Errorf("notify does not have required field `name`")
	}

	if from.Name == "applause" {
		return fmt.Errorf("notify name can not be applause")
	}
	return nil
}

// Send sends an autoupdate message provides as a reader.
func (n *Notify) Send(bs []byte, userID int) error {
	buf := new(bytes.Buffer)

	if err := n.backend.SendNotify(buf.String()); err != nil {
		return fmt.Errorf("sending message: %w", err)
	}

	return nil
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
	From       ChannelID       `json:"channel_id"`
	ToAll      bool            `json:"to_all"`
	ToUsers    []int           `json:"to_users"`
	ToChannels []string        `json:"to_channels"`
	Name       string          `json:"name"`
	Message    json.RawMessage `json:"message"`
}

func (m mail) forMe(uid int, cID ChannelID) bool {
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
