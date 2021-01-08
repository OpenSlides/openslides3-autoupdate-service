package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/auth"
)

// HandleApplause adds the userID to the applause list.
func (n *Notify) HandleApplause(w http.ResponseWriter, r *http.Request) error {
	userID, ok := r.Context().Value(auth.UserIDKey).(int)
	if !ok || userID == 0 {
		return authRequiredError{}
	}

	n.backend.AddApplause(userID)
	return nil
}

// HandleSend is an http.ErrorHandlerFunc for sending a notify message.
func (n *Notify) HandleSend(w http.ResponseWriter, r *http.Request) error {
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

// HandleNotify is an http.ErrorHandlerFunc for receiving notify messages.
func (n *Notify) HandleNotify(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/octet-stream")

	userID, ok := r.Context().Value(auth.UserIDKey).(int)
	if !ok || userID == 0 {
		return authRequiredError{}
	}

	cid := n.cIDGen.generate(userID)
	tid := n.topic.LastID()

	w.WriteHeader(http.StatusOK)

	if _, err := fmt.Fprintf(w, `{"channel_id": "%s"}`, cid); err != nil {
		return noStatusCodeError{err}
	}

	if _, err := fmt.Fprintln(w); err != nil {
		return noStatusCodeError{err}
	}
	w.(http.Flusher).Flush()

	encoder := json.NewEncoder(w)
	var err error

	for {
		tid, err = n.autoupdateLoop(w, r, tid, userID, cid, encoder)
		if err != nil {
			return noStatusCodeError{err}
		}
	}
}

func (n *Notify) autoupdateLoop(w http.ResponseWriter, r *http.Request, tid uint64, uid int, cid channelID, encoder *json.Encoder) (uint64, error) {
	var rMails []string
	var err error
	tid, rMails, err = n.topic.Receive(r.Context(), tid)
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

	w.(http.Flusher).Flush()
	return tid, nil
}
