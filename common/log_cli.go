package common

import (
	"log"
	"os"
)

// CLILogModule is a pluggable module for outputting logs to the CLI
type CLILogModule struct {
	Stop      chan bool
	logModule *LoggerModule
	level     LogLevel
	listener  chan LogEntry
	stdLog    log.Logger
	errLog    log.Logger
}

func (m *CLILogModule) processEntry(entry LogEntry) {
	if entry.Level < m.level {
		return
	}

	if entry.Level >= LogError {
		m.errLog.Print(entry.Message)
	} else {
		m.stdLog.Print(entry.Message)
	}
}

func (m *CLILogModule) listen() {
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

// CreateCLILogModule creates a CLILogModule
func CreateCLILogModule(configuration *Configuration, logModule *LoggerModule) *CLILogModule {
	mod := CLILogModule{
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
