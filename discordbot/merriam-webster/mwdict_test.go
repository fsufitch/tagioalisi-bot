package mwdict

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicClientErrorMessagesNoAPIKey(t *testing.T) {
	// Setup
	bc := NewBasicClient("", "testuseragent")
	target_error := errors.New("no merriam-webster api key found")

	// Tested code
	_, _, err := bc.SearchCollegiate("testing")

	// Asserts
	assert.Equal(t, target_error, err)
}

func TestBasicClientErrorMessagesKeyProvided(t *testing.T) {
	// Setup
	bc := NewBasicClient("api_key", "testuseragent")

	// Tested code
	_, _, err := bc.SearchCollegiate("testing")

	// Asserts
	assert.Equal(t, nil, err)
}
