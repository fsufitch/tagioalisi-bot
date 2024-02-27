package util

import (
	"errors"

	"github.com/bwmarrin/discordgo"

	"github.com/fsufitch/tagioalisi-bot/log"
)

type InteractionUtil struct {
	Log *log.Logger
}

func (iu InteractionUtil) ErrorResponse(s *discordgo.Session, inter *discordgo.Interaction, title string, details string) {
	err := s.InteractionRespond(inter, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Error: " + title,
					Color:       0xff0000,
					Description: details,
				},
			},
		},
	})
	if err != nil {
		iu.Log.Errorf("error response failed; title=%+v -- %s", title, err.Error())
	}
}

func (iu InteractionUtil) UnexpectedError(s *discordgo.Session, inter *discordgo.Interaction, err error) {
	iu.ErrorResponse(s, inter, "Unexpected Error", err.Error())
}

// RespondWithError is a one-liner shortcut for reporting errors for interactions
//
//	If `err` is nil, it returns false and does nothing.
//	Otherwise, it responds to the interaction with an error, and returns true.
func (iu InteractionUtil) RespondWithError(s *discordgo.Session, inter *discordgo.Interaction, err error) bool {
	if err == nil {
		return false
	}
	iu.UnexpectedError(s, inter, err)
	return true
}

func BuildChoices(choiceStrings ...string) []*discordgo.ApplicationCommandOptionChoice {
	choices := []*discordgo.ApplicationCommandOptionChoice{}
	for _, choiceString := range choiceStrings {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  choiceString,
			Value: choiceString,
		})
	}
	return choices
}

func FindInteractionOption(inter *discordgo.Interaction, command string, optPath ...string) *discordgo.ApplicationCommandInteractionDataOption {
	if len(optPath) < 1 {
		panic("tried to FindInteractionOption without any path")
	}

	if !(inter.Type == discordgo.InteractionApplicationCommand || inter.Type == discordgo.InteractionApplicationCommandAutocomplete) {
		// Not an app command, no match
		return nil
	}

	cmdData := inter.ApplicationCommandData()
	if cmdData.Name != command {
		// Wrong command, no match
		return nil
	}

	var currentOption *discordgo.ApplicationCommandInteractionDataOption
	nextOptionList := cmdData.Options

	for len(optPath) > 0 {
		currentOption = nil
		for _, opt := range nextOptionList {
			if opt.Name == optPath[0] {
				currentOption = opt
			}
		}
		if currentOption == nil {
			// Failed to drill down, no match
			return nil
		}
		if len(optPath) == 1 {
			// Found it!
			return currentOption
		}
		// Continue drilling down
		nextOptionList = currentOption.Options
		optPath = optPath[1:]
	}

	// Might have a halfway match here
	return currentOption
}

func InteractionUserID(inter *discordgo.Interaction) (string, error) {
	if inter.User != nil {
		return inter.User.ID, nil
	}
	if inter.Member != nil {
		return inter.Member.User.ID, nil
	}
	return "", errors.New("no user found in interaction")
}
