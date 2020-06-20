package notify

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/auth"
)

func (n *Notify) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	n.mux.ServeHTTP(w, r)
}

func (n *Notify) handleSend(w http.ResponseWriter, r *http.Request) error {
	userID, ok := r.Context().Value(auth.UserIDKey).(int)
	if !ok || userID == 0 {
		return authRequiredError{}
	}

	m, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("reading request body: %w", err)
	}

	var from mail

	if err := json.Unmarshal(m, &from); err != nil {
		return invalidRequestError{err}
	}

	if from.From.uid() != userID {
		return invalidRequestError{fmt.Errorf("invalid channel id")}
	}

	if err := n.backend.SendNotify(string(m)); err != nil {
		return fmt.Errorf("sending message: %w", err)
	}
	return nil
}

func (n *Notify) handleNotify(w http.ResponseWriter, r *http.Request) error {
	userID, ok := r.Context().Value(auth.UserIDKey).(int)
	if !ok || userID == 0 {
		return authRequiredError{}
	}

	cid := n.cIDGen.generate(userID)
	tid := n.topic.LastID()

	if _, err := fmt.Fprintf(w, `{"channelID": "%s"}`, cid); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(w); err != nil {
		return err
	}
	w.(http.Flusher).Flush()

	encoder := json.NewEncoder(w)
	var rMails []string
	var err error

	ticker := new(time.Ticker)
	if n.keepAlive > 0 {
		ticker = time.NewTicker(time.Duration(n.keepAlive) * time.Second)
		defer ticker.Stop()
	}

	for {
		event := make(chan struct{})
		go func() {
			tid, rMails, err = n.topic.Receive(r.Context(), tid)
			close(event)
		}()

		select {
		case <-ticker.C:
			if err := sendKeepAlive(w); err != nil {
				return err
			}
			continue
		case <-event:
			// Received autoupdate event.
			// TODO: Reset ticker. This will be possible in go 1.15 that will be released in august:
			//       https://tip.golang.org/doc/go1.15#time
		}

		if err != nil {
			var closing interface {
				Closing()
			}
			if errors.As(err, &closing) {
				return nil
			}

			if errors.Is(err, context.Canceled) {
				return nil
			}
			return fmt.Errorf("receiving message: %w", err)
		}

		for _, rMail := range rMails {
			var m mail
			if err := json.Unmarshal([]byte(rMail), &m); err != nil {
				return fmt.Errorf("decoding message: %w", err)
			}

			if !m.forMe(userID, cid) {
				continue
			}

			out := struct {
				SenderUserID    int             `json:"sender_user_id"`
				SenderChannelID string          `json:"sender_channel_id"`
				Message         json.RawMessage `json:"message"`
			}{
				m.From.uid(),
				m.From.String(),
				m.Message,
			}

			if err := encoder.Encode(out); err != nil {
				return fmt.Errorf("sending message: %w", err)
			}
		}

		w.(http.Flusher).Flush()
	}
}

// errHandleFunc is like a http.Handler, but has a error as return value.
//
// If the returned error implements the clientError interface, then the error
// message is sent to the client. In other cases the error is interpredet as an
// internal error and logged to stdout.
type errHandleFunc func(w http.ResponseWriter, r *http.Request) error

func (f errHandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		var clientError interface {
			ClientError() string
			Error() string
		}
		if errors.As(err, &clientError) {
			fmt.Fprintf(
				w,
				`{"error": {"type": "%s", "msg": "%s"}}`,
				clientError.ClientError(),
				clientError.Error(),
			)
			fmt.Fprintln(w)
			return
		}

		log.Printf("Internal Error: %v", err)
		fmt.Fprintln(w, `{"error": {"type": "InternalError", "msg": "Ups, something went wrong!"}}`)
	}
}

func sendKeepAlive(w io.Writer) error {
	_, err := fmt.Fprintln(w, `{}`)
	w.(http.Flusher).Flush()
	return err
}
