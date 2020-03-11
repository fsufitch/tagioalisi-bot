package oauth

import "github.com/google/wire"

// ProviderSet contains the wire components necessaryt to set up OAuth support
var ProviderSet = wire.NewSet(
	wire.Value(States{}),
	ProvideMemorySessionStorage,
	wire.Bind(new(SessionStorage), new(MemorySessionStorage)),
)
