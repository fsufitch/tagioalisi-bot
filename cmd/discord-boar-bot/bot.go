package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

// runBot starts the appropriate processes and blocks until either one crashes, or until interrupt
func (r *CLIRuntime) runBot() error {
	webErrChan := make(chan error, 1)
	if r.Configuration.WebEnabled {
		go func() {
			err := r.WebServer.Start()
			webErrChan <- err
		}()
	}

	botErrChan := make(chan error, 1)
	go func() {
		err := r.Bot.Start()
		botErrChan <- err
	}()

	r.Logger.Info("Processes initialized")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	var err error
	select {
	case err = <-webErrChan:
		r.Logger.Error("Shutting down due to fatal web server error")
		r.Logger.Error(err.Error())
	case err = <-botErrChan:
		r.Logger.Error("Shutting down due to fatal Discord bot error")
		r.Logger.Error(err.Error())
	case sig := <-sc:
		r.Logger.Error("Shutting down due to interrupt: " + sig.String())
	}

	r.Bot.Stop <- true
	<-time.After(1 * time.Second)
	r.CLILog.Stop <- true

	return err
}
