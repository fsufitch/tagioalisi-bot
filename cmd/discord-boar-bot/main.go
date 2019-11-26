package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/fsufitch/discord-boar-bot/bot"
	"github.com/fsufitch/discord-boar-bot/common"
	"github.com/fsufitch/discord-boar-bot/web"
	"github.com/google/wire"
)

func main() {
	runtime, err := InitializeCLIRuntime()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	err = runtime.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// CLIRuntime encapsulates a bot runtime in the CLI
type CLIRuntime struct {
	Configuration *common.Configuration
	Logger        *common.LogDispatcher
	CLILog        *common.CLILogReceiver
	WebServer     *web.BoarBotServer
	Bot           *bot.DiscordBoarBot
}

// Start runs the appropriate segment of code based on the run mode
func (r *CLIRuntime) Start() error {
	switch r.Configuration.RunMode {
	case common.Bot:
		return r.runBot()
	case common.Migration:
		return r.runMigration()
	case common.Unknown:
		fallthrough
	default:
		return errors.New("Unknown run mode found")
	}
}

// NewCLIRuntime creates a new CLIRuntime
func NewCLIRuntime(
	configuration *common.Configuration,
	logger *common.LogDispatcher,
	cliLog *common.CLILogReceiver, webServer *web.BoarBotServer,
	bot *bot.DiscordBoarBot,
) *CLIRuntime {
	return &CLIRuntime{
		Configuration: configuration,
		Logger:        logger,
		CLILog:        cliLog,
		WebServer:     webServer,
		Bot:           bot,
	}
}

// CLIRuntimeProviderSet is a wire ProviderSet with the necessities for CLI runtime
var CLIRuntimeProviderSet = wire.NewSet(
	NewCLIRuntime,
	common.ConfigurationFromEnvironment,
	common.NewLogDispatcher,
	common.NewCLILogReceiver,
)
