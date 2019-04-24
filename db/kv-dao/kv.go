package kv

import (
	"database/sql"
	"github.com/fsufitch/discord-boar-bot/db/connection"
	"time"
)

// KeyValueDAO exposes database KV functionality
type KeyValueDAO struct {
	dbConn *connection.DatabaseConnection
}

// KeyValueEntry is a simple container for a KV row
type KeyValueEntry struct {
	Key       string
	Value     string
	Timestamp time.Time
}

// Set sets a key in the KV table
func (dao KeyValueDAO) Set(key string, value string) error {
	tx, err := dao.dbConn.Transaction()
	if err != nil {
		return err
	}

	tx.QueryRow(`
		SELECT key, value, timestamp FROM kv WHERE key=$1
	`, key)

	return tx.Rollback()
}

// Get gets a key from the KV table
func (dao KeyValueDAO) Get(key string) (*KeyValueEntry, error) {
	tx, err := dao.dbConn.Transaction()
	if err != nil {
		return nil, err
	}

	row := tx.QueryRow(`
		SELECT key, value, timestamp FROM kv WHERE key=$1
	`, key)

	output := &KeyValueEntry{}
	err = row.Scan(&output.Key, &output.Value, &output.Timestamp)

	if err == sql.ErrNoRows {
		err = nil
		output = nil
	} else {
		return nil, err
	}

	return output, tx.Rollback()
}

// NewKeyValueDAO creates a new KeyValueDAO
func NewKeyValueDAO(
	dbConn *connection.DatabaseConnection,
) *KeyValueDAO {
	return &KeyValueDAO{
		dbConn: dbConn,
	}
}
