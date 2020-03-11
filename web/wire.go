package web

import (
	"github.com/google/wire"
)

// ProvideWebServer contains all the necessary wire providers to stand up a webserver
var ProvideWebServer = wire.NewSet(
	wire.Struct(new(TagioalisiAPIServer), "*"),
	wire.Struct(new(SecretBearerAuthorizationWrapper), "*"),
	ProvideRouter,
	wire.Struct(new(HelloHandler), "*"),
	wire.Struct(new(SockpuppetHandler), "*"),
	wire.Struct(new(LoginHandler), "*"),
	wire.Struct(new(AuthCodeHandler), "*"),
	wire.Struct(new(LogoutHandler), "*"),
)
