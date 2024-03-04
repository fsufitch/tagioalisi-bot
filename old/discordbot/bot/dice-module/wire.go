package dice

import (
	"github.com/fsufitch/tagioalisi-bot/bot/dice-module/calc"
	"github.com/google/wire"
)

// ProvideModule creates the wiki support provider set
var ProvideModule = wire.NewSet(
	calc.ProvideCalc,
	wire.Struct(new(Module), "*"),
)
