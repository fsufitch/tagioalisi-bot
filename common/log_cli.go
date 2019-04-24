package common

import (
	"log"
	"os"
)

// CLILogReceiver is a pluggable module for outputting logs to the CLI
type CLILogReceiver struct {
	Stop      chan bool
	logModule *LogDispatcher
	level     LogLevel
	listener  chan LogEntry
	stdLog    log.Logger
	errLog    log.Logger
}

func (m *CLILogReceiver) processEntry(entry LogEntry) {
	if entry.Level < m.level {
		return
	}

	if entry.Level >= LogError {
		m.errLog.Print(entry.Message)
	} else {
		m.stdLog.Print(entry.Message)
	}
}

func (m *CLILogReceiver) listen() {
	stopped := false
	for !stopped {
		select {
		case entry := <-m.listener:
			m.processEntry(entry)
		case <-m.Stop:
			stopped = true
			m.errLog.Print("CLI log module stopped")
		}
	}
	m.logModule.RemoveListener(m.listener)
}

// NewCLILogReceiver creates a CLILogReceiver
func NewCLILogReceiver(configuration *Configuration, logModule *LogDispatcher) *CLILogReceiver {
	mod := CLILogReceiver{
		level:     configuration.CLILogLevel,
		logModule: logModule,
		listener:  make(chan LogEntry),
		Stop:      make(chan bool, 1),
		stdLog:    log.Logger{},
		errLog:    log.Logger{},
	}

	mod.stdLog.SetOutput(os.Stdout)
	mod.stdLog.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	mod.errLog.SetOutput(os.Stderr)
	mod.errLog.SetFlags(log.Ldate | log.Ltime | log.LUTC)

	go mod.listen()
	logModule.AddListener(mod.listener)

	logModule.Info("CLI log output set up")

	return &mod
}
