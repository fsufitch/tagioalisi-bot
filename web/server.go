package web

import (
	"fmt"
	"net/http"

	"github.com/fsufitch/discord-boar-bot/common"
)

// BoarBotServer is the webserver of the boar bot
type BoarBotServer struct {
	Running       bool
	configuration *common.Configuration
	logger        *common.LoggerModule
}

// Start is a blocking function that starts and serves the web API
func (s BoarBotServer) Start() error {
	serv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.configuration.WebPort),
		Handler: HelloWorldHandler{},
	}

	s.logger.Info("Starting web server on addr: " + serv.Addr)

	s.Running = true
	err := serv.ListenAndServe()
	s.Running = false

	s.logger.Error("Web server unexpectedly shut down")

	return err
}

// HelloWorldHandler just says hello world to everything
type HelloWorldHandler struct{}

func (h HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world"))
}

// NewBoarBotServer creates a new BoarBotServer
func NewBoarBotServer(configuration *common.Configuration, logger *common.LoggerModule) *BoarBotServer {
	logger.Info("Initializing web server")
	return &BoarBotServer{
		configuration: configuration,
		logger:        logger,
	}
}
