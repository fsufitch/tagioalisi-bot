package util

import "github.com/google/wire"

var ProvideModule = wire.NewSet(
	wire.Struct(new(InteractionUtil), "*"),
)
