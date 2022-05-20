// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/log"
)

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Injectors from wire.go:

func InitializeMain() (Main, func(), error) {
	logger := log.ProvideLogger()
	debugMode, err := config.ProvideDebugModeFromEnvironment()
	if err != nil {
		return Main{}, nil, err
	}
	stdOutReceiver := log.ProvideStdOutReceiver(debugMode)
	stdErrReceiver := log.ProvideStdErrReceiver(debugMode)
	cliLoggingBootstrapper := log.CLILoggingBootstrapper{
		Logger:         logger,
		StdOutReceiver: stdOutReceiver,
		StdErrReceiver: stdErrReceiver,
	}
	databaseString, err := config.ProvideDatabaseStringFromEnvironment()
	if err != nil {
		return Main{}, nil, err
	}
	migrationWrapper := MigrationWrapper{
		Log:   logger,
		DBURL: databaseString,
	}
	mainMain, cleanup, err := ProvideMain(logger, cliLoggingBootstrapper, migrationWrapper)
	if err != nil {
		return Main{}, nil, err
	}
	return mainMain, func() {
		cleanup()
	}, nil
}
