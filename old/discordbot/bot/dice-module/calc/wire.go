package calc

import "github.com/google/wire"

// ProvideCalc creates the calculator module providers
var ProvideCalc = wire.NewSet(
	wire.Struct(new(DiceCalculator), "*"),
)
