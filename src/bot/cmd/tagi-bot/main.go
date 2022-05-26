package main

import (
	"context"
	"os"
	"time"

	"github.com/fsufitch/tagioalisi-bot/bot"
	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/grpc"
	"github.com/fsufitch/tagioalisi-bot/log"
)

var banner = []string{
	"+-------------------------------+",
	"| Discord Boar Bot - now in Go! |",
	"+-------------------------------+",
}

// Main is an initialized runtime with all necessary dependencies injected
type Main struct {
	context    context.Context
	log        *log.Logger
	bot        bot.Bot
	webRun     WebRunFunc
	grpcServer grpc.TagioalisiGRPCServer
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

	grpcError := make(chan error)
	go func() {
		grpcError <- m.grpcServer.Run()
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
func ProvideMain(
	ctx InterruptContext,
	bot bot.Bot,
	log *log.Logger,
	debugMode config.DebugMode,
	cliBS log.CLILoggingBootstrapper,
	webRun WebRunFunc,
	grpcServer grpc.TagioalisiGRPCServer,
) (Main, func(), error) {
	cliBS.Start()
	for _, line := range banner {
		log.Infof(line)
	}
	log.Infof("Debug mode: %v", debugMode)
	return Main{ctx, log, bot, webRun, grpcServer}, func() { cliBS.Stop() }, nil
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
