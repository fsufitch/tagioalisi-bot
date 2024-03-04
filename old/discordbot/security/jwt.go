package security

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fsufitch/tagioalisi-bot/config"
)

// JWTData is a flat structure of the data stored in a session JWT
type JWTData struct {
	SessionID   string
	AccessToken string
	Expiration  time.Time
	Username    string
	UserID      string
	AvatarURL   string
}

type claims struct {
	Username  string `json:"un"`
	Token     string `json:"tok"`
	AvatarURL string `json:"av"`
	jwt.StandardClaims
}

// JWTSupport is an object that can encode/decode JWT tokens and their nested data
type JWTSupport struct {
	JWTHMACSecret config.JWTHMACSecret
	AES           AESSupport
}

// CreateTokenString encodes the given JWT data into a JWT; the access token is encrypted if a cipher block is available
func (j JWTSupport) CreateTokenString(data JWTData) (string, error) {
	token := data.AccessToken
	if j.AES.Ready() {
		encToken, err := j.AES.Encrypt([]byte(token))
		if err != nil {
			return "", err
		}
		token = base64.StdEncoding.EncodeToString(encToken)
	}

	c := claims{
		Username:  data.Username,
		Token:     token,
		AvatarURL: data.AvatarURL,
		StandardClaims: jwt.StandardClaims{
			Id:        data.SessionID,
			Subject:   data.UserID,
			ExpiresAt: data.Expiration.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, c)

	return jwtToken.SignedString([]byte(j.JWTHMACSecret))
}

// ExtractJWTData extracts session data from a JWT; the access token is decrypted if a cipher block is available
func (j JWTSupport) ExtractJWTData(data string) (*JWTData, error) {
	token, err := jwt.ParseWithClaims(data, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.JWTHMACSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if c, ok := token.Claims.(*claims); ok {
		accessToken := c.Token
		if j.AES.Ready() {
			raw, err := base64.StdEncoding.DecodeString(accessToken)
			if err != nil {
				return nil, err
			}
			decToken, err := j.AES.Decrypt(raw)
			if err != nil {
				return nil, err
			}
			accessToken = string(decToken)
		}
		return &JWTData{
			AccessToken: accessToken,
			SessionID:   c.Id,
			Expiration:  time.Unix(c.ExpiresAt, 0),
			UserID:      c.Subject,
			Username:    c.Username,
			AvatarURL:   c.AvatarURL,
		}, nil
	}
	return nil, errors.New("claims type assertion failed")
}

// Errors pertinent to extracting JWTs
var (
	ErrNoJWTFound = errors.New("no jwt found in request")
	ErrJWTExpired = errors.New("jwt expired")
)

// ExtractJWTFromRequest extracts JWT info from a request and validates expiration
func (j JWTSupport) ExtractJWTFromRequest(r *http.Request) (*JWTData, error) {
	jwtString := ""
	authHeader := r.Header.Get("Authorization")

	if r.URL.Query().Get("jwt") != "" {
		jwtString = r.URL.Query().Get("jwt")
	} else if strings.HasPrefix(authHeader, "Bearer ") && authHeader[7:] != "" {
		jwtString = authHeader[7:]
	} else {
		return nil, ErrNoJWTFound
	}

	jwtData, err := j.ExtractJWTData(jwtString)
	if err != nil {
		return nil, err
	}

	if jwtData.Expiration.Before(time.Now()) {
		return nil, ErrJWTExpired
	}

	return jwtData, nil
}
