package groups

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/fsufitch/tagioalisi-bot/bot/util"
	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/log"
)

// Module is a bot module that responds to "!groups" commands
type Module struct {
	Log           *log.Logger
	Prefix        config.ManagedGroupPrefix
	ApplicationID config.ApplicationID
	InterUtil     util.InteractionUtil
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "groups" }

// Register adds this module to the Discord session
func (m *Module) Register(ctx context.Context, session *discordgo.Session) error {
	return nil
}

func (m *Module) RegisterGuild(ctx context.Context, session *discordgo.Session, guildID string) error {
	return m.RegisterApplicationCommand(ctx, session, guildID)
}
