package util

import (
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
		iu.Log.Errorf("error response failed; title=%+v : %s", title, err.Error())
	}
}

func (iu InteractionUtil) UnexpectedError(s *discordgo.Session, inter *discordgo.Interaction, err error) {
	iu.ErrorResponse(s, inter, "Unexpected Error", err.Error())
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

// func seekSuboption(curr *discordgo.ApplicationCommandInteractionDataOption, path ...string) *discordgo.ApplicationCommandInteractionDataOption {
// 	if len(path) < 1 {
// 		panic("this should be impossible")
// 	}
// 	if curr.Name != path[0] {
// 		return nil
// 	}
// 	if len(path) == 1 {
// 		return curr
// 	}
// 	for _, opt := range curr.Options {
// 		if found := seekSuboption(opt, path[1:]...); found != nil {
// 			return found
// 		}
// 	}
// 	return nil
// }

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
		if currentOption != nil && len(optPath) == 1 {
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
