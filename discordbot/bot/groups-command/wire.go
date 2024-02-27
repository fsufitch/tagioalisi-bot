package groupscommand

import "github.com/google/wire"

var ProvideModule = wire.NewSet(
	ProvideApplicationCommand,
	ProvideDefaultPrefixer,
)
