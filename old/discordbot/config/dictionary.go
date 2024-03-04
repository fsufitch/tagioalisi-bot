package config

import (
	"os"
)

// MerriamWebsterAPIKey is a string alias containing the M-W API key
type MerriamWebsterAPIKey string

// ProvideMerriamWebsterAPIKeyFromEnvironment creates a MerriamWebsterAPIKey from the environment
func ProvideMerriamWebsterAPIKeyFromEnvironment() MerriamWebsterAPIKey {
	return MerriamWebsterAPIKey(os.Getenv("MERRIAM_WEBSTER_DICTIONARY_KEY"))
}
