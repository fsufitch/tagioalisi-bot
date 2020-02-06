package db

import (
	"github.com/fsufitch/discord-boar-bot/db/acl-dao"
	"github.com/fsufitch/discord-boar-bot/db/connection"
	"github.com/fsufitch/discord-boar-bot/db/kv-dao"
	"github.com/fsufitch/discord-boar-bot/db/memes-dao"
	"github.com/google/wire"
)

// ProvidePostgresDatabase contains all the necessary wire providers to use the Boar Bot Postgres database
var ProvidePostgresDatabase = wire.NewSet(
	connection.ProvidePostgresDatabaseConnection,
	wire.Struct(new(kv.DAO), "*"),
	wire.Struct(new(memes.DAO), "*"),
	wire.Struct(new(acl.DAO), "*"),
)
