package connection

import (
	"database/sql"

	"github.com/fsufitch/discord-boar-bot/config"
	_ "github.com/lib/pq" // Inject database driver
)

// DatabaseConnection is a struct wrapping *sql.DB and injecting connections to dependent DAOs
type DatabaseConnection *sql.DB

// ProvidePostgresDatabaseConnection creates a new concrete postgres DatabaseConnection
func ProvidePostgresDatabaseConnection(dbURL config.DatabaseString) (DatabaseConnection, func(), error) {
	db, err := sql.Open("postgres", string(dbURL))
	close := func() { db.Close() }

	if err != nil {
		return nil, nil, err
	}

	return DatabaseConnection(db), close, nil
}
