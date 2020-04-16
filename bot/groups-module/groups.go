package groups

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/log"
	"github.com/pkg/errors"
)

// Module is a bot module that responds to "!groups" commands
type Module struct {
	Log    *log.Logger
	Prefix config.ManagedGroupPrefix
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "groups" }

// Register adds this module to the Discord session
func (m *Module) Register(ctx context.Context, session *discordgo.Session) error {
	cancel := session.AddHandler(m.handleCommand)
	go func() {
		<-ctx.Done()
		m.Log.Infof("groups module context done")
		cancel()
	}()
	return nil
}

func (m Module) isGroupManager(session *discordgo.Session, event *discordgo.MessageCreate) (bool, error) {
	member, err := session.GuildMember(event.GuildID, event.Author.ID)
	if err != nil {
		return false, errors.Wrapf(err, "could not convert user to member %v", event.Author.ID)
	}

	allRoles, err := session.GuildRoles(event.GuildID)
	if err != nil {
		return false, errors.Wrap(err, "could not retrieve existing roles")
	}
	roleMap := map[string]*discordgo.Role{}
	for _, role := range allRoles {
		roleMap[role.ID] = role
	}

	for _, roleID := range member.Roles {
		// https://discordapp.com/developers/docs/topics/permissions#permissions
		// Check for "Administrator" permission
		if role, ok := roleMap[roleID]; ok && role.Permissions&0x8 > 0 {
			return true, nil
		}
	}
	return false, nil
}
