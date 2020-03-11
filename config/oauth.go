package config

import (
	"os"

	"golang.org/x/oauth2"
)

var discordOAuth2Scopes = []string{"identify"}

// OAuth2Config is a shared object used to run OAuth2 login; if nil, OAuth2 login is disabled
type OAuth2Config *oauth2.Config

// ProvideOAuth2ConfigFromEnvironment builds an OAuth2Config from the available environment variables
func ProvideOAuth2ConfigFromEnvironment() OAuth2Config {
	var (
		clientID      = os.Getenv("OAUTH_CLIENT_ID")
		clientSecret  = os.Getenv("OAUTH_CLIENT_SECRET")
		authEndpoint  = os.Getenv("OAUTH_AUTH_ENDPOINT")
		tokenEndpoint = os.Getenv("OAUTH_TOKEN_ENDPOINT")
		redirectURL   = os.Getenv("OAUTH_REDIRECT_URL")
	)

	if clientID == "" || clientSecret == "" || authEndpoint == "" || tokenEndpoint == "" || redirectURL == "" {
		return nil
	}

	return OAuth2Config(&oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: tokenEndpoint,
			AuthURL:  authEndpoint,
		},
		RedirectURL: redirectURL,
		Scopes:      discordOAuth2Scopes,
	})
}

// JWTHMACSecret is the secret bytes to be used when using HMAC signing for JWT
type JWTHMACSecret []byte

// ProvideJWTHMACSecretFromEnvironment creates a JWTHMACSecret from the environment
func ProvideJWTHMACSecretFromEnvironment() JWTHMACSecret {
	return JWTHMACSecret(os.Getenv("JWT_HMAC_SECRET"))
}
