package dictionary

import (
	mwdict "github.com/fsufitch/tagioalisi-bot/merriam-webster"
	"github.com/google/wire"
)

// ProvideModule contains the providers necessary to assemble the dictionary module
var ProvideModule = wire.NewSet(
	wire.Struct(new(Module), "*"),
	wire.Bind(new(Client), new(*mwdict.BasicClient)),
	NewClient,
)
