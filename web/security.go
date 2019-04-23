package web

import (
	"net/http"
	"strings"

	"github.com/fsufitch/discord-boar-bot/common"
)

// SecretBearerAuthorizationWrapper is an object for wrapping handlers to only allow authorized requets
type SecretBearerAuthorizationWrapper struct {
	secret string
}

type wrappedHandler struct {
	secret  string
	handler http.Handler
}

func (h wrappedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Bearer ") {
		plainTextResponse(w, http.StatusUnauthorized, []byte("Missing `Authorization: Bearer ...` header"))
		return
	}

	token := auth[7:]
	if token != h.secret {
		plainTextResponse(w, http.StatusForbidden, []byte("Incorrect authorization"))
		return
	}

	h.handler.ServeHTTP(w, r)
}

// Wrap returns a new handler that transparently checks the secret bearer header
func (w SecretBearerAuthorizationWrapper) Wrap(h http.Handler) http.Handler {
	return wrappedHandler{
		secret:  w.secret,
		handler: h,
	}
}

// NewSecretBearerAuthorizationWrapper creates a SecretBearerAuthorizationWrapper
func NewSecretBearerAuthorizationWrapper(
	configuration *common.Configuration,
) *SecretBearerAuthorizationWrapper {
	return &SecretBearerAuthorizationWrapper{
		secret: configuration.WebSecret,
	}
}