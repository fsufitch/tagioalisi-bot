// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/fsufitch/discord-boar-bot/bot"
	log2 "github.com/fsufitch/discord-boar-bot/bot/log-module"
	"github.com/fsufitch/discord-boar-bot/bot/memelink-module"
	"github.com/fsufitch/discord-boar-bot/bot/ping-module"
	"github.com/fsufitch/discord-boar-bot/bot/sockpuppet-module"
	"github.com/fsufitch/discord-boar-bot/config"
	"github.com/fsufitch/discord-boar-bot/db/acl-dao"
	"github.com/fsufitch/discord-boar-bot/db/connection"
	"github.com/fsufitch/discord-boar-bot/db/memes-dao"
	"github.com/fsufitch/discord-boar-bot/log"
	"github.com/fsufitch/discord-boar-bot/web"
)

// Injectors from wire.go:

func InitializeMain() (Main, func(), error) {
	interruptContext := ProvideInterruptContext()
	logger := log.ProvideLogger()
	module := &ping.Module{
		Log: logger,
	}
	debugMode, err := config.ProvideDebugModeFromEnvironment()
	if err != nil {
		return Main{}, nil, err
	}
	discordLogChannel := config.ProvideDiscordLogChannelFromEnvironment()
	logModule := &log2.Module{
		Log:        logger,
		DebugMode:  debugMode,
		LogChannel: discordLogChannel,
	}
	sockpuppetModule := &sockpuppet.Module{
		Log: logger,
	}
	databaseString, err := config.ProvideDatabaseStringFromEnvironment()
	if err != nil {
		return Main{}, nil, err
	}
	databaseConnection, cleanup, err := connection.ProvidePostgresDatabaseConnection(logger, databaseString)
	if err != nil {
		return Main{}, nil, err
	}
	dao := &memes.DAO{
		Conn: databaseConnection,
	}
	aclDAO := &acl.DAO{
		Conn: databaseConnection,
	}
	memelinkModule := &memelink.Module{
		Log:     logger,
		MemeDAO: dao,
		ACLDAO:  aclDAO,
	}
	modules := bot.Modules{
		Ping:       module,
		Log:        logModule,
		SockPuppet: sockpuppetModule,
		MemeLink:   memelinkModule,
	}
	moduleList := bot.ProvideModuleList(modules)
	botModuleBlacklist := config.ProvideBotModuleBlacklistFromEnvironment()
	discordBotToken, err := config.ProvideDiscordBotTokenFromEnvironment()
	if err != nil {
		cleanup()
		return Main{}, nil, err
	}
	discordBoarBot := &bot.DiscordBoarBot{
		Log:             logger,
		Modules:         moduleList,
		ModuleBlacklist: botModuleBlacklist,
		Token:           discordBotToken,
	}
	stdOutReceiver := log.ProvideStdOutReceiver(debugMode)
	stdErrReceiver := log.ProvideStdErrReceiver(debugMode)
	cliLoggingBootstrapper := log.CLILoggingBootstrapper{
		Logger:         logger,
		StdOutReceiver: stdOutReceiver,
		StdErrReceiver: stdErrReceiver,
	}
	webEnabled, err := config.ProvideWebEnabledFromEnvironment()
	if err != nil {
		cleanup()
		return Main{}, nil, err
	}
	webPort, err := config.ProvideWebPortFromEnvironment()
	if err != nil {
		cleanup()
		return Main{}, nil, err
	}
	webSecret := config.ProvideWebSecretFromEnvironment()
	secretBearerAuthorizationWrapper := &web.SecretBearerAuthorizationWrapper{
		Secret: webSecret,
	}
	helloHandler := &web.HelloHandler{}
	sockpuppetHandler := &web.SockpuppetHandler{
		BotModule: sockpuppetModule,
	}
	router := web.ProvideRouter(secretBearerAuthorizationWrapper, helloHandler, sockpuppetHandler)
	boarBotServer := web.BoarBotServer{
		WebPort: webPort,
		Log:     logger,
		Router:  router,
	}
	webRunFunc := ProvideWebRunFunc(webEnabled, boarBotServer)
	mainMain, cleanup2, err := ProvideMain(interruptContext, discordBoarBot, logger, debugMode, cliLoggingBootstrapper, webRunFunc)
	if err != nil {
		cleanup()
		return Main{}, nil, err
	}
	return mainMain, func() {
		cleanup2()
		cleanup()
	}, nil
}
