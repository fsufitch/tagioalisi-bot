package main

import (
	"context"
	"os"
	"time"

	"github.com/fsufitch/tagialisi-bot/bot"
	"github.com/fsufitch/tagialisi-bot/config"
	"github.com/fsufitch/tagialisi-bot/log"
)

var banner = []string{
	"+-------------------------------+",
	"| Discord Boar Bot - now in Go! |",
	"+-------------------------------+",
}

// Main is an initialized runtime with all necessary dependencies injected
type Main struct {
	context context.Context
	log     *log.Logger
	bot     bot.Bot
	webRun  WebRunFunc
}

// Main is what it says on the tin
func (m Main) Main() int {
	botError := make(chan error)
	go func() {
		botError <- m.bot.Run(m.context)
	}()

	webError := make(chan error)
	go func() {
		if m.webRun != nil {
			webError <- m.webRun()
		} else {
			m.log.Infof("web server disabled, not starting")
		}
	}()

	select {
	case err := <-botError:
		m.log.Criticalf("critical bot error: %v", err)
		return 1
	case err := <-webError:
		m.log.Criticalf("critical web error: %v", err)
		return 1
	case <-m.context.Done():
		m.log.Infof("main context canceled, shutting down")
		return 0
	}
}

// ProvideMain initializes the main process
func ProvideMain(ctx InterruptContext, bot bot.Bot, log *log.Logger, debugMode config.DebugMode, cliBS log.CLILoggingBootstrapper, webRun WebRunFunc) (Main, func(), error) {
	cliBS.Start()
	for _, line := range banner {
		log.Infof(line)
	}
	log.Infof("Debug mode: %v", debugMode)
	return Main{ctx, log, bot, webRun}, func() { cliBS.Stop() }, nil
}

func main() {
	m, cleanup, err := InitializeMain()
	if err != nil {
		panic(err)
	}
	defer cleanup()
	status := m.Main()
	<-time.After(500 * time.Millisecond)
	os.Exit(status)
}
