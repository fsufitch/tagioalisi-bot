package news

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/fsufitch/tagioalisi-bot/azure"
	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/log"
)

// Module is a news module implementing RegisterableModule
type Module struct {
	Log   *log.Logger
	News  azure.BingNewsSearch
	AppID config.ApplicationID
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "news" }

// Register adds this module to the Discord session
func (m *Module) Register(ctx context.Context, session *discordgo.Session) error {
	return nil
}

func (m *Module) RegisterGuild(ctx context.Context, session *discordgo.Session, guildID string) error {
	if err := m.RegisterApplicationCommand(ctx, session, guildID); err != nil {
		return err
	}
	return nil
}
