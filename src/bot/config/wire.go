package config

import "github.com/google/wire"

// EnvironmentProviderSet is a Wire provider set for environment configuration
var EnvironmentProviderSet = wire.NewSet(
	ProvideAESBlockFromEnvironment,
	ProvideAzureCredentialsFromEnvironment,
	ProvideBotModuleBlacklistFromEnvironment,
	ProvideBotWebAPIPortFromEnvironment,
	ProvideDatabaseStringFromEnvironment,
	ProvideDebugModeFromEnvironment,
	ProvideDiscordBotTokenFromEnvironment,
	ProvideDiscordLogChannelFromEnvironment,
	ProvideJWTHMACSecretFromEnvironment,
	ProvideManagedGroupPrefixFromEnvironment,
	ProvideOAuth2ConfigFromEnvironment,
	ProvideMerriamWebsterAPIKeyFromEnvironment,
	ProvideUserAgent,
)
