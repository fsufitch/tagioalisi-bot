package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fsufitch/discord-boar-bot/bot/sockpuppet-module"
)

// SockpuppetHandler is a http.Handler that sends messages through the bot
type SockpuppetHandler struct {
	BotModule *sockpuppet.Module
}

// SockpuppetPayload is the incoming payload from a sockpuppet request
type SockpuppetPayload struct {
	ChannelID string `json:"channelID"`
	Message   string `json:"message"`
}

// ServeHTTP passes the message to Discord
func (h SockpuppetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	payload := SockpuppetPayload{}
	json.Unmarshal(body, &payload)

	err := h.BotModule.SendMessage(payload.ChannelID, payload.Message)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
