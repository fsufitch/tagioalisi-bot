// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/fsufitch/discord-boar-bot/config"
	"github.com/fsufitch/discord-boar-bot/log"
	"github.com/google/wire"
)

func InitializeMain() (Main, func(), error) {
	panic(wire.Build(
		wire.Struct(new(MigrationWrapper), "*"),
		ProvideMain,
		log.CLILoggingProviderSet,
		config.EnvironmentProviderSet,
	))
}
