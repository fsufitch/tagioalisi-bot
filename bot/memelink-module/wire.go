package memelink

import "github.com/google/wire"

// ProvideMemeLinkModule provides everything needed to build a meme link module
var ProvideModule = wire.NewSet(
	wire.Struct(new(Module), "Log", "MemeDAO"),
)
