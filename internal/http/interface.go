package http

import "net/http"

// Auther authenticates a request.
type Auther interface {
	Middleware(func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) error
}
