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
	CmdBootstrapper ApplicationCommandModuleBootstrapper
}

// Run is a blocking function that holds the runtime of the Discord bot
func (b TagioalisiBot) Run(ctx context.Context) error {
	session, err := discordgo.New("Bot " + string(b.Token))
	if err != nil {
		return errors.Wrap(err, "could not create bot session")
	}
	defer session.Close()

	err = session.Open()
	if err != nil {
		return errors.Wrap(err, "could not open communication to Discord server")
	}
	b.Log.Infof("init session opened")

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

	cancelGuildRegistrations := b.CmdBootstrapper.registerToAllGuilds(session)
	b.Log.Infof("init complete")

	<-ctx.Done()
	b.Log.Infof("bot context canceled, shutting down")
	cancelGuildRegistrations()
	return session.Close()
}
