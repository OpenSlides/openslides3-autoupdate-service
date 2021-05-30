package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/auth"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/notify"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

var meter = global.GetMeterProvider().Meter("openslides.org")

// RegisterAll registers all routes.
func RegisterAll(mux *http.ServeMux, auth Auther, a *autoupdate.Autoupdate, n *notify.Notify) {
	Health(mux)
	Autoupdate(mux, a, auth)
	Projector(mux, a, auth)
	Notify(mux, n, auth)
	NotifySend(mux, n, auth)
	NotifyApplause(mux, n, auth)
}

// Health registers the health route.
func Health(mux *http.ServeMux) {
	mux.HandleFunc("/system/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"healthy": true}`)
	})
}

// Autoupdate registers the autoupdate route.
func Autoupdate(mux *http.ServeMux, auto *autoupdate.Autoupdate, auther Auther) {
	count := newConnectionCount("autoupdate")

	handler := func(w http.ResponseWriter, r *http.Request) error {
		count.Add()
		defer count.Sub()

		uid := auth.FromContext(r.Context())
		w.Header().Set("Content-Type", "application/octet-stream")

		rawChangeID := r.URL.Query().Get("change_id")
		var changeID int
		if rawChangeID != "" {
			var err error
			changeID, err = strconv.Atoi(rawChangeID)
			if err != nil {
				return invalidRequestError{fmt.Errorf("Change id has to be a number not %s", rawChangeID)}
			}
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"connected":true}`)
		w.(http.Flusher).Flush()

		for {
			all, data, newChangeID, err := auto.Receive(r.Context(), uid, changeID)
			if err != nil {
				return noStatusCodeError{err}
			}

			if len(data) == 0 {
				changeID = newChangeID
				continue
			}

			if err := sendAutoupdateData(w, all, data, changeID, newChangeID); err != nil {
				return noStatusCodeError{err}
			}
			changeID = newChangeID
		}
	}

	mux.Handle("/system/autoupdate", errHandleFunc(middleware(handler, auther)))
}

// Projector registers the projector route.
func Projector(mux *http.ServeMux, auto *autoupdate.Autoupdate, auth Auther) {
	count := newConnectionCount("projector")

	handler := func(w http.ResponseWriter, r *http.Request) error {
		count.Add()
		defer count.Sub()

		w.Header().Set("Content-Type", "application/json")

		projectorIDs, err := projectorIDs(r.URL.Query().Get("projector_ids"))
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"connected":true}`)
		w.(http.Flusher).Flush()

		encoder := json.NewEncoder(w)
		var tid uint64
		for {
			ntid, data, cid, err := auto.Projectors(r.Context(), tid, projectorIDs)
			if err != nil {
				return noStatusCodeError{err}
			}

			out := struct {
				CID  int                     `json:"change_id"`
				Data map[int]json.RawMessage `json:"data"`
			}{
				cid,
				data,
			}

			if err := encoder.Encode(out); err != nil {
				return noStatusCodeError{err}
			}
			w.(http.Flusher).Flush()
			tid = ntid
		}
	}
	mux.Handle("/system/projector", errHandleFunc(middleware(handler, auth)))
}

// Notify registers the notify route.
func Notify(mux *http.ServeMux, n *notify.Notify, auther Auther) {
	count := newConnectionCount("notify")

	handler := func(w http.ResponseWriter, r *http.Request) error {
		count.Add()
		defer count.Sub()

		uid := auth.FromContext(r.Context())
		w.Header().Set("Content-Type", "application/octet-stream")

		cid := n.GenerateChannelID(uid)
		tid := n.LastID()

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"connected":true}`)

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
			tid, err = n.Receive(r.Context(), w, tid, uid, cid, encoder)
			if err != nil {
				return noStatusCodeError{fmt.Errorf("rececing notify data: %w", err)}
			}
			w.(http.Flusher).Flush()
		}
	}
	mux.Handle("/system/notify", errHandleFunc(middleware(handler, auther)))
}

// NotifySend registers the notify/send route.
func NotifySend(mux *http.ServeMux, n *notify.Notify, auther Auther) {
	counter, _ := meter.NewInt64Counter(
		"openslides.notify-send-requests",
		metric.WithDescription("request count to notify send"),
	)

	handler := func(w http.ResponseWriter, r *http.Request) error {
		counter.Add(r.Context(), 1)

		userID := auth.FromContext(r.Context())
		if userID == 0 {
			return authRequiredError{"You have to be logged in to use the notify system."}
		}

		bs, err := io.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("reading message: %w", err)
		}

		if err := notify.ValidateRequest(bs, userID); err != nil {
			return invalidRequestError{err}
		}

		buf := new(bytes.Buffer)
		if err := json.Compact(buf, bs); err != nil {
			return invalidRequestError{err}
		}

		return n.Send(buf.Bytes(), userID)
	}
	mux.Handle("/system/notify/send", errHandleFunc(middleware(handler, auther)))
}

// NotifyApplause registers the notify/applause route.
func NotifyApplause(mux *http.ServeMux, n *notify.Notify, auther Auther) {
	counter, _ := meter.NewInt64Counter(
		"openslides.notify-applause-requests",
		metric.WithDescription("request count to applause send"),
	)

	handler := func(w http.ResponseWriter, r *http.Request) error {
		counter.Add(r.Context(), 1)
		userID := auth.FromContext(r.Context())
		if userID == 0 {
			return authRequiredError{"You have to be logged in to send applause."}
		}

		return n.AddApplause(userID)
	}
	mux.Handle("/system/applause", errHandleFunc(middleware(handler, auther)))
}

// errHandleFunc is like a http.Handler, but has a error as return value.
//
// If the returned error implements the clientError interface, then the error
// message is sent to the client. In other cases the error is interpredet as an
// internal error and logged to stdout.
type errHandleFunc func(w http.ResponseWriter, r *http.Request) error

var errCount, _ = meter.NewInt64Counter(
	"openslides.http-error-count",
	metric.WithDescription("500er send to the client"),
)

func (f errHandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		var closing interface {
			Closing()
		}
		if errors.As(err, &closing) {
			// Server is closing.
			return
		}

		if errors.Is(err, context.Canceled) {
			// Client closed connection.
			return
		}

		var errNet *net.OpError
		if errors.As(err, &errNet) {
			// This happens, when a client closes a connection without telling
			// the server.
			return
		}

		var noStatusErr interface {
			NoStatus()
		}
		var status bool
		if !errors.As(err, &noStatusErr) {
			status = true
		}

		var clientError interface {
			ClientError() string
			Error() string
		}
		if errors.As(err, &clientError) {
			if status {
				w.WriteHeader(http.StatusBadRequest)
			}
			fmt.Fprintf(
				w,
				`{"error": {"type": "%s", "msg": "%s"}}`,
				clientError.ClientError(),
				clientError.Error(),
			)
			fmt.Fprintln(w)
			return
		}

		if status {
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Printf("Internal Error: %v", err)
		fmt.Fprintln(w, `{"error": {"type": "InternalError", "msg": "Ups, something went wrong!"}}`)
		errCount.Add(r.Context(), 1)
	}
}

func getOrPOSTMiddleware(next errHandleFunc) errHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		// Only allow GET or POST requests.
		if !(r.Method == http.MethodPost || r.Method == http.MethodGet) {
			return invalidRequestError{fmt.Errorf("Only GET or POST requests are supported")}
		}

		return next(w, r)
	}
}

// middleware combines all necessary middlewares.
func middleware(next errHandleFunc, auther Auther) errHandleFunc {
	return authMiddleware(getOrPOSTMiddleware(next), auther)
}

func authMiddleware(next errHandleFunc, auther Auther) errHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx, err := auther.Authenticate(r)
		if err != nil {
			var errAuth auth.Error
			if errors.As(err, &errAuth) {
				return authRequiredError{msg: errAuth.Error()}
			}
			return fmt.Errorf("authenticate request: %w", err)
		}
		return next(w, r.WithContext(ctx))
	}
}

func sendAutoupdateData(w io.Writer, all bool, data map[string]json.RawMessage, fromChangeID, toChangeID int) error {
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
		return fmt.Errorf("encode and send output data, error type %T: %w", err, err)
	}
	w.(flusher).Flush()
	return nil
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

type connectionCount struct {
	mu sync.Mutex
	v  int
	m  metric.Int64Counter
}

func newConnectionCount(name string) *connectionCount {
	c := new(connectionCount)

	c.m, _ = meter.NewInt64Counter(
		fmt.Sprintf("openslides.%s-connection-overall", name),
		metric.WithDescription("overall connections to "+name),
	)

	meter.NewInt64UpDownSumObserver(
		fmt.Sprintf("openslides.%s-connection-current", name),
		func(_ context.Context, result metric.Int64ObserverResult) {
			result.Observe(int64(c.Value()))
		},
		metric.WithDescription("current open connections to "+name),
	)
	return c
}

func (c *connectionCount) Add() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.v++
	c.m.Add(context.Background(), 1)
	return c.v
}

func (c *connectionCount) Sub() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.v--
	return c.v
}

func (c *connectionCount) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.v
}
