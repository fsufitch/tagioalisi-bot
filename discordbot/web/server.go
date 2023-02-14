package web

import (
	"fmt"
	"net/http"

	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/log"
)

// TagioalisiAPIServer is the API webserver of Tagioalisi
type TagioalisiAPIServer struct {
	Port   config.BotWebAPIPort
	Log    *log.Logger
	Router Router
	// GRPCWrapperFunc grpc.WrapGRPCWebsocketFunc
}

// Run is a blocking function that starts and serves the web API
func (s TagioalisiAPIServer) Run() error {
	// handler := s.GRPCWrapperFunc(s.Router)

	serv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: s.Router,
	}

	s.Log.Infof("web: starting server on addr: %s ", serv.Addr)

	err := serv.ListenAndServe()

	s.Log.Errorf("web: server unexpectedly shut down with error: %v", err)

	return err
}
