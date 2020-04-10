package web

import (
	"net/http"

	"github.com/fsufitch/tagioalisi-bot/log"
)

// HelloHandler is a http.Handler just says hello world
type HelloHandler struct {
	Log *log.Logger
}

// ServeHTTP just says hello world
func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	h.Log.HTTP(http.StatusOK, r)
	w.Write([]byte("Hello, world!"))
}
