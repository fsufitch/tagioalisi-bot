package log

import (
	"os"

	"github.com/fsufitch/tagioalisi-bot/config"
)

// StdOutReceiver is a receiver to print to stdout
type StdOutReceiver MessageReceiver

// StdErrReceiver WriterMessageReceiver
type StdErrReceiver MessageReceiver

// ProvideStdOutReceiver creates a StdOutReceiver for dependency injection
func ProvideStdOutReceiver(debugMode config.DebugMode) StdOutReceiver {
	receiver := writerMessageReceiver{Writer: os.Stdout}
	if debugMode {
		receiver.MinLevel = Debug
		receiver.MaxLevel = Warning
	} else {
		receiver.MinLevel = Info
		receiver.MaxLevel = Warning
	}
	return receiver
}

// ProvideStdErrReceiver creates a StdErrReceiver for dependency injection
func ProvideStdErrReceiver(debugMode config.DebugMode) StdErrReceiver {
	return writerMessageReceiver{
		Writer:   os.Stderr,
		MinLevel: Error,
		MaxLevel: Critical,
	}
}

// CLILoggingBootstrapper handles stdout/stderr logging setup
type CLILoggingBootstrapper struct {
	Logger         *Logger
	StdOutReceiver StdOutReceiver
	StdErrReceiver StdErrReceiver
}

// Start begins logging to stdout and stderr
func (b CLILoggingBootstrapper) Start() (err error) {
	err = b.Logger.RegisterReceiver("stdout", b.StdOutReceiver)
	if err != nil {
		return err
	}
	err = b.Logger.RegisterReceiver("stderr", b.StdErrReceiver)
	if err != nil {
		return err
	}
	b.Logger.Infof("CLI logging set up")
	return
}

// Stop stops logging to stdout and stderr
func (b CLILoggingBootstrapper) Stop() (err error) {
	b.Logger.Warningf("dismantling CLI logging")
	err = b.Logger.DeregisterReceiver("stdout")
	if err != nil {
		return err
	}
	return b.Logger.DeregisterReceiver("stderr")
}
