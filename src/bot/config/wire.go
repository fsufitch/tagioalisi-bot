package config

import "github.com/google/wire"

// EnvironmentProviderSet is a Wire provider set for environment configuration
var EnvironmentProviderSet = wire.NewSet(
	ProvideAESBlockFromEnvironment,
	ProvideApplicationIDFromOAuth2Config,
	ProvideAzureCredentialsFromEnvironment,
	ProvideBotModuleBlacklistFromEnvironment,
	ProvideBotWebAPIPortFromEnvironment,
	ProvideDatabaseStringFromEnvironment,
	ProvideDebugModeFromEnvironment,
	ProvideDiscordBotTokenFromEnvironment,
	ProvideDiscordLogChannelFromEnvironment,
	ProvideGRPCPortFromEnvironment,
	ProvideJWTHMACSecretFromEnvironment,
	ProvideManagedGroupPrefixFromEnvironment,
	ProvideOAuth2ConfigFromEnvironment,
	ProvideMerriamWebsterAPIKeyFromEnvironment,
	ProvideUserAgent,
	ProvideLaunchTime,
)
