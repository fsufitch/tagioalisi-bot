package ping

import "github.com/bwmarrin/discordgo"

// Module is a bot module that responds to "!ping" with "!pong"
type Module struct{}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "ping" }

// Register adds this module to the Discord session
func (m *Module) Register(session *discordgo.Session) error {
	session.AddHandler(m.pingHandler)
	return nil
}

// NewModule creates a new ping handling module
func NewModule() *Module {
	return &Module{}
}

func (m *Module) pingHandler(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Content == "!ping" {
		s.ChannelMessageSend(msg.ChannelID, "pong!")
	}
}
