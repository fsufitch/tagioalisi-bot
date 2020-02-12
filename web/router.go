package web

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Router is a router for our server
type Router http.Handler

// ProvideRouter creates the router necessary to start the server
func ProvideRouter(
	security *SecretBearerAuthorizationWrapper,
	hello *HelloHandler,
	sockpuppet *SockpuppetHandler,
) Router {
	r := mux.NewRouter()
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Tighten this?
		handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Authorization", "Content-Type", "Content-Language", "Origin"}),
	)

	r.Handle("/", handlers.MethodHandler{"GET": hello})
	r.Handle("/sockpuppet", handlers.MethodHandler{"POST": security.Wrap(sockpuppet)})

	return Router(cors(r))
}
