package sockpuppet

import "github.com/google/wire"

var ProvideModule = wire.NewSet(
	wire.Struct(new(Module), "Log"),
)
