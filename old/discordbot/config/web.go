package config

import (
	"fmt"
	"os"
	"strconv"
)

// BotHTTPSPort is what it says on the tin
type BotHTTPSPort int

func ProvideBotHTTPSPortFromEnvironment() (BotHTTPSPort, error) {
	if port, ok := os.LookupEnv("DISCORDBOT_HTTPS_PORT"); !ok || port == "" {
		return BotHTTPSPort(7443), nil
	} else if portInt, err := strconv.Atoi(port); err != nil {
		return 0, fmt.Errorf("invalid port in DISCORDBOT_HTTPS_PORT: %s; %w", port, err)
	} else {
		return BotHTTPSPort(portInt), nil
	}

}

type BotTLS struct {
	Certificate string
	SecretKey string
}

func ProvideBotTLSFromEnvironment() BotTLS {
	tls := BotTLS{
		Certificate: "/certs/discordbot.crt",
		SecretKey: "/certs/discordbot.key",
	}

	if path, ok := os.LookupEnv("DISCORDBOT_TLS_CERT"); ok {
		tls.Certificate = path
	}

	if path, ok := os.LookupEnv("DISCORDBOT_TLS_KEY"); ok {
		tls.SecretKey = path
	}

	return tls
}