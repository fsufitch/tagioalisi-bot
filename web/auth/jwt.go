package auth

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/fsufitch/tagialisi-bot/config"
)

type JWTSupport struct {
	JWTSecret config.JWTHMACSecret
}

type Claims struct {
	SessionID string `json:"sid"`
}

func (c Claims) Valid() error {
	if c.SessionID == "" {
		return errors.New("empty session id")
	}
	return nil
}

func (s JWTSupport) EncodeSessionID(sessionID string) (encodedJWT string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, Claims{
		SessionID: sessionID,
	})

	return token.SignedString([]byte(s.JWTSecret))
}

func (s JWTSupport) DecodeSessionID(encodedJWT string) (sessionID string, err error) {
	var token *jwt.Token
	if token, err = jwt.ParseWithClaims(encodedJWT, &Claims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(s.JWTSecret), nil
	}); err != nil {
		return
	}

	switch c := token.Claims.(type) {
	case (*Claims):
		sessionID = c.SessionID
	default:
		err = fmt.Errorf("could not use custom claims type (claims was %s)", token.Claims)
	}
	return
}
