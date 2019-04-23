package web

import "github.com/gorilla/mux"

// NewRouter creates the router necessary to start the server
func NewRouter(
	security *SecretBearerAuthorizationWrapper,
	hello *HelloHandler,
	sockpuppet *SockpuppetHandler,
) *mux.Router {
	r := mux.NewRouter()

	r.Handle("/", hello)
	r.Handle("/sockpuppet", security.Wrap(sockpuppet))

	return r
}
