package notify

import (
	"bytes"
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

	if from.Name == "" {
		return invalidRequestError{fmt.Errorf("notify does not have required field `name`")}
	}

	buf := new(bytes.Buffer)
	if err := json.Compact(buf, m); err != nil {
		return invalidRequestError{fmt.Errorf("json is invalid: %w", err)}
	}

	if err := n.backend.SendNotify(buf.String()); err != nil {
		return fmt.Errorf("sending message: %w", err)
	}
	return nil
}

func (n *Notify) handleNotify(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/octet-stream")

	userID, ok := r.Context().Value(auth.UserIDKey).(int)
	if !ok || userID == 0 {
		return authRequiredError{}
	}

	cid := n.cIDGen.generate(userID)
	tid := n.topic.LastID()

	if _, err := fmt.Fprintf(w, `{"channel_id": "%s"}`, cid); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(w); err != nil {
		return err
	}
	w.(http.Flusher).Flush()

	encoder := json.NewEncoder(w)
	var err error

	for {
		tid, err = n.autoupdateLoop(w, r, tid, userID, cid, encoder)
		if err != nil {
			return err
		}
	}
}

func (n *Notify) autoupdateLoop(w http.ResponseWriter, r *http.Request, tid uint64, uid int, cid channelID, encoder *json.Encoder) (uint64, error) {
	ctx := r.Context()
	if n.keepAlive > 0 {
		var cancel func()
		ctx, cancel = context.WithTimeout(r.Context(), time.Duration(n.keepAlive)*time.Second)
		defer cancel()
	}

	var rMails []string
	var err error
	tid, rMails, err = n.topic.Receive(ctx, tid)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			if err := sendKeepAlive(w); err != nil {
				return 0, err
			}
			return tid, nil
		}
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

	w.(http.Flusher).Flush()
	return tid, nil
}

// errHandleFunc is like a http.Handler, but has a error as return value.
//
// If the returned error implements the clientError interface, then the error
// message is sent to the client. In other cases the error is interpredet as an
// internal error and logged to stdout.
type errHandleFunc func(w http.ResponseWriter, r *http.Request) error

func (f errHandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		var closing interface {
			Closing()
		}
		if errors.As(err, &closing) {
			return
		}

		if errors.Is(err, context.Canceled) {
			return
		}

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
