package web

import (
	"encoding/json"
	"net/http"

	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/log"
)

// HelloHandler is a http.Handler just says hello world
type HelloHandler struct {
	Log *log.Logger

	DebugMode          config.DebugMode
	BotModuleBlacklist config.BotModuleBlacklist
	ManagedGroupPrefix config.ManagedGroupPrefix
	OAuth2Config       config.OAuth2Config
}

// HelloResponse contains the JSON payload of the hello response
type HelloResponse struct {
	DebugMode          bool     `json:"debug_mode"`
	DiscordClientID    string   `json:"discord_client_id"`
	BotModuleBlacklist []string `json:"bot_module_blacklist"`
	GroupPrefix        string   `json:"group_prefix"`
}

// ServeHTTP just says hello world
func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		DebugMode:       bool(h.DebugMode),
		DiscordClientID: h.OAuth2Config.ClientID,
		GroupPrefix:     string(h.ManagedGroupPrefix),
	}

	botModuleBlacklist := []string{}
	for module := range h.BotModuleBlacklist {
		botModuleBlacklist = append(botModuleBlacklist, module)
	}
	response.BotModuleBlacklist = botModuleBlacklist

	data, _ := json.Marshal(response)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	h.Log.HTTP(http.StatusOK, r)
	w.Write(data)
}
