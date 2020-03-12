package usersession

import "github.com/bwmarrin/discordgo"

var newSession = func(token string) (*discordgo.Session, error) {
	return discordgo.New("Bearer " + token)
}

type websocketSupport interface {
	Open() error
	Close() error
}

// Identity is a session for getting information about users
type Identity interface {
	websocketSupport
	User(id string) (*discordgo.User, error)
}

// NewIdentity creates an IdentitySession from a token
func NewIdentity(token string) (Identity, error) {
	s, err := newSession(token)
	return Identity(s), err
}
