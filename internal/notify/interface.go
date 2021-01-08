package notify

import "net/http"

// Backend handels sending and receiving of notify messages.
type Backend interface {
	SendNotify(string) error
	ReceiveNotify(closing <-chan struct{}) (mail string, err error)
	AddApplause(userID int) error
	GetApplause(since int64) (int, error)
}

// Auther authenticates a request.
type Auther interface {
	Middleware(func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) error
}
