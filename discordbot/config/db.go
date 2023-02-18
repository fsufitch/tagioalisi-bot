ipackage config

import (
	"errors"
	"fmt"
	"os"
)

// DatabaseString is a string alias containing a database connection string
type DatabaseString string

// ProvideDatabaseStringFromEnvironment creates a DatabaseString from the environment, or errors when it's missing
func ProvideDatabaseStringFromEnvironment() (DatabaseString, error) {
	// 	POSTGRES_USER=tagioalisi
	// POSTGRES_PASSWORD=tagi_secret!7461
	// POSTGRES_DB=tagioalisi
	var (
		user, password, host, db string
		ok                       bool
	)
	if user, ok = os.LookupEnv("POSTGRES_USER"); !ok {
		return "", errors.New("missing env var: POSTGRES_USER")
	}
	if password, ok = os.LookupEnv("POSTGRES_PASSWORD"); !ok {
		return "", errors.New("missing env var: POSTGRES_PASSWORD")
	}
	if host, ok = os.LookupEnv("POSTGRES_HOST"); !ok {
	   	return "", errors.New("missing env var: POSTGRES_HOST")
	}
	if db, ok = os.LookupEnv("POSTGRES_DB"); !ok {
		return "", errors.New("missing env var: POSTGRES_DB")
	}
	databaseURL := fmt.Sprintf("postgres://%s:%s@db/%s?sslmode=disable", user, password, db)
	return DatabaseString(databaseURL), nil

}
