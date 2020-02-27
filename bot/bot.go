package bot

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/discord-boar-bot/config"
	"github.com/fsufitch/discord-boar-bot/log"
	"github.com/pkg/errors"
)

// Bot is a general interface for a runnable bot
type Bot interface {
	Run(context.Context) error
}

// TagioalisiBot is the concrete implementation of Bot
type TagioalisiBot struct {
	Log             *log.Logger
	Modules         ModuleList
	ModuleBlacklist config.BotModuleBlacklist
	Token           config.DiscordBotToken
}

// Run is a blocking function that holds the runtime of the Discord bot
func (b TagioalisiBot) Run(ctx context.Context) error {
	session, err := discordgo.New("Bot " + string(b.Token))
	if err != nil {
		return errors.Wrap(err, "could not create bot session")
	}
	defer session.Close()

	for _, module := range b.Modules {
		if _, ok := b.ModuleBlacklist[module.Name()]; ok {
			b.Log.Infof("not registering blacklisted module `%s`", module.Name())
			continue
		}
		if err = module.Register(ctx, session); err != nil {
			return errors.Wrap(err, "error registering bot module: "+module.Name())
		}
		b.Log.Infof("registered module `%s`", module.Name())
	}

	err = session.Open()
	if err != nil {
		return errors.Wrap(err, "could not open communication to Discord server")
	}
	b.Log.Infof("bot initialized and listening")
	<-ctx.Done()
	b.Log.Infof("bot context canceled, shutting down")
	return session.Close()
}
