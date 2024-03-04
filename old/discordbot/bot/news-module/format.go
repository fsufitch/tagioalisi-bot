package news

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/fsufitch/tagioalisi-bot/azure"
)

// formatter is a common interface for formatting news results
type formatter func(answer *azure.NewsAnswer) []*discordgo.MessageEmbed

func compactFormatter(answer *azure.NewsAnswer) []*discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Color: 0xbeefed,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Want full embeds? Try `!news -v`. Want more/less news? Try `!news -c <count>`.",
		},
		Fields: []*discordgo.MessageEmbedField{},
	}

	for _, article := range answer.Articles {
		// fmt.Printf("%+v\n", article)
		embedName := "(no source provided)"
		if len(article.Providers) > 0 {
			embedName = article.Providers[0].Name
		}
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("__%s__", embedName),
			Value: fmt.Sprintf("[%s](%s)", article.Name, article.URL),
		})
	}

	return []*discordgo.MessageEmbed{embed}
	
}

func verboseFormatter(answer *azure.NewsAnswer) (embeds []*discordgo.MessageEmbed) {
	for _, article := range answer.Articles {
		author := "unknown"
		if len(article.Providers) > 0 {
			author = article.Providers[0].Name
		}
		embed := &discordgo.MessageEmbed{
			Color: 0xbeefed,
			Title: article.Name,
			URL:   article.URL,
			Author: &discordgo.MessageEmbedAuthor{
				Name: author,
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: article.Image.Thumbnail.URL,
			},
			Description: article.Description,
		}
		embeds = append(embeds, embed)
	}
	return embeds
}
