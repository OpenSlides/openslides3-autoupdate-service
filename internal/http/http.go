package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/autoupdate"
)

// Handler handels client requests to the autoupdate service.
type Handler struct {
	autoupdate *autoupdate.Autoupdate
	mux        *http.ServeMux
	auther     Auther
	keepAlive  int
}

// New create a new Handler with the correct urls.
func New(autoupdate *autoupdate.Autoupdate, auther Auther, keepAlive int, addHandler http.Handler) *Handler {
	h := &Handler{
		autoupdate: autoupdate,
		mux:        http.NewServeMux(),
		auther:     auther,
		keepAlive:  keepAlive,
	}
	h.mux.Handle("/system/autoupdate", http.HandlerFunc(h.handleAutoupdate))
	h.mux.Handle("/system/health", http.HandlerFunc(h.handleHealth))
	h.mux.Handle("/", addHandler)
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"healthy": true}`)
}

func (h *Handler) handleAutoupdate(w http.ResponseWriter, r *http.Request) {
	uid, err := h.auther.Auth(r)
	if err != nil {
		sendErr(w, fmt.Errorf("authenticate: %w", err))
		return
	}

	rawChangeID := r.URL.Query().Get("change_id")
	changeID := 0
	if rawChangeID != "" {
		var err error
		changeID, err = strconv.Atoi(rawChangeID)
		if err != nil {
			http.Error(w, fmt.Sprintf("change id has to be a number not %s", rawChangeID), http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.(http.Flusher).Flush()
	log.Printf("connect user %d with change_id %d", uid, changeID)

	var data map[string]json.RawMessage
	var all bool
	var newChangeID int
	ticker := new(time.Ticker)
	if h.keepAlive > 0 {
		ticker = time.NewTicker(time.Duration(h.keepAlive) * time.Second)
		defer ticker.Stop()
	}

	for {
		event := make(chan struct{})
		go func() {
			all, data, newChangeID, err = h.autoupdate.Receive(r.Context(), uid, changeID)
			close(event)
		}()

		select {
		case <-ticker.C:
			if err := sendKeepAlive(w); err != nil {
				sendErr(w, err)
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
				// Server is closing. Close connection.
				return
			}
			if errors.Is(err, context.Canceled) {
				// Client closed connection.
				return
			}
			sendErr(w, err)
			return
		}

		if err := sendData(w, all, data, changeID, newChangeID); err != nil {
			sendErr(w, err)
			return
		}
		changeID = newChangeID
	}
}

func sendData(w io.Writer, all bool, data map[string]json.RawMessage, fromChangeID, toChangeID int) error {
	changed := make(map[string][]json.RawMessage)
	deleted := make(map[string][]int)
	for k := range data {
		parts := strings.Split(k, ":")
		if len(parts) != 2 {
			return fmt.Errorf("invalid key %s, expected exacly one `:`", k)
		}

		collection := parts[0]
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("invalid key %s, id is not a number", k)
		}

		if data[k] == nil {
			if !all {
				deleted[collection] = append(deleted[collection], id)
			}
			continue
		}

		changed[collection] = append(changed[collection], data[k])
	}

	format := struct {
		Changed      map[string][]json.RawMessage `json:"changed"`
		Deleted      map[string][]int             `json:"deleted"`
		FromChangeID int                          `json:"from_change_id"`
		ToChangeID   int                          `json:"to_change_id"`
		AllData      bool                         `json:"all_data"`
	}{
		changed,
		deleted,
		fromChangeID,
		toChangeID,
		all,
	}

	if err := json.NewEncoder(w).Encode(format); err != nil {
		return fmt.Errorf("encode and send output data: %w", err)
	}
	w.(http.Flusher).Flush()
	return nil
}

// internalErr sends a nonsense error message to the client and logs the real
// message to stdout.
func sendErr(w io.Writer, err error) {
	var clientError interface {
		ClientError() string
		Error() string
	}
	if errors.As(err, &clientError) {
		fmt.Fprintf(w, `{"error": {"type": "%s", "msg": "%s"}}`, clientError.ClientError(), clientError.Error())
		fmt.Fprintln(w)
		return
	}

	log.Printf("Internal Error: %v", err)
	fmt.Fprintln(w, `{"error": {"type": "InternalError", "msg": "Ups, something went wrong!"}}`)
	w.(http.Flusher).Flush()
}

func sendKeepAlive(w io.Writer) error {
	_, err := fmt.Fprintln(w, `{}`)
	w.(http.Flusher).Flush()
	return err
}
