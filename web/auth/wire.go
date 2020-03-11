package auth

import "github.com/google/wire"

var ProvideWebAuth = wire.NewSet(
	wire.Struct(new(JWTSupport), "*"),
	wire.Struct(new(CookieSupport), "*"),
	wire.Value(LoginStates{}),
	ProvideMemorySessionStorage,
	wire.Bind(new(SessionStorage), new(MemorySessionStorage)),
)
