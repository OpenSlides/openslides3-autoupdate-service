package http

import (
	"context"
	"net/http"
)

// Auther authenticates a request.
type Auther interface {
	Authenticate(r *http.Request) (context.Context, error)
}

// flusher is the same as http.Flusher
type flusher interface {
	Flush()
}
