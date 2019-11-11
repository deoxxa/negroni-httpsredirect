package negronihttpsredirect

import (
	"net/http"
)

type Option func(h *HTTPSRedirect)

func Header(header string) Option {
	return func(h *HTTPSRedirect) {
		h.header = header
	}
}

func Status(status int) Option {
	return func(h *HTTPSRedirect) {
		h.status = status
	}
}

type HTTPSRedirect struct {
	header string
	status int
}

func New(options ...Option) *HTTPSRedirect {
	h := HTTPSRedirect{header: "x-forwarded-proto", status: http.StatusSeeOther}

	for _, option := range options {
		option(&h)
	}

	return &h
}

func (h *HTTPSRedirect) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if hdr := r.Header.Get(h.header); hdr == "http" {
		http.Redirect(rw, r, "https://"+r.Host+r.RequestURI, h.status)
		return
	}

	next(rw, r)
}
