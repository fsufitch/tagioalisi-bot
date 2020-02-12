package web

import (
	"fmt"
	"net/http"

	"github.com/fsufitch/discord-boar-bot/config"
	"github.com/fsufitch/discord-boar-bot/log"
	"github.com/gorilla/mux"
)

// TagioalisiAPIServer is the API webserver of Tagioalisi
type TagioalisiAPIServer struct {
	WebPort config.WebPort
	Log     *log.Logger
	Router  Router
	CORS    *CORSWrapper
}

// Run is a blocking function that starts and serves the web API
func (s TagioalisiAPIServer) Run() error {
	serv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.WebPort),
		Handler: s.CORS.Wrap((*mux.Router)(s.Router)),
	}

	s.Log.Infof("web: starting server on addr: %s ", serv.Addr)

	err := serv.ListenAndServe()

	s.Log.Errorf("web: server unexpectedly shut down with error: %v", err)

	return err
}
