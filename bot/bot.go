package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/discord-boar-bot/common"
	"github.com/pkg/errors"
)

// DiscordBoarBot encapsulates the Discord bot process
type DiscordBoarBot struct {
	Stop          chan bool
	configuration *common.Configuration
	log           *common.LogDispatcher
	session       *discordgo.Session
	modules       ModuleRegistry
}

// Start is a blocking function that holds the runtime of the Discord bot
func (b *DiscordBoarBot) Start() error {
	session, err := discordgo.New("Bot " + b.configuration.DiscordToken)
	if err != nil {
		return errors.Wrap(err, "Could not create bot session")
	}

	for _, module := range b.modules {
		if err = module.Register(session); err != nil {
			return errors.Wrap(err, "Error registering "+module.Name())
		}
		b.log.Info("Registered bot module: " + module.Name())
	}

	err = session.Open()
	if err != nil {
		return errors.Wrap(err, "Could not open communication to Discord server")
	}

	b.log.Info("Discord bot starting...")
	<-b.Stop
	b.log.Info("Discord bot graceful shutdown...")

	return session.Close()
}

// NewDiscordBoarBot creates a new Discord Boar Bot
func NewDiscordBoarBot(
	configuration *common.Configuration,
	log *common.LogDispatcher,
	modules ModuleRegistry,
) *DiscordBoarBot {
	if configuration.RunMode != common.Bot {
		log.Info("Not initializing bot since run mode is not Bot")
		return nil
	}
	log.Info("Initializing Discord bot")
	return &DiscordBoarBot{
		Stop:          make(chan bool, 1),
		configuration: configuration,
		log:           log,
		modules:       modules,
	}
}
