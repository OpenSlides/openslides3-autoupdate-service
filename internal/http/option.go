package http

import (
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides3-autoupdate-service/internal/notify"
)

// Option for http.New()
type Option func(*Handler)

// WithForceHTTP2 force http2.
func WithForceHTTP2(force bool) Option {
	return func(h *Handler) {
		h.forceHTTP2 = force
	}
}

// WithAutoupdate adds routes for the autoupdate package.
func WithAutoupdate(a *autoupdate.Autoupdate) Option {
	return func(h *Handler) {
		addErrHandleFunc(h, "/system/autoupdate", a.HandleAutoupdate)
		addErrHandleFunc(h, "/system/projector", a.HandleProjector)
	}
}

// WithNotify adds routes for the notify package.
func WithNotify(n *notify.Notify) Option {
	return func(h *Handler) {
		addErrHandleFunc(h, "/system/notify", n.HandleNotify)
		addErrHandleFunc(h, "/system/notify/send", n.HandleSend)
	}
}

// WithErrHandleFunc adds a http.errHandleFunc to the Handler.
func WithErrHandleFunc(pattern string, handler errHandleFunc) Option {
	return func(h *Handler) {
		addErrHandleFunc(h, pattern, handler)
	}
}

func addErrHandleFunc(h *Handler, pattern string, handler errHandleFunc) {
	h.mux.Handle(pattern, errHandleFunc(h.middleware(handler)))
}
