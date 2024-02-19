package groups

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

// var permissionManageGroups int64 = discordgo.PermissionManageRoles

// var cmdGroupsAdmin = discordgo.ApplicationCommand{
// 	Name:                     "groups-admin",
// 	Description:              "administer group roles for the current server",
// 	DefaultMemberPermissions: &permissionManageGroups,
// }

func (m *Module) applicationCommandHandlers() []any {
	return []any{
		m.handleGroupJoin_AutoComplete,
		m.handleGroupJoin,
	}
}

func (m *Module) RegisterApplicationCommand(ctx context.Context, session *discordgo.Session, guildID string) (err error) {
	_, err = session.ApplicationCommandCreate(string(m.ApplicationID), guildID, cmdGroupMember)
	if err != nil {
		return
	}

	// _, err = session.ApplicationCommandCreate(string(m.ApplicationID), guildID, &cmdGroupsAdmin)
	// if err != nil {
	// 	return err
	// }

	for _, handler := range m.applicationCommandHandlers() {
		cancel := session.AddHandler(handler)
		go func() {
			<-ctx.Done()
			cancel()
		}()
	}

	return nil
}
