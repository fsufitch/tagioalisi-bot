package config

import (
	"fmt"
	"os"
	"strconv"
)

// BotWebAPIPort is what it says on the tin
type BotWebAPIPort int

// ProvideBotWebAPIPortFromEnvironment is what it says on the tin
func ProvideBotWebAPIPortFromEnvironment() (BotWebAPIPort, error) {
	if port, ok := os.LookupEnv("BOT_PORT"); !ok || port == "" {
		return BotWebAPIPort(8081), nil
	} else if portInt, err := strconv.Atoi(port); err != nil {
		return 0, fmt.Errorf("invalid port in BOT_PORT: %s; %w", port, err)
	} else {
		return BotWebAPIPort(portInt), nil
	}
}
