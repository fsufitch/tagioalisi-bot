// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/fsufitch/discord-boar-bot/bot"
	"github.com/fsufitch/discord-boar-bot/config"
	"github.com/fsufitch/discord-boar-bot/db"
	"github.com/fsufitch/discord-boar-bot/log"
	"github.com/fsufitch/discord-boar-bot/web"
	"github.com/google/wire"
)

func InitializeMain() (Main, func(), error) {
	panic(wire.Build(
		ProvideMain,
		ProvideInterruptContext,
		ProvideWebRunFunc,
		config.EnvironmentProviderSet,
		log.CLILoggingProviderSet,
		bot.ProvideDiscordBoarBot,
		db.ProvidePostgresDatabase,
		web.ProvideWebServer,
	))
}
