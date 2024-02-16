// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/fsufitch/tagioalisi-bot/azure"
	"github.com/fsufitch/tagioalisi-bot/bot"
	"github.com/fsufitch/tagioalisi-bot/bot/dice-module"
	"github.com/fsufitch/tagioalisi-bot/bot/dice-module/calc"
	"github.com/fsufitch/tagioalisi-bot/bot/dictionary-module"
	"github.com/fsufitch/tagioalisi-bot/bot/groups-module"
	log2 "github.com/fsufitch/tagioalisi-bot/bot/log-module"
	"github.com/fsufitch/tagioalisi-bot/bot/memelink-module"
	"github.com/fsufitch/tagioalisi-bot/bot/news-module"
	"github.com/fsufitch/tagioalisi-bot/bot/ping-module"
	"github.com/fsufitch/tagioalisi-bot/bot/sockpuppet-module"
	"github.com/fsufitch/tagioalisi-bot/bot/wiki-module"
	"github.com/fsufitch/tagioalisi-bot/bot/wiki-module/wikisupport"
	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/db/acl-dao"
	"github.com/fsufitch/tagioalisi-bot/db/connection"
	"github.com/fsufitch/tagioalisi-bot/db/memes-dao"
	"github.com/fsufitch/tagioalisi-bot/grpc"
	"github.com/fsufitch/tagioalisi-bot/log"
	"github.com/fsufitch/tagioalisi-bot/proto"
	"github.com/fsufitch/tagioalisi-bot/security"
	"github.com/fsufitch/tagioalisi-bot/web"
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
	managedGroupPrefix := config.ProvideManagedGroupPrefixFromEnvironment()
	groupsModule := &groups.Module{
		Log:    logger,
		Prefix: managedGroupPrefix,
	}
	multi := _wireMultiValue
	wikiModule := &wiki.Module{
		Log:         logger,
		WikiSupport: multi,
	}
	diceCalculator := calc.DiceCalculator{
		Log: logger,
	}
	oAuth2Config := config.ProvideOAuth2ConfigFromEnvironment()
	applicationID := config.ProvideApplicationIDFromOAuth2Config(oAuth2Config)
	diceModule := &dice.Module{
		Log:        logger,
		Calculator: diceCalculator,
		AppID:      applicationID,
	}
	azureNewsSearchAPIKey := config.ProvideAzureCredentialsFromEnvironment()
	userAgent := config.ProvideUserAgent()
	bingNewsSearch := azure.BingNewsSearch{
		Log:       logger,
		Key:       azureNewsSearchAPIKey,
		UserAgent: userAgent,
	}
	newsModule := &news.Module{
		Log:   logger,
		News:  bingNewsSearch,
		AppID: applicationID,
	}
	merriamWebsterAPIKey := config.ProvideMerriamWebsterAPIKeyFromEnvironment()
	basicClient, err := dictionary.NewClient(merriamWebsterAPIKey, userAgent)
	if err != nil {
		cleanup()
		return Main{}, nil, err
	}
	dictionaryModule := &dictionary.Module{
		Log:    logger,
		Client: basicClient,
	}
	modules := bot.Modules{
		Ping:       module,
		Log:        logModule,
		SockPuppet: sockpuppetModule,
		MemeLink:   memelinkModule,
		Groups:     groupsModule,
		Wiki:       wikiModule,
		Dice:       diceModule,
		News:       newsModule,
		Dictionary: dictionaryModule,
	}
	moduleList := bot.ProvideModuleList(modules)
	botModuleBlacklist := config.ProvideBotModuleBlacklistFromEnvironment()
	discordBotToken, err := config.ProvideDiscordBotTokenFromEnvironment()
	if err != nil {
		cleanup()
		return Main{}, nil, err
	}
	tagioalisiBot := &bot.TagioalisiBot{
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
	botHTTPSPort, err := config.ProvideBotHTTPSPortFromEnvironment()
	if err != nil {
		cleanup()
		return Main{}, nil, err
	}
	botTLS := config.ProvideBotTLSFromEnvironment()
	launchTime := config.ProvideLaunchTime()
	helloHandler := &web.HelloHandler{
		Log:                logger,
		DebugMode:          debugMode,
		LaunchTime:         launchTime,
		BotModuleBlacklist: botModuleBlacklist,
		ManagedGroupPrefix: managedGroupPrefix,
		OAuth2Config:       oAuth2Config,
	}
	jwthmacSecret := config.ProvideJWTHMACSecretFromEnvironment()
	aesBlock, err := config.ProvideAESBlockFromEnvironment()
	if err != nil {
		cleanup()
		return Main{}, nil, err
	}
	aesSupport := security.AESSupport{
		Block: aesBlock,
	}
	jwtSupport := security.JWTSupport{
		JWTHMACSecret: jwthmacSecret,
		AES:           aesSupport,
	}
	sockpuppetHandler := &web.SockpuppetHandler{
		BotModule: sockpuppetModule,
		Log:       logger,
		JWT:       jwtSupport,
	}
	loginHandler := &web.LoginHandler{
		OAuth2Config: oAuth2Config,
		AES:          aesSupport,
	}
	authCodeHandler := &web.AuthCodeHandler{
		OAuth2Config: oAuth2Config,
		JWT:          jwtSupport,
		AES:          aesSupport,
	}
	logoutHandler := &web.LogoutHandler{}
	whoAmIHandler := &web.WhoAmIHandler{
		Log: logger,
		JWT: jwtSupport,
	}
	router := web.ProvideRouter(helloHandler, sockpuppetHandler, loginHandler, authCodeHandler, logoutHandler, whoAmIHandler)
	tagioalisiAPIServer := &web.TagioalisiAPIServer{
		Port:   botHTTPSPort,
		TLS:    botTLS,
		Log:    logger,
		Router: router,
	}
	grpcPort, err := config.ProvideGRPCPortFromEnvironment()
	if err != nil {
		cleanup()
		return Main{}, nil, err
	}
	unimplementedGreeterServer := proto.UnimplementedGreeterServer{}
	greeterServer := grpc.GreeterServer{
		UnimplementedGreeterServer: unimplementedGreeterServer,
		Log:                        logger,
	}
	tagioalisiGRPC := &grpc.TagioalisiGRPC{
		Log:           logger,
		Port:          grpcPort,
		GreeterServer: greeterServer,
	}
	mainMain, cleanup2, err := ProvideMain(interruptContext, tagioalisiBot, logger, debugMode, cliLoggingBootstrapper, tagioalisiAPIServer, tagioalisiGRPC)
	if err != nil {
		cleanup()
		return Main{}, nil, err
	}
	return mainMain, func() {
		cleanup2()
		cleanup()
	}, nil
}

var (
	_wireMultiValue = wikisupport.DefaultMultiWikiSupport
)
