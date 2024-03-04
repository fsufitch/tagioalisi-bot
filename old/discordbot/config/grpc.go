package config

import (
	"fmt"
	"os"
	"strconv"
)

// BotWebAPIPort is what it says on the tin
type GRPCPort int

func ProvideGRPCPortFromEnvironment() (GRPCPort, error) {
	if port, ok := os.LookupEnv("DISCORDBOT_GRPC_PORT"); !ok || port == "" {
		return GRPCPort(9000), nil
	} else if portInt, err := strconv.Atoi(port); err != nil {
		return 0, fmt.Errorf("invalid port in DISCORDBOT_GRPC_PORT: %s; %w", port, err)
	} else {
		return GRPCPort(portInt), nil
	}
}
