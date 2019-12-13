package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// RunMode is an enum representing the mode in which the binary is running
type RunMode int

// Values for RunMode
const (
	Unknown RunMode = iota
	Bot
	Migration
)

// Configuration is a container for start-of-process runtime configuration values
type Configuration struct {
	RunMode             RunMode
	WebEnabled          bool
	WebPort             int
	WebSecret           string
	DiscordToken        string
	DatabaseURL         string
	CLILogLevel         LogLevel
	DiscordLogLevel     LogLevel
	DiscordLogChannel   string
	BlacklistBotModules map[string]bool
	MigrationDir        string
}

// ConfigurationFromEnvironment bootstraps a configuration object based on environment variables
func ConfigurationFromEnvironment() (*Configuration, error) {
	c := Configuration{}

	switch strings.ToLower(os.Getenv("RUN_MODE")) {
	case "", "bot":
		c.RunMode = Bot
	case "migration":
		c.RunMode = Migration
	default:
		c.RunMode = Unknown
	}

	if webEnabled, err := strconv.ParseBool(os.Getenv("WEB_ENABLED")); err == nil {
		c.WebEnabled = webEnabled
	} else {
		return nil, errors.Wrap(err, "invalid or missing value for WEB_ENABLED")
	}

	if c.WebEnabled {
		if webPort, err := strconv.ParseInt(os.Getenv("PORT"), 10, 64); err == nil {
			c.WebPort = int(webPort)
		} else {
			return nil, errors.Wrap(err, "invalid or missing value for PORT")
		}

		if webSecret, ok := os.LookupEnv("WEB_SECRET"); ok && webSecret != "" {
			c.WebSecret = webSecret
		} else {
			return nil, errors.New("missing value for WEB_SECRET")
		}
	}

	if token, ok := os.LookupEnv("DISCORD_TOKEN"); ok {
		c.DiscordToken = token
	} else {
		return nil, errors.New("missing value for DISCORD_TOKEN")
	}

	if dbURL, ok := os.LookupEnv("DATABASE_URL"); ok {
		c.DatabaseURL = dbURL
	} else {
		return nil, errors.New("missing value for DATABASE_URL")
	}

	c.CLILogLevel = LogInfo
	if logLevelString, ok := os.LookupEnv("LOG_LEVEL"); ok {
		switch logLevelString {
		case "debug":
			c.CLILogLevel = LogDebug
		case "info":
			c.CLILogLevel = LogInfo
		case "warn", "warning":
			c.CLILogLevel = LogWarning
		case "error":
			c.CLILogLevel = LogError
		default:
			return nil, fmt.Errorf("invalid value for LOG_LEVEL: %s", logLevelString)
		}
	}

	c.DiscordLogLevel = LogInfo
	if logLevelString, ok := os.LookupEnv("DISCORD_LOG_LEVEL"); ok {
		switch logLevelString {
		case "debug":
			c.DiscordLogLevel = LogDebug
		case "info":
			c.DiscordLogLevel = LogInfo
		case "warn", "warning":
			c.DiscordLogLevel = LogWarning
		case "error":
			c.DiscordLogLevel = LogError
		default:
			return nil, fmt.Errorf("invalid value for DISCORD_LOG_LEVEL: %s", logLevelString)
		}
	}

	c.DiscordLogChannel = os.Getenv("DISCORD_LOG_CHANNEL")

	blacklistString := os.Getenv("BLACKLIST_BOT_MODULES")
	c.BlacklistBotModules = map[string]bool{}
	for _, moduleName := range strings.Split(blacklistString, ",") {
		c.BlacklistBotModules[moduleName] = true
	}

	c.MigrationDir = os.Getenv("MIGRATION_DIR")

	return &c, nil
}
