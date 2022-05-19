package sockpuppet

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagioalisi-bot/log"
	"github.com/pkg/errors"
)

// Module is a bot module that sockpuppets messages from elsewhere
type Module struct {
	session *discordgo.Session
	done    bool
	Log     *log.Logger
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "sockpuppet" }

// Register adds this module to the Discord session
func (m *Module) Register(ctx context.Context, session *discordgo.Session) error {
	m.session = session
	go func() {
		<-ctx.Done()
		m.Log.Infof("sockpuppet module context done")
		m.done = true
	}()
	return nil
}

// ErrSendingNotPermitted is an error indicating sending a message is not allowed
var ErrSendingNotPermitted = errors.New("not allowed to sockpuppet in this channel")

// VerifyCanSend returns an error if the user is not allowed to sockpuppet
func (m *Module) VerifyCanSend(senderUserID string, channelID string) error {
	perm, err := m.session.UserChannelPermissions(senderUserID, channelID)
	if err != nil || perm&0x00002000 == 0 {
		// Manage Messages required
		return ErrSendingNotPermitted
	}
	return nil
}

// SendMessage is used to send a message via the sockpuppet
func (m *Module) SendMessage(channelID string, message string, senderUserID string) error {
	if err := m.VerifyCanSend(senderUserID, channelID); err != nil {
		return err
	}

	if _, err := m.session.ChannelMessageSend(channelID, message); err != nil {
		return errors.Wrap(err, "sockpuppet: could not send message")
	}
	return nil
}
