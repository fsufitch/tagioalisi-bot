package config

import "os"

// ManagedGroupPrefix is a configurable prefix for the `groups` module
type ManagedGroupPrefix string

// ProvideManagedGroupPrefixFromEnvironment creates a ManagedGroupPrefix from the environment, or defaults to "g-"
func ProvideManagedGroupPrefixFromEnvironment() ManagedGroupPrefix {
	prefix := os.Getenv("GROUP_PREFIX")
	if prefix == "" {
		return "g-"
	}
	return ManagedGroupPrefix(prefix)
}
