package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// LoginStates is a map of currently active login attempts
type LoginStates map[string]LoginState

// LoginState is an active login attempt
type LoginState struct {
	ID        string
	Time      time.Time
	ReturnURL string
}

// ErrLoginStateNotFound is the standard error for when a login state does not exist, or expired
var ErrLoginStateNotFound = errors.New(`state not found`)

var stateTimeout = time.Minute * 10

// New creates a new login state
func (s LoginStates) New(returnURL string) LoginState {
	state := LoginState{
		ID:        uuid.New().String(),
		Time:      time.Now(),
		ReturnURL: returnURL,
	}

	s[state.ID] = state
	return state
}

// Get retrieves a login state, if possible
func (s LoginStates) Get(id string) *LoginState {
	if state, ok := s[id]; !ok {
		return nil
	} else if time.Now().Sub(state.Time) > stateTimeout {
		s.Clear(id)
		return nil
	} else {
		return &state
	}
}

// Clear deletes a login state
func (s LoginStates) Clear(id string) {
	delete(s, id)
}
