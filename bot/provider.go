package bot

import (
	"github.com/fsufitch/discord-boar-bot/bot/log-module"
	"github.com/fsufitch/discord-boar-bot/bot/memelink-module"
	"github.com/fsufitch/discord-boar-bot/bot/ping-module"
	"github.com/fsufitch/discord-boar-bot/bot/sockpuppet-module"
	"github.com/fsufitch/discord-boar-bot/db"
	"github.com/google/wire"
)

// BotProviderSet contains all the necessary wire providers to stand up a Discord Boar Bot
var BotProviderSet = wire.NewSet(
	NewDiscordBoarBot,
	InitModuleRegistry,
	ping.NewModule,
	sockpuppet.NewModule,
	memelink.NewModule,
	log.NewModule,
	db.DBProviderSet, // Database access *required* for bot functionality
)
