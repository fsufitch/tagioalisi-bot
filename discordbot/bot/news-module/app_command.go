package news

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var minResults float64 = 1
var maxResults float64 = 10

const (
	cmdNewsSearch  = "news-search"
	argSearchQuery = "search-query"
	argCompact     = "compact"
	argMaxResults  = "max-results"
)

func (m *Module) RegisterApplicationCommand(ctx context.Context, session *discordgo.Session, guildID string) error {
	cmd := discordgo.ApplicationCommand{
		ApplicationID: string(m.AppID),
		Name:          cmdNewsSearch,
		Type:          discordgo.ChatApplicationCommand,
		Description:   "Find some recent relevant news articles",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        argSearchQuery,
				Description: "Keywords to search for",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        argCompact,
				Description: "Show smaller embeds",
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Required:    false,
			},
			{
				Name:        argMaxResults,
				Description: "Maximum number of results to retrieve",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    false,
				MinValue:    &minResults,
				MaxValue:    maxResults,
			},
		},
	}

	_, err := session.ApplicationCommandCreate(string(m.AppID), guildID, &cmd)
	if err != nil {
		return err
	}
	cancel := session.AddHandler(m.handleApplicationCommand)
	go func() {
		<-ctx.Done()
		m.Log.Infof("news application command context done")
		cancel()
	}()
	m.Log.Infof("Registered news application command")
	return nil
}

func (m *Module) handleApplicationCommand(s *discordgo.Session, event *discordgo.InteractionCreate) {
	if event.ApplicationCommandData().Name != cmdNewsSearch {
		return
	}

	searchQuery := ""
	compact := false
	maxResults := 3
	for _, opt := range event.ApplicationCommandData().Options {
		switch opt.Name {
		case argSearchQuery:
			searchQuery = strings.TrimSpace(opt.StringValue())
		case argCompact:
			compact = opt.BoolValue()
		case argMaxResults:
			maxResults = int(opt.IntValue())
		}
	}

	if searchQuery == "" {
		s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: ":face_with_raised_eyebrow: I couldn't find any search terms in your command!",
			},
		})
		return
	}

	answer, err := m.News.Search(context.TODO(), searchQuery, int32(maxResults))
	if err != nil {
		s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Search failed; error: `%s`", err.Error()),
			},
		})
		return
	}

	var f formatter = verboseFormatter
	if compact {
		f = compactFormatter
	}
	embeds := f(answer)

	err = s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embeds,
		},
	})

	if err != nil {
		m.Log.Errorf("failed sending response: %+v", err)
	}
}
