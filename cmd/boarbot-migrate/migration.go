package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigration() error {
	// sourceURL := fmt.Sprintf("file://%s", r.Configuration.MigrationDir)
	// r.Logger.Info(fmt.Sprintf("Starting migration using source=%s and db=%s",
	// 	sourceURL, r.Configuration.DatabaseURL))
	// m, err := migrate.New(sourceURL, r.Configuration.DatabaseURL)
	// if err != nil {
	// 	return errors.Wrap(err, "could not begin migration")
	// }
	// if err := m.Up(); err == migrate.ErrNoChange {
	// 	r.Logger.Info("No migration necessary")
	// } else if err != nil {
	// 	return errors.Wrap(err, "error during migration")
	// }

	// version, dirty, err := m.Version()

	// r.Logger.Info(fmt.Sprintf("Migration done; version=%d dirty=%v err=%v", version, dirty, err))

	// <-time.After(1 * time.Second)
	return nil
}
