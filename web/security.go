package web

import (
	"net/http"
	"strings"

	"github.com/fsufitch/discord-boar-bot/config"
	"github.com/fsufitch/discord-boar-bot/log"
)

// SecretBearerAuthorizationWrapper is an object for wrapping handlers to only allow authorized requets
type SecretBearerAuthorizationWrapper struct {
	Secret config.WebSecret
	Log    *log.Logger
}

type wrappedHandler struct {
	log     *log.Logger
	secret  string
	handler http.Handler
}

func (h wrappedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	w.Header().Set("Access-Control-Allow-Origin", "*") // TODO: make this smarter?
	if !strings.HasPrefix(auth, "Bearer ") {
		plainTextResponse(w, http.StatusUnauthorized, []byte("Missing `Authorization: Bearer ...` header"))
		h.log.HTTP(http.StatusUnauthorized, r)
		return
	}

	token := auth[7:]
	if token != h.secret {
		plainTextResponse(w, http.StatusForbidden, []byte("Incorrect authorization"))
		h.log.HTTP(http.StatusForbidden, r)
		return
	}

	h.handler.ServeHTTP(w, r)
}

// Wrap returns a new handler that transparently checks the secret bearer header
func (w SecretBearerAuthorizationWrapper) Wrap(h http.Handler) http.Handler {
	return wrappedHandler{
		secret:  string(w.Secret),
		handler: h,
		log:     w.Log,
	}
}
