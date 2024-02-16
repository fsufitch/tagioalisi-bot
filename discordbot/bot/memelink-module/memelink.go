package memelink

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/fsufitch/tagioalisi-bot/db/acl-dao"
	"github.com/fsufitch/tagioalisi-bot/db/memes-dao"
	"github.com/fsufitch/tagioalisi-bot/log"
)

// Module is a meme link module implementing RegisterableModule
type Module struct {
	Log     *log.Logger
	MemeDAO *memes.DAO
	ACLDAO  *acl.DAO

	session *discordgo.Session
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "memelink" }

// Register adds this module to the Discord session
func (m *Module) Register(ctx context.Context, session *discordgo.Session) error {
	m.session = session
	cancel1 := m.session.AddHandler(m.handleCommand)
	cancel2 := m.session.AddHandler(m.handleLink)

	go func() {
		<-ctx.Done()
		m.Log.Infof("memelink module context done")
		cancel1()
		cancel2()
	}()
	return nil
}

func (m *Module) RegisterGuild(ctx context.Context, session *discordgo.Session, guildID string) error {
	return nil
}

