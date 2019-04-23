package main

import (
	"github.com/fsufitch/discord-boar-bot/common"
	"github.com/google/wire"
)

// CLIRuntime encapsulates a bot runtime in the CLI
type CLIRuntime struct {
	Configuration *common.Configuration
	Logger        *common.LoggerModule
}

// NewCLIRuntime creates a new CLIRuntime
func NewCLIRuntime(configuration *common.Configuration, logger *common.LoggerModule, _ *common.CLILogModule) *CLIRuntime {
	return &CLIRuntime{
		Configuration: configuration,
		Logger:        logger,
	}
}

// CLIRuntimeProviderSet is a wire ProviderSet with the bare necessities for CLI runtime
var CLIRuntimeProviderSet = wire.NewSet(
	NewCLIRuntime,
	common.ConfigurationFromEnvironment,
	common.NewLoggerModule,
	common.CreateCLILogModule,
)
