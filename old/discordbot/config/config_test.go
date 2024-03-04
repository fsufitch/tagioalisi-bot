package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseString(t *testing.T) {
	// Setup
	os.Setenv("DATABASE", "test")

	// Tested code
	db, err := ProvideDatabaseStringFromEnvironment()

	// Asserts
	assert.Nil(t, err)
	assert.Equal(t, "test", string(db))
}

func TestDatabaseString_Missing(t *testing.T) {
	// Setup
	os.Unsetenv("DATABASE")

	// Tested code
	_, err := ProvideDatabaseStringFromEnvironment()

	// Asserts
	assert.NotNil(t, err)
}

func TestWebPort(t *testing.T) {
	// Setup
	os.Setenv("DISCORDBOT_HTTPS_PORT", "1234")

	// Tested code
	port, err := ProvideBotHTTPSPortFromEnvironment()

	// Asserts
	assert.Nil(t, err)
	assert.Equal(t, 1234, int(port))
}

func TestWebPort_Default(t *testing.T) {
	// Setup
	os.Unsetenv("DISCORDBOT_HTTPS_PORT")

	// Tested code
	port, err := ProvideBotHTTPSPortFromEnvironment()

	// Asserts
	assert.Nil(t, err)
	assert.Equal(t, 7443, int(port))
}

func TestWebPort_Invalid(t *testing.T) {
	// Setup
	os.Setenv("DISCORDBOT_HTTPS_PORT", "Not a real port")

	// Tested code
	_, err := ProvideBotHTTPSPortFromEnvironment()

	// Asserts
	assert.NotNil(t, err)
}

func TestDebugMode(t *testing.T) {
	// Setup
	os.Setenv("DEBUG", "1")

	// Tested code
	debug, err := ProvideDebugModeFromEnvironment()

	// Asserts
	assert.Nil(t, err)
	assert.True(t, bool(debug))
}

func TestDebugMode_Default(t *testing.T) {
	// Setup
	os.Unsetenv("DEBUG")

	// Tested code
	debug, err := ProvideDebugModeFromEnvironment()

	// Asserts
	assert.Nil(t, err)
	assert.False(t, bool(debug))
}

func TestDebugMode_Invalid(t *testing.T) {
	// Setup
	os.Setenv("DEBUG", "not a real bool")

	// Tested code
	_, err := ProvideDebugModeFromEnvironment()

	// Asserts
	assert.NotNil(t, err)
}
