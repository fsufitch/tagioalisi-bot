package dictionary

import (
	"context"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagioalisi-bot/bot/util"
	"github.com/fsufitch/tagioalisi-bot/log"
)

// Module is a dictionary module that works with the Merriam-Webster dictionary
type Module struct {
	Log    *log.Logger
	Client Client
}

// Name returns the name of the module
func (m Module) Name() string { return "dictionary" }

// Register adds this module to the Discord session
func (m *Module) Register(ctx context.Context, session *discordgo.Session) error {
	cancel := session.AddHandler(m.handleCommand)
	go func() {
		<-ctx.Done()
		m.Log.Infof("dictionary module context done")
		cancel()
	}()
	return nil
}

func (m *Module) define(ctx commandContext, word string) error {
	word = strings.TrimSpace(strings.ToLower(word))
	results, suggestions, err := m.Client.SearchCollegiate(word)
	if err != nil {
		m.Log.Errorf("failed SearchCollegiate: %v", err)
		return errorFormatter(ctx.session, ctx.messageCreate.ChannelID, word, "Unexpected error", err.Error())
	}

	if len(results) > 0 {
		if err := collegiateResultFormatter(ctx.session, ctx.messageCreate.ChannelID, word, results); err != nil {
			m.Log.Errorf("dictionary: error formatting: %v", err)
			return util.DiscordMessageSendRawBlock(ctx.session, ctx.messageCreate.ID, err.Error())
		}
	} else if len(suggestions) > 0 {
		if err := suggestionFormatter(ctx.session, ctx.messageCreate.ChannelID, word, suggestions); err != nil {
			m.Log.Errorf("dictionary: error formatting suggestions: %v", err)
			return util.DiscordMessageSendRawBlock(ctx.session, ctx.messageCreate.ID, err.Error())
		}
	} else {
		return errorFormatter(ctx.session, ctx.messageCreate.ChannelID, word, "No results found", "")
	}

	return nil
}
