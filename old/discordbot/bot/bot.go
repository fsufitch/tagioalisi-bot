package bot

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"

	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/log"
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

	// Register global modules
	session.AddHandlerOnce(func(s *discordgo.Session, event *discordgo.Ready) {
		for _, module := range b.Modules {
			if _, ok := b.ModuleBlacklist[module.Name()]; ok {
				b.Log.Infof("not registering blacklisted module (global): module=%+v", module.Name())
				continue
			}
			if err = module.Register(ctx, session); err != nil {
				b.Log.Errorf("error registering bot module (global): module=%s -- %s", module.Name(), err)
			} else {
				b.Log.Infof("registered module (global): module=%+v", module.Name())
			}
		}
	})


	session.AddHandler(func(s *discordgo.Session, event *discordgo.GuildCreate) {
		b.Log.Infof("registering guild: %s (%s)", event.Guild.Name, event.Guild.ID)
		for _, module := range b.Modules {
			if _, ok := b.ModuleBlacklist[module.Name()]; ok {
				b.Log.Infof("not registering blacklisted module: module=%+v guild=%v", module.Name(), event.Guild.ID)
				continue
			}
			if err = module.RegisterGuild(ctx, session, event.Guild.ID); err != nil {
				b.Log.Errorf("error registering bot module to guild: module=%s guild=%s -- %s", module.Name(), event.Guild.ID, err)
			} else {
				b.Log.Infof("registered module to guild: module=%+v guild=%v", module.Name(), event.Guild.ID)
			}
		}
	})

	err = session.Open()
	if err != nil {
		return errors.Wrap(err, "could not open communication to Discord server")
	}
	b.Log.Infof("bot initialized and listening")

	<-ctx.Done()
	b.Log.Infof("bot context canceled, shutting down")
	return session.Close()
}
