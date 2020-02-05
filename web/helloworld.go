package web

import "net/http"

// HelloHandler is a http.Handler just says hello world
type HelloHandler struct{}

// ServeHTTP just says hello world
func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}
