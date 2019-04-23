package web

import (
	"github.com/google/wire"
)

// WebProviderSet contains all the necessary wire providers to stand up a webserver
var WebProviderSet = wire.NewSet(
	NewBoarBotServer,
	NewSecretBearerAuthorizationWrapper,
	NewRouter,
	NewHelloHandler,
	NewSockpuppetHandler,
)
