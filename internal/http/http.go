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

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/autoupdate"
)

// Handler handels client requests to the autoupdate service.
type Handler struct {
	autoupdate *autoupdate.Autoupdate
	mux        *http.ServeMux
	auther     Auther
}

// New create a new Handler with the correct urls.
func New(autoupdate *autoupdate.Autoupdate, auther Auther, addHandler http.Handler) *Handler {
	h := &Handler{
		autoupdate: autoupdate,
		mux:        http.NewServeMux(),
		auther:     auther,
	}
	h.mux.Handle("/system/autoupdate", validRequest(http.HandlerFunc(h.handleAutoupdate)))
	h.mux.Handle("/system/health", validRequest(http.HandlerFunc(h.handleHealth)))
	h.mux.Handle("/system/projector", validRequest(http.HandlerFunc(h.handleProjector)))
	if addHandler != nil {
		h.mux.Handle("/", addHandler)
	}
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
	w.Header().Set("Content-Type", "application/octet-stream")

	uid, err := h.auther.Auth(r)
	if err != nil {
		sendErr(w, fmt.Errorf("authenticate: %w", err))
		return
	}

	rawChangeID := r.URL.Query().Get("change_id")
	var changeID int
	if rawChangeID != "" {
		var err error
		changeID, err = strconv.Atoi(rawChangeID)
		if err != nil {
			http.Error(w, fmt.Sprintf("change id has to be a number not %s", rawChangeID), http.StatusBadRequest)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.(http.Flusher).Flush()
	log.Printf("connect user %d with change_id %d", uid, changeID)

	for {
		changeID, err = h.autoupdateLoop(w, r, changeID, uid)
		if err != nil {
			sendErr(w, err)
			return
		}
	}
}

func (h *Handler) autoupdateLoop(w http.ResponseWriter, r *http.Request, cid, uid int) (int, error) {
	all, data, newChangeID, err := h.autoupdate.Receive(r.Context(), uid, cid)
	if err != nil {
		return 0, err
	}

	if len(data) == 0 {
		return newChangeID, nil
	}

	if err := sendData(w, all, data, cid, newChangeID); err != nil {
		return 0, err
	}
	return newChangeID, nil
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

func (h *Handler) handleProjector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// The uid is not needed. But this makes sure the user is authenticated or
	// anonymous is enabled.
	_, err := h.auther.Auth(r)
	if err != nil {
		sendErr(w, fmt.Errorf("authenticate: %w", err))
		return
	}

	projectorIDs, err := projectorIDs(r.URL.Query().Get("projector_ids"))
	if err != nil {
		sendErr(w, err)
		return
	}

	encoder := json.NewEncoder(w)
	var tid uint64
	for {
		tid, err = h.projectorLoop(r.Context(), w, encoder, tid, projectorIDs)
		if err != nil {
			sendErr(w, err)
			return
		}
	}
}

func (h *Handler) projectorLoop(ctx context.Context, w io.Writer, encoder *json.Encoder, tid uint64, pids []int) (uint64, error) {
	ntid, data, cid, err := h.autoupdate.Projectors(ctx, tid, pids)
	if err != nil {
		return 0, err
	}

	out := struct {
		CID  int                     `json:"change_id"`
		Data map[int]json.RawMessage `json:"data"`
	}{
		cid,
		data,
	}

	if err := encoder.Encode(out); err != nil {
		return 0, nil
	}
	w.(http.Flusher).Flush()
	return ntid, nil
}

// internalErr sends a nonsense error message to the client and logs the real
// message to stdout.
func sendErr(w io.Writer, err error) {
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

func projectorIDs(raw string) ([]int, error) {
	parts := strings.Split(raw, ",")
	ids := make([]int, len(parts))
	for i, rpid := range parts {
		id, err := strconv.Atoi(rpid)
		if err != nil {
			return nil, invalidRequestError{fmt.Errorf("projector_ids has to be a list of ints not `%s`", raw)}
		}

		ids[i] = id
	}
	return ids, nil
}

func validRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only allow http2 requests.
		if !r.ProtoAtLeast(2, 0) {
			http.Error(w, "Only http2 is supported", http.StatusBadRequest)
			return
		}

		// Only allow GET or POST requests.
		if !(r.Method == http.MethodPost || r.Method == http.MethodGet) {
			http.Error(w, "Only GET or POST requests are supported", http.StatusMethodNotAllowed)
			return
		}

		h.ServeHTTP(w, r)
	})
}
