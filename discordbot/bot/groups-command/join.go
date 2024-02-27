package groupscommand

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"

	guildcache "github.com/fsufitch/tagioalisi-bot/bot/guild-cache"
	"github.com/fsufitch/tagioalisi-bot/bot/util/interactions"
)

var cmdJoinGroup = &discordgo.ApplicationCommandOption{
	Name:        "join",
	Description: "join a group",
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:         "group",
			Description:  "group to join",
			Type:         discordgo.ApplicationCommandOptionString,
			Required:     true,
			Autocomplete: true,
		},
	},
}

func (cmd *GroupsCommandModule) subcommandJoin(iw *interactions.InteractionWrapper) {
	cmd.Logger.Debugf("..... %+v\n", iw.InteractionCreate().Data)
	cmd.Logger.Debugf("-> subcommandJoin ?")
	if iw.GetCommandOption("join") == nil {
		// This isn't a join command
		return
	}
	cmd.Logger.Debugf("-> subcommandJoin !")
	groupName := ""
	groupNameOpt := iw.GetCommandOption("join", "group")
	if groupNameOpt != nil {
		groupName = groupNameOpt.StringValue()
	}
	if groupName == "" {
		panic("group name is empty")
	}

	guildMember := iw.Interaction().Member
	if guildMember == nil {
		iw.RespondError(interactions.InteractionError{
			Title: "Error: interest groups are only supported in a guild context",
		})
	}

	guildID := iw.Interaction().GuildID
	roleName := cmd.Prefixer.AddGroupRolePrefix(groupName)

	cmd.Logger.Debugf("have cache manager: %+v", cmd.GuildCacheManager)
	role, err := cmd.GuildCacheManager.Cache(guildID).Role(iw.Session(), guildcache.HasName(roleName))
	if errors.Is(err, guildcache.ErrRoleNotFound) {
		iw.RespondError(interactions.InteractionError{
			Title:       "Failed joining group",
			Description: fmt.Sprintf("No such group: \"%s\" (role=\"%s\")", groupName, roleName),
		})
		return
	} else if err != nil {
		panic(err)
	}

	if slices.Contains(guildMember.Roles, role.ID) {
		iw.RespondError(interactions.InteractionError{
			Title:       "Failed joining group",
			Description: fmt.Sprintf("%s is already a member of \"%s\"", guildMember.Mention(), groupName),
		})
		return
	}

	err = iw.Session().GuildMemberRoleAdd(guildID, guildMember.User.ID, role.ID)
	if err != nil {
		panic(err)
	}

	iw.RespondEmbed(&discordgo.MessageEmbed{
		Title:       "Group joined!",
		Description: fmt.Sprintf("%s joined \"%s\"", guildMember.Mention(), groupName),
	})
}

func (cmd *GroupsCommandModule) subcommandJoinAutocomplete(iw *interactions.InteractionWrapper) []string {
	if iw.GetCommandOption("join") == nil {
		// This isn't a join command
		return nil
	}

	groupNameOpt := iw.GetCommandOption("join", "group")
	if groupNameOpt == nil {
		cmd.Logger.Warningf("no group option found; this should not happen")
		return []string{}
	}
	partialGroupName := groupNameOpt.StringValue()

	roleList, err := cmd.GuildCacheManager.Cache(iw.Interaction().GuildID).Roles(iw.Session(), guildcache.AnyRole())
	if err != nil {
		panic(fmt.Errorf("failed to fetch role list: %v", err))
	}

	groupChoices := []string{}
	for _, role := range roleList {
		groupName, err := cmd.Prefixer.RemoveGroupRolePrefix(role.Name)
		if err != nil {
			// Fine, skip over it
			continue
		}
		if strings.HasPrefix(strings.ToLower(groupName), strings.ToLower(partialGroupName)) {
			groupChoices = append(groupChoices, groupName)
		}
	}
	return groupChoices
}
