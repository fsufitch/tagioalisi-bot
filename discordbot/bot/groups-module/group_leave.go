package groups

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/fsufitch/tagioalisi-bot/bot/util"
)

var cmdGroupMember_Leave = &discordgo.ApplicationCommandOption{
	Name:        "leave",
	Description: "leave a group",
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:         "group",
			Description:  "group to leave",
			Type:         discordgo.ApplicationCommandOptionString,
			Required:     true,
			Autocomplete: true,
		},
	},
}

func (m *Module) handleGroupLeave(s *discordgo.Session, inter *discordgo.InteractionCreate) {
	if !(inter.Type == discordgo.InteractionApplicationCommand) {
		return // Looking for an app command, not autocomplete etc
	}

	leaveGroupOpt := util.FindInteractionOption(inter.Interaction, "groups", "leave", "group")
	if leaveGroupOpt == nil {
		return // No command match
	}

	leaveGroupName := strings.ToLower(leaveGroupOpt.StringValue())

	userID, err := util.InteractionUserID(inter.Interaction)
	if err != nil {
		m.InterUtil.UnexpectedError(s, inter.Interaction, err)
		return
	}

	m.Log.Debugf("user %+v wants to leave group %+v", inter.Member.User.Username, leaveGroupName)

	role, err := getRoleByName(s, inter.GuildID, groupToRole(leaveGroupName))
	if errors.Is(err, errNoSuchGroup) {
		m.InterUtil.ErrorResponse(s, inter.Interaction, "No such group", fmt.Sprintf("Group `%s` does not exist", leaveGroupName))
		return
	}

	if isIn, _ := userIsInGroup(s, inter.GuildID, userID, leaveGroupName); !isIn {
		m.InterUtil.ErrorResponse(s, inter.Interaction, "Already not in group", fmt.Sprintf("%s is already not in `%s`", inter.Member.Mention(), leaveGroupName))
		return
	}

	if err = s.GuildMemberRoleRemove(inter.GuildID, inter.Member.User.ID, role.ID); err != nil {
		m.InterUtil.UnexpectedError(s, inter.Interaction, err)
		return
	}

	s.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%s left group `%s`", inter.Member.Mention(), leaveGroupName),
		},
	})
}

func (m *Module) handleGroupLeave_AutoComplete(s *discordgo.Session, inter *discordgo.InteractionCreate) {
	if !(inter.Type == discordgo.InteractionApplicationCommandAutocomplete) {
		return // Looking for an autocomplete command
	}

	leaveGroupOpt := util.FindInteractionOption(inter.Interaction, "groups", "leave", "group")
	if leaveGroupOpt == nil {
		return // No command match
	}

	userID, err := util.InteractionUserID(inter.Interaction)
	if err != nil {
		m.InterUtil.UnexpectedError(s, inter.Interaction, err)
		return
	}

	userGroups, err := getMemberGroups(s, inter.GuildID, userID)
	if err != nil {
		m.InterUtil.UnexpectedError(s, inter.Interaction, err)
	}

	partialGroupName := strings.ToLower(leaveGroupOpt.StringValue())

	choices := []string{}
	for _, group := range userGroups {
		if !strings.HasPrefix(strings.ToLower(group), partialGroupName) {
			continue
		}
		choices = append(choices, group)
	}

	s.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: util.BuildChoices(choices...),
		},
	})
}
