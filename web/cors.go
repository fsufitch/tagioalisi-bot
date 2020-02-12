package web

import (
	"net/http"
)

// CORSWrapper is a HTTPHandler wrapper for setting CORS headers
type CORSWrapper struct{}

type corsWrappedHandler struct {
	handler http.Handler
}

func (h corsWrappedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // TODO: make this smarter?
	h.handler.ServeHTTP(w, r)
}

// Wrap returns a new handler that transparently adds CORS headers
func (w CORSWrapper) Wrap(h http.Handler) http.Handler {
	return corsWrappedHandler{handler: h}
}
