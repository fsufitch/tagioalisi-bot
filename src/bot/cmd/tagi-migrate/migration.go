package main

import (
	"net/http"
	"time"

	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/db"
	"github.com/fsufitch/tagioalisi-bot/log"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/pkg/errors"
)

// MigrationWrapper holds the configuration needed to run a migration
type MigrationWrapper struct {
	Log   *log.Logger
	DBURL config.DatabaseString
}

// Run runs a migration using the configured values
func (w MigrationWrapper) Run() error {
	var err error

	// Workaround to use fs.FS with migrate: https://github.com/golang-migrate/migrate/issues/471#issuecomment-782442944
	source, err := httpfs.New(http.FS(db.MigrationFS), "migrations")
	if err != nil {
		return err
	}
	w.Log.Infof("Starting migration using db=%s", w.DBURL)

	var m *migrate.Migrate

	for i := 0; i < 5; i++ {
		m, err = migrate.NewWithSourceInstance("embeddedMigrations", source, string(w.DBURL))
		if err == nil {
			break
		}
		w.Log.Warningf("failed to start migration (attempt #%d): %v", i+1, err)
		if i < 4 {
			<-time.After(2 * time.Second)
		}
	}

	if err != nil {
		return errors.Wrap(err, "could not start migration")
	}
	if err := m.Up(); err == migrate.ErrNoChange {
		w.Log.Infof("No migration necessary")
	} else if err != nil {
		return errors.Wrap(err, "error during migration")
	}

	version, dirty, err := m.Version()
	w.Log.Infof("Migration done; version=%d dirty=%v err=%v", version, dirty, err)
	return nil
}
