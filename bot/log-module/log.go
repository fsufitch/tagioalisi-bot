package log

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/discord-boar-bot/common"
)

// Module is a bot module that outputs log messages
type Module struct {
	session      *discordgo.Session
	log          *common.LogDispatcher
	logLevel     common.LogLevel
	logChannelID string
	listener     chan common.LogEntry
	Stop         chan bool
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "log" }

// Register adds this module to the Discord session
func (m *Module) Register(session *discordgo.Session) error {
	// TODO: "log here" handler with permissions
	m.session = session

	if m.logChannelID != "" {
		m.log.AddListener(m.listener)
		go m.listen()
		m.log.Info("Enabled discord logging")
	} else {
		m.log.Warn("Discord logging channel ID not set")
	}
	return nil
}

func (m *Module) send(message string) {
	if _, err := m.session.ChannelMessageSend(m.logChannelID, message); err != nil {
		m.Stop <- true
		m.log.Error("Could not print to discord channel log, stopping discord logging")
	}
}

func (m *Module) listen() {
	stopped := false
	for !stopped {
		select {
		case entry := <-m.listener:
			if entry.Level >= m.logLevel {
				m.send(entry.Message)
			}
		case <-m.Stop:
			stopped = true
			m.log.Warn("CLI log module stopped")
		}
	}
	m.log.RemoveListener(m.listener)
}

// NewModule creates a new ping handling module
func NewModule(config *common.Configuration, log *common.LogDispatcher) *Module {
	return &Module{
		log:          log,
		logLevel:     config.DiscordLogLevel,
		logChannelID: config.DiscordLogChannel,
		listener:     make(chan common.LogEntry),
		Stop:         make(chan bool, 1),
	}
}
