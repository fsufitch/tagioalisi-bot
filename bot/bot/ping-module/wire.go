package ping

import "github.com/google/wire"

// ProvideModule contains the providers necessary to assemble the ping module
var ProvideModule = wire.NewSet(
	wire.Struct(new(Module), "*"),
)
