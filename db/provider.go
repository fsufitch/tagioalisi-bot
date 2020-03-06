package db

import (
	"github.com/fsufitch/tagialisi-bot/db/acl-dao"
	"github.com/fsufitch/tagialisi-bot/db/connection"
	"github.com/fsufitch/tagialisi-bot/db/kv-dao"
	"github.com/fsufitch/tagialisi-bot/db/memes-dao"
	"github.com/google/wire"
)

// ProvidePostgresDatabase contains all the necessary wire providers to use the Tagioalisi Postgres database
var ProvidePostgresDatabase = wire.NewSet(
	connection.ProvidePostgresDatabaseConnection,
	wire.Struct(new(kv.DAO), "*"),
	wire.Struct(new(memes.DAO), "*"),
	wire.Struct(new(acl.DAO), "*"),
)
