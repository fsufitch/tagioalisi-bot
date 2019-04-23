// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/fsufitch/discord-boar-bot/bot"
	"github.com/fsufitch/discord-boar-bot/common"
	"github.com/fsufitch/discord-boar-bot/web"
)

// Injectors from wire.go:

func InitializeCLIRuntime() (*CLIRuntime, error) {
	configuration, err := common.ConfigurationFromEnvironment()
	if err != nil {
		return nil, err
	}
	loggerModule := common.NewLoggerModule()
	cliLogModule := common.CreateCLILogModule(configuration, loggerModule)
	boarBotServer := web.NewBoarBotServer(configuration, loggerModule)
	discordBoarBot := bot.NewDiscordBoarBot(configuration, loggerModule)
	cliRuntime := NewCLIRuntime(configuration, loggerModule, cliLogModule, boarBotServer, discordBoarBot)
	return cliRuntime, nil
}
