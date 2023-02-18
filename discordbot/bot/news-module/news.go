package news

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagioalisi-bot/azure"
	"github.com/fsufitch/tagioalisi-bot/bot/util"
	"github.com/fsufitch/tagioalisi-bot/log"
)

// Module is a news module implementing RegisterableModule
type Module struct {
	Log  *log.Logger
	News azure.NewsSearch

	session *discordgo.Session
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "news" }

// Register adds this module to the Discord session
func (m *Module) Register(ctx context.Context, session *discordgo.Session) error {
	m.session = session

	cancel := m.session.AddHandler(m.handleCommand)
	go func() {
		<-ctx.Done()
		m.Log.Infof("news module context done")
		cancel()
	}()

	return nil
}

// DoSearch performs an actual news search
func (m *Module) DoSearch(
	ctx context.Context,
	session *discordgo.Session,
	channelID string,
	query string,
	count int,
	verbose bool,
) error {
	m.Log.Debugf("news: searching for `%s`", query)

	if !m.News.Ready() {
		return util.DiscordMessageSendRawBlock(session, channelID, "News API not properly instantiated. Sorry! :(")
	}
	m.Log.Debugf("news: api ready")

	results, err := m.News.Search(ctx, query, int32(count))
	if err != nil {
		return err
	}
	m.Log.Debugf("news: got %d results for %s", results.Len(), query)

	var f formatter
	if verbose {
		f = verboseFormatter
	} else {
		f = compactFormatter
	}

	return f(session, channelID, results)
}
