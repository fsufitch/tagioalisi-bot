package dice

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (m *Module) RegisterApplicationCommand(ctx context.Context, session *discordgo.Session) error {
	cmd := discordgo.ApplicationCommand{
		ApplicationID: string(m.AppID),
		Name:          "roll-dice",
		Type:          discordgo.ChatApplicationCommand,
		Description:   "Roll some dice!",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "dice-expression",
				Description: "What dice should be rolled, or calculations done? e.g. 5d8 + 3d6 - 10",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        "verbose",
				Description: "Be more verbose in the output",
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Required:    false,
			},
		},
	}

	_, err := session.ApplicationCommandCreate(string(m.AppID), "327526752203177984", &cmd)
	if err != nil {
		return err
	}
	cancel := session.AddHandler(m.handleApplicationCommand)
	go func() {
		<-ctx.Done()
		m.Log.Infof("dice application command context done")
		cancel()
	}()
	m.Log.Infof("Registered dice application command")
	return nil
}

func (m *Module) handleApplicationCommand(s *discordgo.Session, event *discordgo.InteractionCreate) {
	if event.ApplicationCommandData().Name != "roll-dice" {
		return
	}

	diceExpression := ""
	for _, opt := range event.ApplicationCommandData().Options {
		if opt.Name == "dice-expression" {
			diceExpression = opt.StringValue()
		}
	}
	diceExpression = strings.TrimSpace(diceExpression)

	if diceExpression == "" {
		s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: ":face_with_raised_eyebrow: I couldn't find a dice expression in your command! ",
			},
		})
	}

	result, err := m.Calculator.Calculate(diceExpression)
	if err != nil {
		s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf(":warning: Calculator error! %v", err),
			},
		})
		return
	}

	responseEmbed := discordgo.MessageEmbed{
		Title:       fmt.Sprintf(":game_die: Your result: %d", result.Value),
		Description: fmt.Sprintf("You rolled: %s", diceExpression),
	}

	for _, roll := range result.Rolls {
		total := 0
		shownResults := []string{}
		for i, value := range roll.Results {
			if i < 30 {
				shownResults = append(shownResults, fmt.Sprintf("%d", value))
			}
			total += value
		}

		resultsText := strings.Join(shownResults, " · ")

		if len(roll.Results) > 30 {
			resultsText += fmt.Sprintf(" ··· and %d more", len(roll.Results)-30)
		}

		responseEmbed.Fields = append(responseEmbed.Fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%dd%d = %d", roll.Count, roll.Sides, total),
			Inline: true,
			Value:  resultsText,
		})
	}

	s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{&responseEmbed},
		},
	},
	)
}
