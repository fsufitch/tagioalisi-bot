//+build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeCLIRuntime() (*CLIRuntime, error) {
	wire.Build(CLIRuntimeProviderSet)

	return &CLIRuntime{}, nil
}
