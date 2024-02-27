package interactions

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"

	"github.com/fsufitch/tagioalisi-bot/bot/util"
	"github.com/fsufitch/tagioalisi-bot/log"
)

type ApplicationCommandWrapper struct {
	Command *discordgo.ApplicationCommand
	Logger  *log.Logger

	Handlers             []HandlerFunc
	AutocompleteHandlers []AutocompleteHandlerFunc
}

type RegisterableApplicationCommandModule interface {
	Register(session *discordgo.Session, guildID string) (*discordgo.ApplicationCommand, error)
}

type HandlerFunc func(*InteractionWrapper)

type AutocompleteHandlerFunc func(*InteractionWrapper) []string

type ApplicationCommandRegistration struct {
	GuildID            string
	ApplicationCommand *discordgo.ApplicationCommand
	RegisterError      error
}

func (acw ApplicationCommandWrapper) Register(session *discordgo.Session, guildID string) (ccmd *discordgo.ApplicationCommand, cancel func(), err error) {
	cancelCh := util.FunctionCallBuffer()
	cancel = func() { close(cancelCh) }

	if acw.Command.ApplicationID == "" {
		return nil, cancel, fmt.Errorf("app command has no app ID set; cmd='%s' guild=%s", acw.Command.Name, acw.Command.GuildID)
	}
	ccmd, err = session.ApplicationCommandCreate(acw.Command.ApplicationID, guildID, acw.Command)
	if err != nil {
		return ccmd, cancel, err
	}

	cancelCh <- session.AddHandler(acw.HandleInteractionAutocomplete)
	cancelCh <- session.AddHandler(acw.HandleInteractionCommand)
	return
}

func (acw *ApplicationCommandWrapper) AutoRegisterToSessionGuilds(session *discordgo.Session) (registrationCh <-chan ApplicationCommandRegistration, cancel func()) {
	outputRegistrationCh := make(chan ApplicationCommandRegistration, 128)

	// cancelCh receives functions to call, then once the channel is closed, calls them
	cancelCh := util.FunctionCallBuffer()

	cancelGuildCreate := session.AddHandler(func(session2 *discordgo.Session, event *discordgo.GuildCreate) {
		ccmd, cancelRegistration, err := acw.Register(session, event.Guild.ID)
		outputRegistrationCh <- ApplicationCommandRegistration{
			GuildID:            event.Guild.ID,
			ApplicationCommand: ccmd,
			RegisterError:      err,
		}
		cancelCh <- cancelRegistration
	})

	cancelCh <- cancelGuildCreate

	return outputRegistrationCh, func() { close(cancelCh) }
}

func (acw *ApplicationCommandWrapper) HandleInteractionCommand(session *discordgo.Session, inter *discordgo.InteractionCreate) {
	if inter.Type != discordgo.InteractionApplicationCommand {
		return
	}

	iw, err := WrapInteraction(session, inter, acw.Logger)
	if err != nil {
		acw.Logger.Errorf("failed to wrap interaction: %v", err)
		return
	}
	defer acw.recoverPanic(iw)

	for _, handlerFunc := range acw.Handlers {
		handlerFunc(iw)
		if iw.Acknowledged() {
			break
		}
	}

	if !iw.Acknowledged() {
		panic("interaction was not acknowledged")
	}
}

func (acw *ApplicationCommandWrapper) HandleInteractionAutocomplete(session *discordgo.Session, inter *discordgo.InteractionCreate) {
	if inter.Type != discordgo.InteractionApplicationCommandAutocomplete {
		return
	}

	iw, err := WrapInteraction(session, inter, acw.Logger)
	if err != nil {
		acw.Logger.Errorf("failed to wrap interaction: %v", err)
		return
	}
	defer acw.recoverPanic(iw)

	var choices []string
	for _, handlerFunc := range acw.AutocompleteHandlers {
		choices = handlerFunc(iw)
		if choices != nil {
			break
		}
	}

	iw.RespondAutocomplete(choices...)

	if !iw.Acknowledged() {
		acw.Logger.Errorf(
			"autocomplete interaction was not acknowledged: gid=%s cid=%s ",
			iw.Interaction().GuildID,
			iw.Interaction().ChannelID,
		)
		iw.RespondAutocomplete()
	}
}

func (acw *ApplicationCommandWrapper) recoverPanic(iw *InteractionWrapper) {
	panicID := new(string)
	var panicReason any
	if panicReason = recover(); panicReason == nil {
		return
	}
	*panicID = uuid.NewString()

	indentedStack := strings.ReplaceAll(string(debug.Stack()), "\n", "\n    ")
	acw.Logger.Errorf("Panic [%s]: %+v\nStack:\n%s", *panicID, panicReason, indentedStack)
	iw.RespondError(InteractionError{
		Title:       "Unexpeced error (panic)",
		Description: fmt.Sprintf("ID: `%s`\n\nReason: ```%+v```\n\nCheck logs for more details.", *panicID, panicReason),
	})

	defer func() {
		if repanicReason := recover(); repanicReason != nil {
			origPanicID := "UNKNOWN"
			if panicID != nil {
				origPanicID = *panicID
			}
			acw.Logger.Errorf("panicked while handling panic: [%s] reason=%v orig=%v", origPanicID, repanicReason, panicReason)
		}
	}()
}
