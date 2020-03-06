// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/fsufitch/tagialisi-bot/bot"
	"github.com/fsufitch/tagialisi-bot/config"
	"github.com/fsufitch/tagialisi-bot/db"
	"github.com/fsufitch/tagialisi-bot/log"
	"github.com/fsufitch/tagialisi-bot/web"
	"github.com/google/wire"
)

func InitializeMain() (Main, func(), error) {
	panic(wire.Build(
		ProvideMain,
		ProvideInterruptContext,
		ProvideWebRunFunc,
		config.EnvironmentProviderSet,
		log.CLILoggingProviderSet,
		bot.ProvideTagioalisiBot,
		db.ProvidePostgresDatabase,
		web.ProvideWebServer,
	))
}
