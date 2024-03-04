package acl

import (
	"github.com/fsufitch/tagioalisi-bot/db/connection"
)

// DAO exposes database ACL functionality
type DAO struct {
	Conn connection.DatabaseConnection
}
