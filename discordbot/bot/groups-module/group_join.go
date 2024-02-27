package groups

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/fsufitch/tagioalisi-bot/bot/util"
)

var cmdGroupMember_Join = &discordgo.ApplicationCommandOption{
	Name:        "join",
	Description: "join a group",
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "group",
			Description: "group to join",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
			// Autocomplete: true,
		},
	},
}

func (m *Module) handleGroupJoin(s *discordgo.Session, inter *discordgo.InteractionCreate) {
	// s.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
	// 	Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	// 	Data: &discordgo.InteractionResponseData{
	// 		Content: "loading",
	// 	},
	// })
	fmt.Printf("AAAAA inter=%s\n", inter.Interaction.ID)
	fmt.Println(inter.Interaction.ID + " " + inter.Type.String())
	if !(inter.Type == discordgo.InteractionApplicationCommand) {
		return // Looking for an app command, not autocomplete etc
	}

	joinGroupOpt := util.FindInteractionOption(inter.Interaction, "groups", "join", "group")
	if joinGroupOpt == nil {
		return // No command match
	}

	fmt.Printf("11111 inter=%s\n", inter.Interaction.ID)

	joinGroupName := strings.ToLower(joinGroupOpt.StringValue())

	m.Log.Debugf("user %+v wants to join group %+v", inter.Member.User.Username, joinGroupName)

	role, err := getRoleByName(s, inter.GuildID, groupToRole(joinGroupName))
	fmt.Printf("22222 inter=%s\n", inter.Interaction.ID)
	if errors.Is(err, errNoSuchGroup) {
		m.InterUtil.ErrorResponse(s, inter.Interaction, "Invalid group name", fmt.Sprintf("No such group: %s", joinGroupName))
		return
	}
	if err != nil {
		m.InterUtil.UnexpectedError(s, inter.Interaction, err)
		return
	}

	fmt.Printf("BBBBB inter=%s\n", inter.Interaction.ID)

	isIndbg, err := userIsInGroup(s, inter.GuildID, inter.Member.User.ID, joinGroupName)
	fmt.Printf("isin=%+v err=%+v\n", isIndbg, err)

	if isIn, _ := userIsInGroup(s, inter.GuildID, inter.Member.User.ID, joinGroupName); isIn {
		m.InterUtil.ErrorResponse(s, inter.Interaction, "Already in group", fmt.Sprintf("%s is already in `%s`", inter.Member.Mention(), joinGroupName))
		return
	}
	fmt.Printf("CCCCC inter=%s\n", inter.Interaction.ID)

	err = s.GuildMemberRoleAdd(inter.GuildID, inter.Member.User.ID, role.ID)
	if err != nil {
		m.InterUtil.UnexpectedError(s, inter.Interaction, err)
		return
	}

	fmt.Printf("DDDDD inter=%s\n", inter.Interaction.ID)

	err = s.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%s joined `%s`", inter.Member.Mention(), joinGroupName),
		},
	})
	fmt.Printf("EEEEE inter=%s\n", inter.Interaction.ID)

	if err != nil {
		m.InterUtil.UnexpectedError(s, inter.Interaction, err)
		return
	}
}

func (m *Module) handleGroupJoin_AutoComplete(s *discordgo.Session, inter *discordgo.InteractionCreate) {
	if !(inter.Type == discordgo.InteractionApplicationCommandAutocomplete) {
		return // Looking for an autocomplete command
	}

	joinGroupOpt := util.FindInteractionOption(inter.Interaction, "groups", "join", "group")
	if joinGroupOpt == nil {
		return // No command match
	}

	userID, err := util.InteractionUserID(inter.Interaction)
	if err != nil {
		m.InterUtil.UnexpectedError(s, inter.Interaction, err)
		return
	}

	partialGroupName := strings.ToLower(joinGroupOpt.StringValue())
	validGroups, err := getGuildGroups(s, inter.GuildID)
	if err != nil {
		m.InterUtil.UnexpectedError(s, inter.Interaction, err)
		return
	}

	acChoices := []string{}
	for _, group := range validGroups {
		if !strings.HasPrefix(strings.ToLower(group), partialGroupName) {
			continue
		}
		if userInGroup, _ := userIsInGroup(s, inter.GuildID, userID, group); userInGroup {
			continue
		}
		acChoices = append(acChoices, group)
	}

	s.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: util.BuildChoices(acChoices...),
		},
	})
}
