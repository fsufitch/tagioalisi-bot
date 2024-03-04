package wikisupport

import "github.com/google/wire"

// ProvideMulti creates the wiki support provider set
var ProvideMulti = wire.NewSet(
	wire.Value(DefaultMultiWikiSupport),
)
