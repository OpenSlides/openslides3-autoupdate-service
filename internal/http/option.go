package http

import (
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/notify"
)

// Option for http.New()
type Option func(*Handler)

// WithNotify adds routes for the notify package.
func WithNotify(n *notify.Notify) Option {
	return func(h *Handler) {
		h.mux.Handle("/system/notify", errHandleFunc(h.middleware(n.HandleNotify)))
		h.mux.Handle("/system/notify/send", errHandleFunc(h.middleware(n.HandleSend)))
	}
}

// WithAutoupdate adds routes for the autoupdate package.
func WithAutoupdate(a *autoupdate.Autoupdate) Option {
	return func(h *Handler) {
		h.mux.Handle("/system/autoupdate", errHandleFunc(h.middleware(a.HandleAutoupdate)))
		h.mux.Handle("/system/projector", errHandleFunc(h.middleware(a.HandleProjector)))
	}
}
