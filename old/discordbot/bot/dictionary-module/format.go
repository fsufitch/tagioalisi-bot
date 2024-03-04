package dictionary

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagioalisi-bot/merriam-webster/types"
)

func collegiateResultFormatter(session *discordgo.Session, channelID string, word string, results []types.CollegiateResult) error {
	embed := &discordgo.MessageEmbed{
		Color: 0x305f7a,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Brought to you by Merriam-Webster",
		},
		Fields: []*discordgo.MessageEmbedField{},
		Title:  fmt.Sprintf(`Short definition for "%s"`, word),
		URL:    fmt.Sprintf("https://www.merriam-webster.com/dictionary/%s", word),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://merriam-webster.com/assets/mw/static/app-css-images/logos/mw-logo.png",
		},
	}

	resultsWithShortDefinitions := []types.CollegiateResult{}
	for _, result := range results {
		if len(result.ShortDefinitions) > 0 {
			resultsWithShortDefinitions = append(resultsWithShortDefinitions, result)
		}
	}

	for i, result := range resultsWithShortDefinitions {
		field := &discordgo.MessageEmbedField{}

		nameParts := []string{
			fmt.Sprintf("%d. %s — %s", i+1, result.HeadwordInfo.Headword, result.Function),
		}
		for _, p := range result.HeadwordInfo.Pronounciations {
			nameParts = append(nameParts, p.MerriamWebsterFormat)
		}
		field.Name = strings.Join(nameParts, "; ")

		lines := []string{}

		for _, shortDef := range result.ShortDefinitions {
			lines = append(lines, shortDef)
		}

		if len(lines) > 1 {
			for i, line := range lines {
				lines[i] = fmt.Sprintf("%s. %s", string('a'+i), line)
			}
		}

		field.Value = strings.Join(lines, "\n")

		embed.Fields = append(embed.Fields, field)
	}

	_, err := session.ChannelMessageSendEmbed(channelID, embed)

	return err
}

func suggestionFormatter(session *discordgo.Session, channelID string, word string, suggestions []string) error {
	embed := &discordgo.MessageEmbed{
		Color:       0x305f7a,
		Title:       fmt.Sprintf(`"%s": no definition found, but maybe try these?`, word),
		Description: strings.Join(suggestions, ", "),
		URL:         fmt.Sprintf("https://www.merriam-webster.com/dictionary/%s", word),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://merriam-webster.com/assets/mw/static/app-css-images/logos/mw-logo.png",
		},
	}

	_, err := session.ChannelMessageSendEmbed(channelID, embed)
	return err
}

func errorFormatter(session *discordgo.Session, channelID string, word string, title string, text string) error {
	embed := &discordgo.MessageEmbed{
		Color:       0x305f7a,
		Title:       fmt.Sprintf(`"%s": %s`, word, title),
		Description: text,
		URL:         fmt.Sprintf("https://www.merriam-webster.com/dictionary/%s", word),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://merriam-webster.com/assets/mw/static/app-css-images/logos/mw-logo.png",
		},
	}

	_, err := session.ChannelMessageSendEmbed(channelID, embed)
	return err
}
