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
	log           *common.LoggerModule
	session       *discordgo.Session
}

// Start is a blocking function that holds the runtime of the Discord bot
func (b *DiscordBoarBot) Start() error {
	session, err := discordgo.New("Bot " + b.configuration.DiscordToken)
	if err != nil {
		return errors.Wrap(err, "Could not create bot session")
	}

	// TODO: injected handler system
	session.AddHandler(pingHandler)

	err = session.Open()
	if err != nil {
		return errors.Wrap(err, "Could not open communication to Discord server")
	}

	b.log.Info("Discord bot starting...")
	<-b.Stop
	b.log.Info("Discord bot graceful shutdown...")

	return session.Close()
}

func pingHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "pong!")
	}
}

// NewDiscordBoarBot creates a new Discord Boar Bot
func NewDiscordBoarBot(
	configuration *common.Configuration,
	log *common.LoggerModule,
) *DiscordBoarBot {
	log.Info("Initializing Discord bot")
	return &DiscordBoarBot{
		Stop:          make(chan bool, 1),
		configuration: configuration,
		log:           log,
	}
}
