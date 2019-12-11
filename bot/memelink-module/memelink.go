package memelink

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/discord-boar-bot/common"
	"github.com/fsufitch/discord-boar-bot/db/connection"
	"github.com/fsufitch/discord-boar-bot/db/memes-dao"
	"github.com/urfave/cli/v2"
)

// Module is a meme link module implementing RegisterableModule
type Module struct {
	session    *discordgo.Session
	db         *connection.DatabaseConnection
	log        *common.LogDispatcher
	memeDAO    *memes.DAO
	cliAppLazy *cli.App
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "memelink" }

// Register adds this module to the Discord session
func (m *Module) Register(session *discordgo.Session) error {
	m.session = session
	// Nothing to do here, it does not react to anything (yet)
	m.session.AddHandler(m.handleCommand)
	m.session.AddHandler(m.handleLink)
	return nil
}

// NewModule creates a new memelink module
func NewModule(db *connection.DatabaseConnection, log *common.LogDispatcher, memeDAO *memes.DAO) *Module {
	return &Module{
		db:      db,
		log:     log,
		memeDAO: memeDAO,
	}
}
