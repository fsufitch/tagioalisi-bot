package db

import (
	"github.com/fsufitch/tagioalisi-bot/db/acl-dao"
	"github.com/fsufitch/tagioalisi-bot/db/connection"
	"github.com/fsufitch/tagioalisi-bot/db/kv-dao"
	"github.com/fsufitch/tagioalisi-bot/db/memes-dao"
	"github.com/google/wire"
)

// ProvidePostgresDatabase contains all the necessary wire providers to use the Tagioalisi Postgres database
var ProvidePostgresDatabase = wire.NewSet(
	connection.ProvidePostgresDatabaseConnection,
	wire.Struct(new(kv.DAO), "*"),
	wire.Struct(new(memes.DAO), "*"),
	wire.Struct(new(acl.DAO), "*"),
)
