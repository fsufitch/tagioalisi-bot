package main

import (
	"fmt"
	"time"

	"github.com/fsufitch/tagialisi-bot/config"
	"github.com/fsufitch/tagialisi-bot/log"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
)

// MigrationWrapper holds the configuration needed to run a migration
type MigrationWrapper struct {
	Log   *log.Logger
	Dir   config.MigrationDir
	DBURL config.DatabaseString
}

// Run runs a migration using the configured values
func (w MigrationWrapper) Run() error {
	sourceURL := fmt.Sprintf("file://%s", w.Dir)
	w.Log.Infof("Starting migration using source=%s and db=%s", sourceURL, w.DBURL)

	var m *migrate.Migrate
	var err error

	for i := 0; i < 5; i++ {
		m, err = migrate.New(sourceURL, string(w.DBURL))
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
