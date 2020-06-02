package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/fsufitch/tagioalisi-bot/security"
	"github.com/google/uuid"
)

var stateTimeout = time.Minute * 10

// LoginStates is a map of currently active login attempts
type LoginStates map[string]LoginState

// LoginState is an active login attempt
type LoginState struct {
	ID        string    `json:"id"`
	Time      time.Time `json:"time"`
	ReturnURL string    `json:"return_url"`
}

// NewLoginState creates a new LoginState with a return URL
func NewLoginState(returnURL string) LoginState {
	return LoginState{
		ID:        uuid.New().String(),
		Time:      time.Now(),
		ReturnURL: returnURL,
	}
}

// LoginStateFromStateParam extracts a LoginState from the state string; encrypts if AES is available
func LoginStateFromStateParam(stateStr string, aes security.AESSupport) (*LoginState, error) {
	bytes, err := base64.StdEncoding.DecodeString(stateStr)
	if err != nil {
		return nil, err
	}

	if aes.Ready() {
		bytes, err = aes.Decrypt(bytes)
		if err != nil {
			return nil, err
		}
	}

	state := LoginState{}
	if err = json.Unmarshal(bytes, &state); err != nil {
		return nil, err
	}

	if time.Now().Sub(state.Time) > stateTimeout {
		return nil, errors.New("state timed out, please try again")
	}

	return &state, nil
}

// ToStateParam turns the LoginState into a string useable as a state param; encrypts if AES is available
func (s LoginState) ToStateParam(aes security.AESSupport) (string, error) {
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	if aes.Ready() {
		bytes, err = aes.Encrypt(bytes)
		if err != nil {
			return "", err
		}
	}

	encoded := base64.StdEncoding.EncodeToString(bytes)
	return encoded, nil
}

// ErrLoginStateNotFound is the standard error for when a login state does not exist, or expired
var ErrLoginStateNotFound = errors.New(`state not found`)

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
