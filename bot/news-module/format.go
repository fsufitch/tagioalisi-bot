package news

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagioalisi-bot/azure"
)

// formatter is a common interface for formatting news results
type formatter func(session *discordgo.Session, channelID string, results azure.NewsResults) error

func compactFormatter(session *discordgo.Session, channelID string, results azure.NewsResults) error {
	embed := &discordgo.MessageEmbed{
		Color: 0xbeefed,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Want full embeds? Try `!news -v`. Want more/less news? Try `!news -n <number>`.",
		},
		Fields: []*discordgo.MessageEmbedField{},
	}

	for i := 0; i < results.Len(); i++ {
		article, _ := results.Get(i)
		fmt.Printf("%+v\n", article)
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("__%s__", article.Source()),
			Value: fmt.Sprintf("[%s](%s)", article.Title(), article.URL()),
		})
	}
	_, err := session.ChannelMessageSendEmbed(channelID, embed)
	return err
}

func verboseFormatter(session *discordgo.Session, channelID string, results azure.NewsResults) error {
	for i := 0; i < results.Len(); i++ {
		article, _ := results.Get(i)
		embed := &discordgo.MessageEmbed{
			Color: 0xbeefed,
			Title: article.Title(),
			URL:   article.URL(),
			Author: &discordgo.MessageEmbedAuthor{
				Name: article.Source(),
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: article.ThumbnailURL(),
			},
			Description: article.Description(),
		}
		_, err := session.ChannelMessageSendEmbed(channelID, embed)
		if err != nil {
			return err
		}
	}
	return nil
}
