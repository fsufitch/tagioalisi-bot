package memes

import "github.com/fsufitch/discord-boar-bot/db/connection"

// MemeDAO is a database abstraction around the meme feature set
type MemeDAO struct {
	dbConn *connection.DatabaseConnection
}

type Meme struct {
	id int
}

// NewMemeDAO creates a new MemeDAO
func NewMemeDAO(dbConn *connection.DatabaseConnection) *MemeDAO {
	return &MemeDAO{dbConn: dbConn}
}
