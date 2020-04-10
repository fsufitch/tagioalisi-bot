package wiki

import (
	"bytes"
	"fmt"

	"github.com/fsufitch/tagioalisi-bot/bot/util"
)

type wikiSupportStruct struct {
	wikis       map[string]wikiStruct
	defaultWiki string
}

type wikiStruct struct {
	id          string
	name        string
	defaultLang string
	url         func(lang string) string
}

var wikiSupport = wikiSupportStruct{
	defaultWiki: "wp",
	wikis: map[string]wikiStruct{
		"w": {
			id:          "w",
			name:        "Wikipedia",
			defaultLang: "en",
			url:         func(lang string) string { return fmt.Sprintf("https://%s.wikipedia.org/w/api.php", lang) },
		},
		"d": {
			id:          "d",
			name:        "Wiktionary",
			defaultLang: "en",
			url:         func(lang string) string { return fmt.Sprintf("https://%s.wiktionary.org/w/api.php", lang) },
		},
		// TODO: add more options?
	},
}

func (m *Module) showOptions(ctx commandContext) error {
	buf := bytes.NewBufferString("Options for -w:\n")

	for _, wiki := range wikiSupport.wikis {
		languageSupport := "no"
		if wiki.defaultLang != "" {
			languageSupport = fmt.Sprintf("yes (default: %s)", wiki.defaultLang)
		}
		fmt.Fprintf(buf, "%s - %s; Language support: %s\n", wiki.id, wiki.name, languageSupport)
	}

	return util.DiscordMessageSendRawBlock(ctx.session, ctx.messageCreate.ChannelID, buf.String())
}

func (m *Module) queryWiki(ctx commandContext, wikiID string, lang string, query string) error {
	return util.DiscordMessageSendRawBlock(ctx.session, ctx.messageCreate.ChannelID, fmt.Sprintf("Querying: %s %s -> %s", wikiID, lang, query))
}
