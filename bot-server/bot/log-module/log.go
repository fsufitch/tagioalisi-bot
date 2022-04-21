package log

import (
	"context"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/log"
)

// Module is a bot module that outputs log messages
type Module struct {
	Log        *log.Logger
	DebugMode  config.DebugMode
	LogChannel config.DiscordLogChannel

	session   *discordgo.Session
	sendQueue chan log.Message
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "log" }

// Register adds this module to the Discord session
func (m *Module) Register(ctx context.Context, session *discordgo.Session) error {
	// TODO: "log here" handler with permissions
	m.session = session

	if m.LogChannel != "" {
		err := m.Log.RegisterReceiver("discord", m)
		if err != nil {
			return err
		}
		m.sendQueue = make(chan log.Message, 64)
		go m.sendWorker()
		return nil
	}
	m.Log.Warningf("log: Discord logging channel ID not set, cannot log to Discord")
	return nil
}

// Receive implements log.MessageReceiver
func (m Module) Receive(messageChan <-chan log.Message) {
	for message := range messageChan {
		select {
		case m.sendQueue <- message:
			// OK
		default:
			// Oh no
			fmt.Fprint(os.Stderr, "XXX: unable to queue log message to discord\n")
		}
	}
}

func (m Module) sendWorker() {
	for message := range m.sendQueue {
		if m.DebugMode && message.Level < log.Info {
			continue
		}
		if !m.DebugMode && message.Level < log.Warning {
			continue
		}

		text := format(message)

		if _, err := m.session.ChannelMessageSend(string(m.LogChannel), text); err != nil {
			// oh boy
			fmt.Fprint(os.Stderr, "XXX: unable to send log message to discord\n")
		}
	}
}

func format(message log.Message) string {
	var levelPrefix string
	switch message.Level {
	case log.Debug:
		levelPrefix = "[DEBUG]"
	case log.Info:
		levelPrefix = "[INFO]"
	case log.Warning:
		levelPrefix = "[WARNING]"
	case log.Error:
		levelPrefix = "[ERROR]"
	case log.Critical:
		levelPrefix = "[CRITICAL]"
	}
	return fmt.Sprintf("**%s** %s", levelPrefix, fmt.Sprintf(message.Format, message.Values...))
}
