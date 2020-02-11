package main

import (
	"os"
	"time"

	"github.com/fsufitch/discord-boar-bot/log"
)

type Main struct {
	Log              *log.Logger
	MigrationWrapper MigrationWrapper
}

func (m Main) Main() int {
	if err := m.MigrationWrapper.Run(); err != nil {
		m.Log.Criticalf("error running migration: %v", err)
		return 1
	}
	m.Log.Infof("migration successful")
	return 0
}

func ProvideMain(log *log.Logger, cliBS log.CLILoggingBootstrapper, mw MigrationWrapper) (Main, func(), error) {
	cliBS.Start()

	return Main{log, mw}, func() { cliBS.Stop() }, nil
}

func main() {
	main, cleanup, err := InitializeMain()
	if err != nil {
		panic(err)
	}

	code := main.Main()
	cleanup()
	<-time.After(500 * time.Millisecond)
	os.Exit(code)
}
