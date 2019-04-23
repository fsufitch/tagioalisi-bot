package web

import (
	"fmt"
	"net/http"

	"github.com/fsufitch/discord-boar-bot/common"
	"github.com/gorilla/mux"
)

// BoarBotServer is the webserver of the boar bot
type BoarBotServer struct {
	Running       bool
	configuration *common.Configuration
	logger        *common.LoggerModule
	router        *mux.Router
}

// Start is a blocking function that starts and serves the web API
func (s BoarBotServer) Start() error {
	serv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.configuration.WebPort),
		Handler: s.router,
	}

	s.logger.Info("Starting web server on addr: " + serv.Addr)

	s.Running = true
	err := serv.ListenAndServe()
	s.Running = false

	s.logger.Error("Web server unexpectedly shut down")

	return err
}

// NewBoarBotServer creates a new BoarBotServer
func NewBoarBotServer(
	configuration *common.Configuration,
	logger *common.LoggerModule,
	router *mux.Router,
) *BoarBotServer {
	logger.Info("Initializing web server")
	return &BoarBotServer{
		configuration: configuration,
		logger:        logger,
		router:        router,
	}
}
