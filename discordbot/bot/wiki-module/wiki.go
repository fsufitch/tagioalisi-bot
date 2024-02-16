package wiki

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/fsufitch/tagioalisi-bot/bot/wiki-module/wikisupport"
	"github.com/fsufitch/tagioalisi-bot/log"
)

// Module is a bot module that responds to "!wiki" commands
type Module struct {
	Log         *log.Logger
	WikiSupport wikisupport.Multi
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "wiki" }

// Register adds this module to the Discord session
func (m *Module) Register(ctx context.Context, session *discordgo.Session) error {
	cancel := session.AddHandler(m.handleCommand)
	go func() {
		<-ctx.Done()
		m.Log.Infof("wiki module context done")
		cancel()
	}()
	return nil
}

func (m *Module) RegisterGuild(ctx context.Context, session *discordgo.Session, guildID string) error {
	return nil
}
