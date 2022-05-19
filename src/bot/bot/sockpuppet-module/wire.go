package sockpuppet

import "github.com/google/wire"

// ProvideModule specifies everything necessary to build the sockpuppet module
var ProvideModule = wire.NewSet(
	wire.Struct(new(Module), "Log"),
)
