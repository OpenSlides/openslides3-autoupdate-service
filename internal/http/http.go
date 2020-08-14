package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Handler handels client requests to the autoupdate service.
type Handler struct {
	mux        *http.ServeMux
	auther     Auther
	forceHTTP2 bool
}

// New create a new Handler with the correct urls.
func New(auther Auther, options ...Option) *Handler {
	h := &Handler{
		mux:    http.NewServeMux(),
		auther: auther,
	}
	h.mux.Handle("/system/health", errHandleFunc(h.handleHealth))

	for _, f := range options {
		f(h)
	}

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"healthy": true}`)
	return nil
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
func (h *Handler) middleware(next errHandleFunc) errHandleFunc {
	r := h.auther.Middleware(getOrPOSTMiddleware(next))

	if h.forceHTTP2 {
		r = http2Middleware(r)
	}
	return r
}
