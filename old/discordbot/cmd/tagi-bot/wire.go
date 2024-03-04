//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/fsufitch/tagioalisi-bot/azure"
	"github.com/fsufitch/tagioalisi-bot/bot"
	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/db"
	"github.com/fsufitch/tagioalisi-bot/grpc"
	"github.com/fsufitch/tagioalisi-bot/log"
	"github.com/fsufitch/tagioalisi-bot/security"
	"github.com/fsufitch/tagioalisi-bot/web"
	"github.com/google/wire"
)

func InitializeMain() (Main, func(), error) {
	panic(wire.Build(
		ProvideMain,
		ProvideInterruptContext,
		config.EnvironmentProviderSet,
		log.CLILoggingProviderSet,
		azure.AzureProviderSet,
		bot.ProvideTagioalisiBot,
		grpc.GRPCProviderSet,
		db.ProvidePostgresDatabase,
		web.ProvideWebServer,
		security.ProvideSecurity,
	))
}
