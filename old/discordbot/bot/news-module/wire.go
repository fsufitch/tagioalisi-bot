package news

import "github.com/google/wire"

// ProvideModule provides everything needed to build a news module
var ProvideModule = wire.NewSet(
	wire.Struct(new(Module), "*"),
)
