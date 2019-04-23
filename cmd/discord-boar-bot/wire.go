//+build wireinject

package main

import (
	"github.com/fsufitch/discord-boar-bot/bot"
	"github.com/fsufitch/discord-boar-bot/web"
	"github.com/google/wire"
)

func InitializeCLIRuntime() (*CLIRuntime, error) {
	wire.Build(
		CLIRuntimeProviderSet,
		web.WebProviderSet,
		bot.BotProviderSet,
	)

	return &CLIRuntime{}, nil
}
