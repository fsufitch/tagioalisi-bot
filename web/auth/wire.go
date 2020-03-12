package auth

import "github.com/google/wire"

// ProvideWebAuth is a wire set containing code necessary to manage auth sessions
var ProvideWebAuth = wire.NewSet(
	wire.Value(LoginStates{}),
	ProvideMemorySessionStorage,
	wire.Bind(new(SessionStorage), new(MemorySessionStorage)),
)
