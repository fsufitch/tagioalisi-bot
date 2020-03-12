package auth

import (
	"net/http"
	"strings"
)

// GetSessionID returns and the session ID associated with a request
func GetSessionID(r *http.Request) string {
	if r.URL.Query().Get("sid") != "" {
		return r.URL.Query().Get("sid")
	}

	authHeader := r.Header.Get("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") && authHeader[7:] != "" {
		return authHeader[7:]
	}
	return ""
}
