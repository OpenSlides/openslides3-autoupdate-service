package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/OpenSlides/openslides3-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/notify"
)

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
func Autoupdate(mux *http.ServeMux, auto *autoupdate.Autoupdate, auth Auther) {
	mux.Handle("/system/autoupdate", errHandleFunc(middleware(auto.HandleAutoupdate, auth)))
}

// Projector registers the projector route.
func Projector(mux *http.ServeMux, auto *autoupdate.Autoupdate, auth Auther) {
	mux.Handle("/system/projector", errHandleFunc(middleware(auto.HandleProjector, auth)))
}

// Notify registers the notify route.
func Notify(mux *http.ServeMux, n *notify.Notify, auth Auther) {
	mux.Handle("/system/notify", errHandleFunc(middleware(n.HandleNotify, auth)))
}

// NotifySend registers the notify/send route.
func NotifySend(mux *http.ServeMux, n *notify.Notify, auth Auther) {
	mux.Handle("/system/notify/send", errHandleFunc(middleware(n.HandleSend, auth)))
}

// NotifyApplause registers the notify/applause route.
func NotifyApplause(mux *http.ServeMux, n *notify.Notify, auth Auther) {
	mux.Handle("/system/applause", errHandleFunc(middleware(n.HandleApplause, auth)))
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
			// Server is closing.
			return
		}

		if errors.Is(err, context.Canceled) {
			// Client closed connection.
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
	}
}

func http2Middleware(next errHandleFunc) errHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		// Only allow http2 requests.
		if !r.ProtoAtLeast(2, 0) {
			return invalidRequestError{fmt.Errorf("Only http2 is supported")}
		}

		return next(w, r)
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
func middleware(next errHandleFunc, a Auther) errHandleFunc {
	r := a.Middleware(getOrPOSTMiddleware(next))
	r = http2Middleware(r)
	return r
}
