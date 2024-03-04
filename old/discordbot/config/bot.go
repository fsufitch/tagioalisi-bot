package config

import (
	"errors"
	"os"
	"strings"
)

// DiscordBotToken is what it says on the tin
type DiscordBotToken string

// ProvideDiscordBotTokenFromEnvironment creates a DiscordBotToken from the environment
func ProvideDiscordBotTokenFromEnvironment() (DiscordBotToken, error) {
	if token, ok := os.LookupEnv("DISCORD_TOKEN"); ok {
		return DiscordBotToken(token), nil
	}
	return "", errors.New("DISCORD_TOKEN not set")
}

// BotModuleBlacklist is a set of modules to not enable
type BotModuleBlacklist map[string]interface{}

// ProvideBotModuleBlacklistFromEnvironment creates a BotModuleBlacklist from the environment
func ProvideBotModuleBlacklistFromEnvironment() BotModuleBlacklist {
	blacklist := map[string]interface{}{}
	blacklistString := os.Getenv("BLACKLIST_BOT_MODULES")
	for _, moduleName := range strings.Split(blacklistString, ",") {
		blacklist[moduleName] = nil
	}
	return blacklist
}
