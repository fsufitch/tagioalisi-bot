package main

import (
	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/web"
)

// WebRunFunc is a function to launch the web server, or nil if web is disabled
type WebRunFunc func() error

// ProvideWebRunFunc assigns the appropriate web run function
func ProvideWebRunFunc(webEnabled config.WebEnabled, server web.TagioalisiAPIServer) WebRunFunc {
	if !webEnabled {
		return nil
	}
	return server.Run
}
