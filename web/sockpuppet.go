package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fsufitch/tagioalisi-bot/bot/sockpuppet-module"
	"github.com/fsufitch/tagioalisi-bot/log"
	"github.com/fsufitch/tagioalisi-bot/web/auth"
	"github.com/fsufitch/tagioalisi-bot/web/usersession"
	"github.com/pkg/errors"
)

// SockpuppetHandler is a http.Handler that sends messages through the bot
type SockpuppetHandler struct {
	BotModule *sockpuppet.Module
	Log       *log.Logger
	Sessions  auth.SessionStorage
}

// SockpuppetPayload is the incoming payload from a sockpuppet request
type SockpuppetPayload struct {
	ChannelID string `json:"channelID"`
	Message   string `json:"message"`
}

// ServeHTTP passes the message to Discord
func (h SockpuppetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var userID string
	if session := h.Sessions.Get(auth.GetSessionID(r)); session == nil {
		http.Error(w, "Not authenticated", http.StatusUnauthorized)
		return
	} else if identity, err := usersession.NewIdentity(session.OAuth2Token.AccessToken); err != nil {
		http.Error(w, "Failed to initialize identity client: "+err.Error(), http.StatusUnauthorized)
		return
	} else if user, err := identity.User("@me"); err != nil {
		http.Error(w, "Failed to get @me: "+err.Error(), http.StatusUnauthorized)
	} else {
		userID = user.ID
	}

	body, _ := ioutil.ReadAll(r.Body)
	payload := SockpuppetPayload{}
	json.Unmarshal(body, &payload)

	err := h.BotModule.SendMessage(payload.ChannelID, payload.Message, userID)

	if errors.Is(err, sockpuppet.ErrSendingNotPermitted) {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	if err != nil {
		http.Error(w, "Unknown error sending message: "+err.Error(), http.StatusInternalServerError)
		h.Log.Errorf("sockpuppet-web: error sending message: %v", err)
		h.Log.HTTP(http.StatusBadRequest, r)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNoContent)
	h.Log.HTTP(http.StatusNoContent, r)
}
