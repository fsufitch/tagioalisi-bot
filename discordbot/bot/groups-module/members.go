package groups

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/fsufitch/tagioalisi-bot/bot/util"
)

var cmdGroupMember_Search = &discordgo.ApplicationCommandOption{
	Name:        "search",
	Description: "search for a group",
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "query",
			Description: "wtf",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}

var cmdGroupMember_Join = &discordgo.ApplicationCommandOption{
	Name:        "join",
	Description: "join a group",
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:         "group",
			Description:  "wtf",
			Type:         discordgo.ApplicationCommandOptionString,
			Required:     true,
			Autocomplete: true,
		},
	},
}

var cmdGroupMember_Leave = &discordgo.ApplicationCommandOption{
	Name:        "leave",
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Description: "leave a group",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:         "group",
			Description:  "wtf",
			Type:         discordgo.ApplicationCommandOptionString,
			Required:     true,
			Autocomplete: true,
		},
	},
}

var cmdGroupMember = &discordgo.ApplicationCommand{
	Name:        "groups",
	Description: "interact with group roles",
	Options: []*discordgo.ApplicationCommandOption{
		cmdGroupMember_Search,
		cmdGroupMember_Join,
		cmdGroupMember_Leave,
	},
	// Options: []*discordgo.ApplicationCommandOption{
	// 	{
	// 		Name:     "action",
	// 		Required: true,
	// 		Type:     discordgo.ApplicationCommandOptionSubCommand,
	// 		Options: []*discordgo.ApplicationCommandOption{
	// 			cmdGroupMember_Search,
	// 			cmdGroupMember_Join,
	// 			cmdGroupMember_Leave,
	// 		},
	// 	},
	// },
}

func (m *Module) handleGroupJoin(s *discordgo.Session, inter *discordgo.InteractionCreate) {
	if !(inter.Type == discordgo.InteractionApplicationCommand) {
		return // Looking for an app command, not autocomplete etc
	}

	joinGroupOpt := util.FindInteractionOption(inter.Interaction, "groups", "join", "group")
	if joinGroupOpt == nil {
		return // No command match
	}

	joinGroupName := strings.ToLower(joinGroupOpt.StringValue())

	m.Log.Debugf("user %+v wants to join group %+v", inter.Member.User.Username, joinGroupName)

	role, err := getRoleByName(s, inter.GuildID, groupToRole(joinGroupName))
	if errors.Is(err, errNoSuchGroup) {
		m.InterUtil.ErrorResponse(s, inter.Interaction, "Invalid group name", fmt.Sprintf("No such group: %s", joinGroupName))
		return
	}
	if err != nil {
		m.InterUtil.UnexpectedError(s, inter.Interaction, err)
		return
	}

	err = s.GuildMemberRoleAdd(inter.GuildID, inter.Member.User.ID, role.ID)
	if err != nil {
		m.InterUtil.UnexpectedError(s, inter.Interaction, err)
		return
	}

	s.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%s joined `%s`", inter.Member.Mention(), joinGroupName),
		},
	})
}

func (m *Module) handleGroupJoin_AutoComplete(s *discordgo.Session, inter *discordgo.InteractionCreate) {
	if !(inter.Type == discordgo.InteractionApplicationCommandAutocomplete) {
		return // Looking for an autocomplete command
	}

	joinGroupOpt := util.FindInteractionOption(inter.Interaction, "groups", "join", "group")
	if joinGroupOpt == nil {
		return // No command match
	}

	partialGroupName := strings.ToLower(joinGroupOpt.StringValue())
	validGroups, err := allGroups(s, inter.GuildID)
	if err != nil {
		m.InterUtil.UnexpectedError(s, inter.Interaction, err)
		return
	}

	acChoices := []string{}
	for _, group := range validGroups {
		lowerGroup := strings.ToLower(group)
		if strings.HasPrefix(lowerGroup, partialGroupName) {
			acChoices = append(acChoices, group)
		}
	}

	s.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: util.BuildChoices(acChoices...),
		},
	})
}
