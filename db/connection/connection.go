package connection

import (
	"database/sql"

	"github.com/fsufitch/discord-boar-bot/common"
	_ "github.com/lib/pq" // Inject database driver
)

// DatabaseConnection is a struct wrapping *sql.DB and injecting connections to dependent DAOs
type DatabaseConnection struct {
	db *sql.DB
}

// Transaction wraps db.Begin
func (d DatabaseConnection) Transaction() (*sql.Tx, error) {
	return d.db.Begin()
}

// NewDatabaseConnection creates a new DatabaseConnection
func NewDatabaseConnection(
	configuration *common.Configuration,
	logger *common.LogDispatcher,
) (*DatabaseConnection, error) {
	if configuration.RunMode != common.Bot {
		logger.Info("Not initializing database since run mode is not Bot")
		return nil, nil
	}
	db, err := sql.Open("postgres", configuration.DatabaseURL)

	if err != nil {
		return nil, err
	}

	return &DatabaseConnection{
		db: db,
	}, nil
}
