package groups

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagioalisi-bot/bot/util"
	"github.com/pkg/errors"
)

func (m Module) groupList(session *discordgo.Session, event *discordgo.MessageCreate) error {
	m.Log.Debugf("groups: message %v listing groups", event.ID)
	roles, err := session.GuildRoles(event.GuildID)
	if err != nil {
		return errors.Wrap(err, "could not retrieve existing roles")
	}

	managedGroupRoles := []*discordgo.Role{}
	for _, role := range roles {
		if strings.HasPrefix(role.Name, string(m.Prefix)) {
			managedGroupRoles = append(managedGroupRoles, role)
		}
	}

	buf := bytes.NewBufferString("")
	buf.WriteString("Your requested list of groups:\n")
	for _, role := range managedGroupRoles {
		groupName := role.Name[len(m.Prefix):]
		buf.WriteString(fmt.Sprintf(" - %s (role: %s)\n", groupName, role.Name))
	}

	ch, err := session.UserChannelCreate(event.Author.ID)
	if err != nil {
		return errors.Wrap(err, "could not open private channel to user")
	}

	return util.DiscordMessageSendRawBlock(session, ch.ID, buf.String())
}

func (m Module) groupJoin(session *discordgo.Session, event *discordgo.MessageCreate, groupName string) error {
	m.Log.Debugf("groups: message %v joining group", event.ID)
	roles, err := session.GuildRoles(event.GuildID)
	if err != nil {
		return errors.Wrap(err, "could not retrieve existing roles")
	}

	roleName := fmt.Sprintf("%s%s", m.Prefix, strings.ToLower(groupName))

	var foundRole *discordgo.Role
	for _, role := range roles {
		if role.Name == roleName {
			foundRole = role
		}
	}
	if foundRole == nil {
		_, err := session.ChannelMessageSend(event.ChannelID, "Could not find a group by that name.")
		return errors.Wrap(err, "could not send message")
	}

	targetUserIDs := []string{}
	targetOthers := false

	for _, mention := range event.Mentions {
		targetUserIDs = append(targetUserIDs, mention.ID)
		if mention.ID != event.Author.ID {
			targetOthers = true
		}
	}
	if len(targetUserIDs) == 0 {
		targetUserIDs = []string{event.Author.ID}
	}

	if targetOthers {
		if isGroupManager, err := m.isGroupManager(session, event); err != nil {
			return errors.Wrap(err, "could not check group manager permissions")
		} else if !isGroupManager {
			_, err := session.ChannelMessageSend(event.ChannelID, "You are not allowed to add others to groups. Administrator permissions required.")
			return errors.Wrap(err, "could not send message")
		}
	}

	for _, target := range targetUserIDs {
		if err := session.GuildMemberRoleAdd(event.GuildID, target, foundRole.ID); err != nil {
			return errors.Wrapf(err, "could not add member %v to role", target)
		}
	}

	msg := fmt.Sprintf("Added %d user(s) to the group", len(targetUserIDs))
	_, err = session.ChannelMessageSend(event.ChannelID, msg)
	return errors.Wrap(err, "could not send message")
}

func (m Module) groupLeave(session *discordgo.Session, event *discordgo.MessageCreate, groupName string) error {
	m.Log.Debugf("groups: message %v leaving group", event.ID)
	roles, err := session.GuildRoles(event.GuildID)
	if err != nil {
		return errors.Wrap(err, "could not retrieve existing roles")
	}

	roleName := fmt.Sprintf("%s%s", m.Prefix, strings.ToLower(groupName))

	var foundRole *discordgo.Role
	for _, role := range roles {
		if role.Name == roleName {
			foundRole = role
		}
	}
	if foundRole == nil {
		_, err := session.ChannelMessageSend(event.ChannelID, "Could not find a group by that name.")
		return errors.Wrap(err, "could not send message")
	}

	targetUserIDs := []string{}
	targetOthers := false

	for _, mention := range event.Mentions {
		targetUserIDs = append(targetUserIDs, mention.ID)
		if mention.ID != event.Author.ID {
			targetOthers = true
		}
	}
	if len(targetUserIDs) == 0 {
		targetUserIDs = []string{event.Author.ID}
	}

	if targetOthers {
		if isGroupManager, err := m.isGroupManager(session, event); err != nil {
			return errors.Wrap(err, "could not check group manager permissions")
		} else if !isGroupManager {
			_, err := session.ChannelMessageSend(event.ChannelID, "You are not allowed to remove others from groups. Administrator permissions required.")
			return errors.Wrap(err, "could not send message")
		}
	}

	for _, target := range targetUserIDs {
		if err := session.GuildMemberRoleRemove(event.GuildID, target, foundRole.ID); err != nil {
			return errors.Wrapf(err, "could not remove member %v from role", target)
		}
	}

	msg := fmt.Sprintf("Removed %d user(s) from the group", len(targetUserIDs))
	_, err = session.ChannelMessageSend(event.ChannelID, msg)
	return errors.Wrap(err, "could not send message")
}
