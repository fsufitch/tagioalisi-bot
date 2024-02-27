package bot

import (
	"github.com/google/wire"

	"github.com/fsufitch/tagioalisi-bot/bot/dice-module"
	"github.com/fsufitch/tagioalisi-bot/bot/dictionary-module"
	groupscommand "github.com/fsufitch/tagioalisi-bot/bot/groups-command"
	"github.com/fsufitch/tagioalisi-bot/bot/groups-module"
	guildcache "github.com/fsufitch/tagioalisi-bot/bot/guild-cache"
	"github.com/fsufitch/tagioalisi-bot/bot/log-module"
	"github.com/fsufitch/tagioalisi-bot/bot/memelink-module"
	"github.com/fsufitch/tagioalisi-bot/bot/news-module"
	"github.com/fsufitch/tagioalisi-bot/bot/ping-module"
	"github.com/fsufitch/tagioalisi-bot/bot/sockpuppet-module"
	"github.com/fsufitch/tagioalisi-bot/bot/util"
	"github.com/fsufitch/tagioalisi-bot/bot/wiki-module"
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
	guildcache.ProvideModule,
	wiki.ProvideModule,
	dice.ProvideModule,
	news.ProvideModule,
	dictionary.ProvideModule,
	util.ProvideModule,
	ProvideApplicationCommandModuleBootstrapper,
	groupscommand.ProvideModule,
)
