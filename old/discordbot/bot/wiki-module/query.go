package wiki

import (
	"bytes"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagioalisi-bot/bot/util"
	"github.com/pkg/errors"
)

func (m *Module) showOptions(ctx commandContext) error {
	buf := bytes.NewBufferString("Options for -w:\n")

	for _, wiki := range m.WikiSupport.Wikis {
		languageSupport := "no"
		if wiki.DefaultLang != "" {
			languageSupport = fmt.Sprintf("yes (default: %s)", wiki.DefaultLang)
		}
		fmt.Fprintf(buf, "%s - %s; Language support: %s\n", wiki.ID, wiki.Name, languageSupport)
	}

	return util.DiscordMessageSendRawBlock(ctx.session, ctx.messageCreate.ChannelID, buf.String())
}

func (m *Module) queryWiki(ctx commandContext, wikiID string, lang string, page string) error {
	client, err := m.WikiSupport.Wikis[wikiID].Client(lang)
	if err != nil {
		return errors.Wrapf(err, "failed to initialize wiki client (id=%s, lang=%s)", wikiID, lang)
	}

	m.Log.Debugf("wiki: sending query id=%s, lang=%s, query=%s", wikiID, lang, page)
	result, err := client.Query(page)
	if err != nil {
		return errors.Wrapf(err, "unexpected error running query (id=%s, lang=%s, query=%s)", wikiID, lang, page)
	}
	m.Log.Debugf("wiki: got response %+v", result)

	if !result.Found {
		_, err = ctx.session.ChannelMessageSend(ctx.messageCreate.ChannelID, fmt.Sprintf(
			"Could not find a page on %s (%s) with the title: %s", m.WikiSupport.Wikis[wikiID].Name, lang, page,
		))
	} else if result.Ambiguous {
		_, err = ctx.session.ChannelMessageSendEmbed(ctx.messageCreate.ChannelID, &discordgo.MessageEmbed{
			URL:         result.URL,
			Title:       result.Title,
			Description: "This is a disambiguation page. Visit the link to see the possibilities.",
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: result.Thumbnail,
			},
			Author: &discordgo.MessageEmbedAuthor{
				IconURL: m.WikiSupport.Wikis[wikiID].IconURL,
				Name:    m.WikiSupport.Wikis[wikiID].Name,
			},
		})
	} else {
		_, err = ctx.session.ChannelMessageSendEmbed(ctx.messageCreate.ChannelID, &discordgo.MessageEmbed{
			URL:         result.URL,
			Title:       result.Title,
			Description: result.Text,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: result.Thumbnail,
			},
			Author: &discordgo.MessageEmbedAuthor{
				IconURL: m.WikiSupport.Wikis[wikiID].IconURL,
				Name:    m.WikiSupport.Wikis[wikiID].Name,
			},
		})
	}
	return err
}
