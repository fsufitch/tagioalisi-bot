package bot

import (
	"github.com/fsufitch/discord-boar-bot/bot/ping-module"
	"github.com/fsufitch/discord-boar-bot/bot/sockpuppet-module"
	"github.com/google/wire"
)

// BotProviderSet contains all the necessary wire providers to stand up a Discord Boar Bot
var BotProviderSet = wire.NewSet(
	NewDiscordBoarBot,
	InitModuleRegistry,
	ping.NewModule,
	sockpuppet.NewModule,
)
