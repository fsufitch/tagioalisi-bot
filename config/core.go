package config

import (
	"fmt"
	"os"
	"strconv"
)

// DebugMode is whether a debug state is set
type DebugMode bool

// ProvideDebugModeFromEnvironment creates a DebugMode based on the value in the DEBUG env var
func ProvideDebugModeFromEnvironment() (DebugMode, error) {
	debugString, ok := os.LookupEnv("DEBUG")
	fmt.Println(debugString)
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
