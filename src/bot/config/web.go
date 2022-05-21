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
		portString = "80"
	}
	port, err := strconv.ParseInt(portString, 0, 0)
	return WebPort(port), err
}
