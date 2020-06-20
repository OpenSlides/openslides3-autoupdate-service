package notify

import "net/http"

// Backend handels sending and receiving of notify messages.
type Backend interface {
	SendNotify(string) error
	ReceiveNotify(closing <-chan struct{}) (mail string, err error)
}

// Auther authenticates a request.
type Auther interface {
	Middleware(http.Handler) http.Handler
}
