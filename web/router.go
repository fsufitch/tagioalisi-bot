package web

import (
	"github.com/gorilla/mux"
)

// Router is a router for our server
type Router *mux.Router

// ProvideRouter creates the router necessary to start the server
func ProvideRouter(
	security *SecretBearerAuthorizationWrapper,
	hello *HelloHandler,
	sockpuppet *SockpuppetHandler,
) Router {
	r := mux.NewRouter()

	r.Handle("/", hello)
	r.Handle("/sockpuppet", security.Wrap(sockpuppet))

	return Router(r)
}
