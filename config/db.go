package config

import (
	"errors"
	"os"
)

// DatabaseString is a string alias containing a database connection string
type DatabaseString string

// ProvideDatabaseStringFromEnvironment creates a DatabaseString from the environment, or errors when it's missing
func ProvideDatabaseStringFromEnvironment() (DatabaseString, error) {
	if envValue, ok := os.LookupEnv("DATABASE_URL"); ok {
		return DatabaseString(envValue), nil
	}
	return "", errors.New("missing env var: DATABASE_URL")

}

// MigrationDir is a directory where to find database migrations
type MigrationDir string

// ProvideMigrationDirFromEnvironment creates a MigrationDir from the environment, or errors when it's missing
func ProvideMigrationDirFromEnvironment() (MigrationDir, error) {
	if envValue, ok := os.LookupEnv("MIGRATION_DIR"); ok {
		return MigrationDir(envValue), nil
	}
	return "", errors.New("missing env var: MIGRATION_DIR")
}
