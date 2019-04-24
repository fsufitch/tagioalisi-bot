package db

import (
	"github.com/fsufitch/discord-boar-bot/db/connection"
	"github.com/fsufitch/discord-boar-bot/db/kv-dao"
	"github.com/google/wire"
)

// DBProviderSet contains all the necessary wire providers to use the Boar Bot database
var DBProviderSet = wire.NewSet(
	connection.NewDatabaseConnection,
	kv.NewKeyValueDAO,
)
