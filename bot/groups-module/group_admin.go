package groups

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (m Module) groupCreate(session *discordgo.Session, event *discordgo.MessageCreate, name string) error {
	if isGroupManager, err := m.isGroupManager(session, event); err != nil {
		return err
	} else if !isGroupManager {
		session.ChannelMessageSend(event.ChannelID, "You are not allowed to create groups. Administrator permissions required.")
		return nil
	}

	newRoleName := fmt.Sprintf("%s%s", m.Prefix, strings.ToLower(name))

	roles, err := session.GuildRoles(event.GuildID)
	if err != nil {
		return err
	}
	for _, role := range roles {
		if newRoleName == role.Name {
			msg := fmt.Sprintf("Role with name `%s` already exists.", role.Name)
			_, err := session.ChannelMessageSend(event.ChannelID, msg)
			return err
		}
	}

	m.Log.Debugf("Creating role; guildID=%s", event.GuildID)

	role, err := session.GuildRoleCreate(event.GuildID)
	if err != nil {
		m.Log.Debugf("%v", err)
		return err
	}

	m.Log.Debugf("Editing role; guildID=%s, roleID=%s, roleName=%s", event.GuildID, role.ID, newRoleName)
	roleEdited, err := session.GuildRoleEdit(event.GuildID, role.ID, newRoleName, 0, false, 0, true)
	if err != nil {
		if err2 := session.GuildRoleDelete(event.GuildID, role.ID); err2 != nil {
			return fmt.Errorf("%v; %v", err, err2)
		}
		return err
	}
	msg := fmt.Sprintf("Created role `%s` representing the group `%s`", roleEdited.Name, name)
	_, err = session.ChannelMessageSend(event.ChannelID, msg)
	return err
}

func (m Module) groupDelete(session *discordgo.Session, event *discordgo.MessageCreate, name string) error {
	if isGroupManager, err := m.isGroupManager(session, event); err != nil {
		return err
	} else if !isGroupManager {
		_, err := session.ChannelMessageSend(event.ChannelID, "You are not allowed to delete groups. Administrator permissions required.")
		return err
	}

	roleName := fmt.Sprintf("%s%s", m.Prefix, strings.ToLower(name))

	roles, err := session.GuildRoles(event.GuildID)
	if err != nil {
		return err
	}

	foundID := ""
	for _, role := range roles {
		if roleName == role.Name {
			foundID = role.ID
			break
		}
	}
	if foundID == "" {
		_, err := session.ChannelMessageSend(event.ChannelID, "Could not find a group by that name.")
		return err
	}

	if err := session.GuildRoleDelete(event.GuildID, foundID); err != nil {
		return err
	}

	msg := fmt.Sprintf("Deleted role `%s` representing the group `%s`", roleName, name)
	_, err = session.ChannelMessageSend(event.ChannelID, msg)
	return err
}
