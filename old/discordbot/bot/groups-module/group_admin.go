package groups

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func (m Module) groupCreate(session *discordgo.Session, event *discordgo.MessageCreate, name string) error {
	m.Log.Debugf("groups: message %s creating group `%s`", event.ID, name)
	if isGroupManager, err := m.isGroupManager(session, event); err != nil {
		return errors.Wrap(err, "could not check group manager permissions")
	} else if !isGroupManager {
		session.ChannelMessageSend(event.ChannelID, "You are not allowed to create groups. Administrator permissions required.")
		return nil
	}

	newRoleName := fmt.Sprintf("%s%s", m.Prefix, strings.ToLower(name))

	roles, err := session.GuildRoles(event.GuildID)
	if err != nil {
		return errors.Wrap(err, "could not retrieve existing roles")
	}
	for _, role := range roles {
		if newRoleName == role.Name {
			msg := fmt.Sprintf("Role with name `%s` already exists.", role.Name)
			_, err := session.ChannelMessageSend(event.ChannelID, msg)
			return errors.Wrap(err, "could not send message")
		}
	}

	m.Log.Debugf("groups: creating role; guildID=%s", event.GuildID)

	role, err := session.GuildRoleCreate(event.GuildID, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create role")
	}

	m.Log.Debugf("groups: editing role; guildID=%s, roleID=%s, roleName=%s", event.GuildID, role.ID, newRoleName)
	roleEdited, err := session.GuildRoleEdit(event.GuildID, role.ID, &discordgo.RoleParams{
		Name: newRoleName,
	})
	if err != nil {
		if err2 := session.GuildRoleDelete(event.GuildID, role.ID); err2 != nil {
			return errors.Wrapf(err, "could not edit role; could not revert role creation: %v", err2)
		}
		return errors.Wrap(err, "could not edit role")
	}
	msg := fmt.Sprintf("Created role `%s` representing the group `%s`", roleEdited.Name, name)
	_, err = session.ChannelMessageSend(event.ChannelID, msg)
	return errors.Wrap(err, "could not send message")
}

func (m Module) groupDelete(session *discordgo.Session, event *discordgo.MessageCreate, name string) error {
	m.Log.Debugf("groups: message %s deleting group `%s`", event.ID, name)
	if isGroupManager, err := m.isGroupManager(session, event); err != nil {
		return errors.Wrap(err, "could not check group manager permissions")
	} else if !isGroupManager {
		session.ChannelMessageSend(event.ChannelID, "You are not allowed to delete groups. Administrator permissions required.")
		return nil
	}

	roleName := fmt.Sprintf("%s%s", m.Prefix, strings.ToLower(name))

	roles, err := session.GuildRoles(event.GuildID)
	if err != nil {
		return errors.Wrap(err, "could not retrieve existing roles")
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
		return errors.Wrap(err, "could not send message")
	}

	if err := session.GuildRoleDelete(event.GuildID, foundID); err != nil {
		return errors.Wrap(err, "could not delete role")
	}

	msg := fmt.Sprintf("Deleted role `%s` representing the group `%s`", roleName, name)
	_, err = session.ChannelMessageSend(event.ChannelID, msg)
	return errors.Wrap(err, "could not send message")
}
