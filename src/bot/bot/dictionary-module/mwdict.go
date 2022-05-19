package dictionary

import (
	"errors"

	"github.com/fsufitch/tagioalisi-bot/config"
	mwdict "github.com/fsufitch/tagioalisi-bot/merriam-webster"
)

// Client is a rename of mwDict.Client for dependency injection
type Client mwdict.Client

// NewClient creates a new client for the dictionary module
func NewClient(apiKey config.MerriamWebsterAPIKey, userAgent config.UserAgent) (*mwdict.BasicClient, error) {
	if apiKey == "" {
		return nil, errors.New("no Merriam-Webster API key found")
	}
	return mwdict.NewBasicClient(string(apiKey), string(userAgent)), nil
}
