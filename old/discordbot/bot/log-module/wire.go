package log

import "github.com/google/wire"

// ProvideModule is a set for building the log module
var ProvideModule = wire.NewSet(
	wire.Struct(new(Module), "Log", "DebugMode", "LogChannel"),
)
