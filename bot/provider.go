package bot

import "github.com/google/wire"

// BotProviderSet contains all the necessary wire providers to stand up a Discord Boar Bot
var BotProviderSet = wire.NewSet(
	NewDiscordBoarBot,
)
