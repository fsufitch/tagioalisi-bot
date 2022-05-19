package web

import (
	"net/http"
)

func plainTextResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write(data)
}
