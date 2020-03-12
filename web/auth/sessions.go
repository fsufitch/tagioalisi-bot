package auth

import (
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

// SessionStorage is an abstract place to store Discord OAuth2 Tokens
type SessionStorage interface {
	Get(id string) *Session
	New(*oauth2.Token) *Session
	Clear(id string)
	Verify(id string, r *http.Request) error
}

// Session contains the details of one user's login session
type Session struct {
	ID          string
	OAuth2Token *oauth2.Token
	// TODO: add verification data
}

// MemorySessionStorage is a map containing active user tokens
type MemorySessionStorage map[string]Session

// Get recovers a user session from memory
func (s MemorySessionStorage) Get(id string) *Session {
	if session, ok := s[id]; ok {
		return &session
	}
	return nil
}

// New saves a new user session in memory
func (s MemorySessionStorage) New(token *oauth2.Token) *Session {
	session := Session{
		ID:          uuid.New().String(),
		OAuth2Token: token,
	}
	s[session.ID] = session
	return &session
}

// Clear removes a user session from memory
func (s MemorySessionStorage) Clear(id string) {
	delete(s, id)
}

// Verify confirms the session is valid to use for a request
func (s MemorySessionStorage) Verify(id string, r *http.Request) error {
	// TODO
	return nil
}

// ProvideMemorySessionStorage creates a new in-memory session storage
func ProvideMemorySessionStorage() MemorySessionStorage {
	return MemorySessionStorage{}
}
