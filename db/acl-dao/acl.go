package acl

import (
	"github.com/fsufitch/tagialisi-bot/db/connection"
)

// DAO exposes database ACL functionality
type DAO struct {
	Conn connection.DatabaseConnection
}
