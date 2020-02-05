package web

import (
	"fmt"
	"net/http"

	"github.com/fsufitch/discord-boar-bot/config"
	"github.com/fsufitch/discord-boar-bot/log"
	"github.com/gorilla/mux"
)

// BoarBotServer is the webserver of the boar bot
type BoarBotServer struct {
	WebPort config.WebPort
	Log     *log.Logger
	Router  Router
}

// Run is a blocking function that starts and serves the web API
func (s BoarBotServer) Run() error {
	serv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.WebPort),
		Handler: (*mux.Router)(s.Router),
	}

	s.Log.Infof("Starting web server on addr: %s " + serv.Addr)

	err := serv.ListenAndServe()

	s.Log.Errorf("Web server unexpectedly shut down with error: %v", err)

	return err
}
