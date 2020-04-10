package wiki

import "github.com/google/wire"

// ProvideModule provides everything needed to build a wiki module
var ProvideModule = wire.NewSet(
	wire.Struct(new(Module), "*"),
)
