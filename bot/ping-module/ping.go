package ping

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/discord-boar-bot/log"
)

// Module is a bot module that responds to "!ping" with "!pong"
type Module struct {
	Log *log.Logger
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "ping" }

// Register adds this module to the Discord session
func (m *Module) Register(ctx context.Context, session *discordgo.Session) error {
	cancel := session.AddHandler(m.pingHandler)
	go func() {
		<-ctx.Done()
		m.Log.Warningf("ping module context done")
		cancel()
	}()
	return nil
}

func (m *Module) pingHandler(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Content == "!ping" {
		m.Log.Debugf("ping received")
		s.ChannelMessageSend(msg.ChannelID, "pong!")
	}
}
