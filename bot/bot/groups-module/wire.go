package groups

import "github.com/google/wire"

// ProvideModule provides everything needed to build a groups module
var ProvideModule = wire.NewSet(
	wire.Struct(new(Module), "*"),
)
