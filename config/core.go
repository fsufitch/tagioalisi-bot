package config

import (
	"os"
	"strconv"
)

// DebugMode is whether a debug state is set
type DebugMode bool

// ProvideDebugModeFromEnvironment creates a DebugMode based on the value in the DEBUG env var
func ProvideDebugModeFromEnvironment() (DebugMode, error) {
	debugString, ok := os.LookupEnv("DEBUG")
	if !ok {
		return false, nil
	}
	mode, err := strconv.ParseBool(debugString)
	return DebugMode(mode), err
}

// DiscordLogChannel is the ID of the channel to post logs in
type DiscordLogChannel string

// ProvideDiscordLogChannelFromEnvironment creates a DiscordLogChannel from the environment
func ProvideDiscordLogChannelFromEnvironment() DiscordLogChannel {
	return DiscordLogChannel(os.Getenv("DISCORD_LOG_CHANNEL"))
}

// UserAgent is the user agent the bot should use when making external queries
type UserAgent string

// ProvideUserAgent creates a basic user agent to use
func ProvideUserAgent() UserAgent {
	return "tagioalisi-bot"
}
