package http

import "net/http"

// Auther authenticates a request.
type Auther interface {
	Auth(r *http.Request) (uid int, err error)
}
