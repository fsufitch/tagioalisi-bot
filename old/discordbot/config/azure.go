package config

import "os"

// AzureNewsSearchAPIKey is the API key for using the Azure News Search API
type AzureNewsSearchAPIKey string

// ProvideAzureCredentialsFromEnvironment retrieves the API key from the environment
func ProvideAzureCredentialsFromEnvironment() AzureNewsSearchAPIKey {
	return AzureNewsSearchAPIKey(os.Getenv("AZURE_NEWS_SEARCH_KEY"))
}
