package bot

import (
	"github.com/fsufitch/tagioalisi-bot/bot/groups-module"
	"github.com/fsufitch/tagioalisi-bot/bot/log-module"
	"github.com/fsufitch/tagioalisi-bot/bot/memelink-module"
	"github.com/fsufitch/tagioalisi-bot/bot/ping-module"
	"github.com/fsufitch/tagioalisi-bot/bot/sockpuppet-module"
	"github.com/google/wire"
)

// ProvideTagioalisiBot contains all the necessary wire providers to stand up a Tagioalisi Bot
var ProvideTagioalisiBot = wire.NewSet(
	wire.Struct(new(TagioalisiBot), "*"),
	wire.Bind(new(Bot), new(*TagioalisiBot)),
	wire.Struct(new(Modules), "*"),
	ProvideModuleList,
	ping.ProvideModule,
	sockpuppet.ProvideModule,
	memelink.ProvideModule,
	log.ProvideModule,
	groups.ProvideModule,
)
