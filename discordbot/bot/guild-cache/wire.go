package guildcache

import "github.com/google/wire"

var ProvideModule = wire.NewSet(
	ProvideGuildCacheManager,
)
