package sockpuppet

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

// Module is a bot module that sockpuppets messages from elsewhere
type Module struct {
	session *discordgo.Session
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "sockpuppet" }

// Register adds this module to the Discord session
func (m *Module) Register(session *discordgo.Session) error {
	m.session = session
	// Nothing to do here, it does not react to anything
	return nil
}

// NewModule creates a new sockpuppet module
func NewModule() *Module {
	return &Module{}
}

// SendMessage is used to send a message via the sockpuppet
func (m *Module) SendMessage(channelID string, message string) error {
	if m.session == nil {
		return errors.New("No session to send messages through")
	}

	_, err := m.session.ChannelMessageSend(channelID, message)
	return err
}
