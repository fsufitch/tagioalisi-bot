package interactions

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/fsufitch/tagioalisi-bot/log"
)

type InteractionWrapper struct {
	command           string
	logger            *log.Logger
	session           *discordgo.Session
	interactionCreate *discordgo.InteractionCreate
	chAcknowledged    chan struct{}
}

func WrapInteraction(s *discordgo.Session, ic *discordgo.InteractionCreate, logger *log.Logger) (*InteractionWrapper, error) {
	if s == nil {
		return nil, errors.New("cannot wrap a nil session")
	}
	if ic == nil {
		return nil, errors.New("cannot wrap a nil interaction")
	}
	if logger == nil {
		return nil, errors.New("cannot wrap a nil logger")
	}

	wrapper := InteractionWrapper{
		logger:            logger,
		session:           s,
		interactionCreate: ic,
		chAcknowledged:    make(chan struct{}),
	}

	return &wrapper, nil
}

func (iw InteractionWrapper) Command() string {
	return iw.command
}

func (iw InteractionWrapper) InteractionCreate() *discordgo.InteractionCreate {
	return iw.interactionCreate
}

func (iw InteractionWrapper) Session() *discordgo.Session {
	return iw.session
}

func (iw InteractionWrapper) Acknowledged() bool {
	select {
	case <-iw.chAcknowledged:
		return true
	default:
		return false
	}
}

func (iw InteractionWrapper) Acknowledge() {
	if iw.Acknowledged() {
		panic("interaction was already acknowledged")
	}
	close(iw.chAcknowledged)
}

func (iw InteractionWrapper) UserID() string {
	if iw.interactionCreate.User != nil {
		return iw.interactionCreate.User.ID
	}
	if iw.interactionCreate.Member != nil && iw.interactionCreate.Member.User != nil {
		return iw.interactionCreate.Member.User.ID
	}
	panic("could not find a User or Member for the interaction")
}

func (iw InteractionWrapper) InteractionGuild() *discordgo.Guild {
	guild, err := iw.session.Guild(iw.interactionCreate.GuildID)
	if err != nil {
		panic("failed to get guild of interaction")
	}
	return guild
}

func (iw InteractionWrapper) Interaction() *discordgo.Interaction {
	return iw.interactionCreate.Interaction
}

func (iw InteractionWrapper) RespondError(err InteractionError) {
	iw.Acknowledge()
	interactionError := iw.session.InteractionRespond(
		iw.Interaction(),
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Error: " + err.Title,
						Color:       0xff0000,
						Description: err.Description,
					},
				},
			},
		})
	if interactionError != nil {
		panic(fmt.Sprintf("error response failed: %s", err.Error()))
	}
}

func (iw InteractionWrapper) Response() *discordgo.Message {
	msg, err := iw.session.InteractionResponse(iw.Interaction())
	if err != nil {
		iw.logger.Errorf("error getting interaction response; is it supposed to be there? :: %v", err)
	}
	return msg
}

func (iw InteractionWrapper) RespondEmbed(embeds ...*discordgo.MessageEmbed) {
	iw.Acknowledge()
	err := iw.session.InteractionRespond(iw.Interaction(), &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embeds,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to respond with embeds: %s", err.Error()))
	}
}

func (iw InteractionWrapper) RespondDeferred() {
	iw.Acknowledge()
	err := iw.session.InteractionRespond(iw.Interaction(), &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to respond with deferral: %s", err.Error()))
	}
}

func (iw InteractionWrapper) UpdateDeferredEmbeds(embeds ...*discordgo.MessageEmbed) {
	if !iw.Acknowledged() {
		panic("cannot update a deferred response if it was not acknowledged already")
	}
	_, err := iw.session.InteractionResponseEdit(iw.Interaction(), &discordgo.WebhookEdit{
		Embeds: &embeds,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to update deferred response: %s", err.Error()))
	}
}

func (iw InteractionWrapper) RespondAutocomplete(choices ...string) {
	iw.Acknowledge()
	choicesStructs := []*discordgo.ApplicationCommandOptionChoice{}
	for _, choiceStr := range choices {
		choicesStructs = append(choicesStructs, &discordgo.ApplicationCommandOptionChoice{
			Name:  choiceStr,
			Value: choiceStr,
		})
	}
	err := iw.session.InteractionRespond(iw.Interaction(), &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: choicesStructs,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to respond to autocorrect: %s", err.Error()))
	}
}

func (iw *InteractionWrapper) GetCommandOption(lookup ...string) *discordgo.ApplicationCommandInteractionDataOption {
	inter := iw.Interaction()
	if inter == nil ||
		!(inter.Type == discordgo.InteractionApplicationCommand ||
			inter.Type == discordgo.InteractionApplicationCommandAutocomplete) {
		// This is not an interaction-command that has option data
		iw.logger.Debugf("tried to look up a non-appcommand interaction's options")
		return nil
	}

	if len(lookup) == 0 {
		panic("extracting command options requires actual options specified")
	}

	lookupQueue := append([]string{}, lookup...)

	interData := inter.ApplicationCommandData()

	var currentOption *discordgo.ApplicationCommandInteractionDataOption
	var nextOptions = interData.Options
	for len(lookupQueue) > 0 {
		// var nextOptions []*discordgo.ApplicationCommandInteractionDataOption
		var nextOptName string
		nextOptName, lookupQueue = lookupQueue[0], lookupQueue[1:]
		currentOption = nil
		for _, opt := range nextOptions {
			iw.logger.Debugf("    search option '%s': '%s'", nextOptName, opt.Name)
			if opt.Name != nextOptName {
				continue
			}
			currentOption = opt
			break
		}
		if currentOption == nil {
			break
		}
		nextOptions = currentOption.Options
	}

	if currentOption == nil {
		iw.logger.Debugf("looked up %+v, found nothing")
	} else {
		iw.logger.Debugf("looked up %+v, found: name=%s type=%s value='%s'", lookup, currentOption.Name, currentOption.Type.String(), currentOption.Value)

	}

	return currentOption
}
