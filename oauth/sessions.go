package oauth

import (
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

// SessionStorage is an abstract place to store Discord OAuth2 Tokens
type SessionStorage interface {
	Get(id string) *oauth2.Token
	Set(token *oauth2.Token) string
	Clear(id string)
}

// MemorySessionStorage is a map containing active user tokens
type MemorySessionStorage map[string]*oauth2.Token

// Get recovers a user session from memory
func (s MemorySessionStorage) Get(id string) *oauth2.Token {
	if token, ok := s[id]; ok {
		return token
	}
	return nil
}

// Set saves a new user session in memory
func (s MemorySessionStorage) Set(token *oauth2.Token) string {
	sessionID := uuid.New().String()
	s[sessionID] = token
	return sessionID
}

// Clear removes a user session from memory
func (s MemorySessionStorage) Clear(id string) {
	delete(s, id)
}

// ProvideMemorySessionStorage creates a new in-memory session storage
func ProvideMemorySessionStorage() MemorySessionStorage {
	return MemorySessionStorage{}
}
