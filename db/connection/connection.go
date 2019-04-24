package connection

import "database/sql"
import _ "github.com/lib/pq" // Inject database driver

import "github.com/fsufitch/discord-boar-bot/common"

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
) (*DatabaseConnection, error) {
	db, err := sql.Open("postgres", configuration.DatabaseURL)

	if err != nil {
		return nil, err
	}

	return &DatabaseConnection{
		db: db,
	}, nil
}
