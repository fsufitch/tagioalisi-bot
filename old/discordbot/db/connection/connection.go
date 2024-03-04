package connection

import (
	"database/sql"
	"time"

	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/log"
	_ "github.com/lib/pq" // Inject database driver
)

// DatabaseConnection is a struct wrapping *sql.DB and injecting connections to dependent DAOs
type DatabaseConnection *sql.DB

// ProvidePostgresDatabaseConnection creates a new concrete postgres DatabaseConnection
func ProvidePostgresDatabaseConnection(log *log.Logger, dbURL config.DatabaseString) (DatabaseConnection, func(), error) {
	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", string(dbURL))
		if err == nil {
			err = db.Ping()
		}
		if err == nil {
			break
		}
		log.Warningf("failed to connect to database (attempt #%d): %v", i+1, err)
		if i < 4 {
			<-time.After(2 * time.Second)
		}
	}

	if err != nil {
		return nil, nil, err
	}

	close := func() { db.Close() }

	log.Infof("database connection successful")

	return DatabaseConnection(db), close, nil
}
