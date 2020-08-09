package autoupdate

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/auth"
)

// HandleAutoupdate handels http requests.
func (a *Autoupdate) HandleAutoupdate(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/octet-stream")

	uid, ok := r.Context().Value(auth.UserIDKey).(int)
	if !ok || uid == 0 {
		return authRequiredError{}
	}

	rawChangeID := r.URL.Query().Get("change_id")
	var changeID int
	if rawChangeID != "" {
		var err error
		changeID, err = strconv.Atoi(rawChangeID)
		if err != nil {
			return fmt.Errorf("change id has to be a number not %s", rawChangeID)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.(http.Flusher).Flush()
	log.Printf("connect user %d with change_id %d", uid, changeID)

	for {
		all, data, newChangeID, err := a.Receive(r.Context(), uid, changeID)
		if err != nil {
			return err
		}

		if err := sendData(w, all, data, changeID, newChangeID); err != nil {
			return err
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

// HandleProjector handels http requests to get projector data.
func (a *Autoupdate) HandleProjector(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	projectorIDs, err := projectorIDs(r.URL.Query().Get("projector_ids"))
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(w)
	var tid uint64
	for {
		ntid, data, cid, err := a.Projectors(r.Context(), tid, projectorIDs)
		if err != nil {
			return err
		}

		out := struct {
			CID  int                     `json:"change_id"`
			Data map[int]json.RawMessage `json:"data"`
		}{
			cid,
			data,
		}

		if err := encoder.Encode(out); err != nil {
			return err
		}
		w.(http.Flusher).Flush()
		tid = ntid
	}
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
