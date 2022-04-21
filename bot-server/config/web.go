package config

import (
	"os"
	"strconv"
)

// WebPort is the port that the web service will run on
type WebPort int

// ProvideWebPortFromEnvironment creates a WebPort from the environment, defaulting to 9999 when missing
func ProvideWebPortFromEnvironment() (WebPort, error) {
	portString, ok := os.LookupEnv("PORT")
	if !ok {
		portString = "9999"
	}
	port, err := strconv.ParseInt(portString, 0, 0)
	return WebPort(port), err
}

// WebEnabled is whether the web control UI is enabled
type WebEnabled bool

// ProvideWebEnabledFromEnvironment creates a WebEnabled from the environment
func ProvideWebEnabledFromEnvironment() (WebEnabled, error) {
	if enabledString, ok := os.LookupEnv("WEB_ENABLED"); ok {
		enabledBool, err := strconv.ParseBool(enabledString)
		return WebEnabled(enabledBool), err
	}
	return false, nil
}
