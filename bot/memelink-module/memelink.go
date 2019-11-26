package memelink

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/discord-boar-bot/db/connection"
)

// Module is a meme link module implementing RegisterableModule
type Module struct {
	session *discordgo.Session
	db      *connection.DatabaseConnection
}

// Name returns the name of the module, for blacklisting
func (m Module) Name() string { return "memelink" }

// Register adds this module to the Discord session
func (m *Module) Register(session *discordgo.Session) error {
	m.session = session
	// Nothing to do here, it does not react to anything (yet)
	return nil
}

// NewModule creates a new memelink module
func NewModule(db *connection.DatabaseConnection) *Module {
	return &Module{
		db: db,
	}
}
