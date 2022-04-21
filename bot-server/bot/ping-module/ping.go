package ping

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagioalisi-bot/log"
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
		m.Log.Infof("ping module context done")
		cancel()
	}()
	return nil
}

func (m *Module) pingHandler(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Content == "!ping" {
		m.Log.Debugf("ping: pinged")
		if _, err := s.ChannelMessageSend(msg.ChannelID, "pong!"); err != nil {
			m.Log.Errorf("ping: could not send message: %v", err)
		}
	}
}
