package kv

import (
	"database/sql"
	"time"

	"github.com/fsufitch/tagioalisi-bot/db/connection"
)

// DAO exposes database KV functionality
type DAO struct {
	Conn connection.DatabaseConnection
}

// Entry is a simple container for a KV row
type Entry struct {
	Key       string
	Value     string
	Timestamp time.Time
}

// Set sets a key in the KV table
func (dao DAO) Set(key string, value string) error {
	tx, err := (*sql.DB)(dao.Conn).Begin()
	if err != nil {
		return err
	}

	now := time.Now()

	tx.QueryRow(`
		INSERT INTO kv (key, value, timestamp)
		VALUES ($1, $2, $3)
		ON CONFLICT (key)
			DO UPDATE SET value = $4, timestamp = $5
	`, key, value, now, value, now)

	return tx.Rollback()
}

// Get gets a key from the KV table
func (dao DAO) Get(key string) (*Entry, error) {
	tx, err := (*sql.DB)(dao.Conn).Begin()
	if err != nil {
		return nil, err
	}

	row := tx.QueryRow(`
		SELECT key, value, timestamp FROM kv WHERE key=$1
	`, key)

	output := &Entry{}
	err = row.Scan(&output.Key, &output.Value, &output.Timestamp)

	if err == sql.ErrNoRows {
		err = nil
		output = nil
	} else {
		return nil, err
	}

	return output, tx.Rollback()
}
