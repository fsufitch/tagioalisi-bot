package sockpuppet

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagialisi-bot/log"
	"github.com/pkg/errors"
)

// Module is a bot module that sockpuppets messages from elsewhere
type Module struct {
	session *discordgo.Session
	done    bool
	Log     *log.Logger
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "sockpuppet" }

// Register adds this module to the Discord session
func (m *Module) Register(ctx context.Context, session *discordgo.Session) error {
	m.session = session
	go func() {
		<-ctx.Done()
		m.Log.Infof("sockpuppet module context done")
		m.done = true
	}()
	return nil
}

// SendMessage is used to send a message via the sockpuppet
func (m *Module) SendMessage(channelID string, message string) error {
	if m.session == nil {
		return errors.New("No session to send messages through")
	}

	if _, err := m.session.ChannelMessageSend(channelID, message); err != nil {
		return errors.Wrap(err, "sockpuppet: could not send message")
	}
	return nil
}
