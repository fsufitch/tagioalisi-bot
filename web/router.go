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
	hello *HelloHandler,
	sockpuppet *SockpuppetHandler,
	login *LoginHandler,
	authCode *AuthCodeHandler,
	logout *LogoutHandler,
	whoami *WhoAmIHandler,
) Router {
	r := mux.NewRouter()
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Tighten this?
		handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Authorization", "Content-Type", "Content-Language", "Origin"}),
	)

	r.Handle("/", handlers.MethodHandler{"GET": hello})
	r.Handle("/sockpuppet", handlers.MethodHandler{"POST": sockpuppet})
	r.Handle("/login", handlers.MethodHandler{"GET": login})
	r.Handle("/login/redirect", handlers.MethodHandler{"GET": authCode})
	r.Handle("/logout", handlers.MethodHandler{"GET": logout})
	r.Handle("/whoami", handlers.MethodHandler{"GET": whoami})

	return Router(cors(r))
}
