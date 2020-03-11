package web

import (
	"encoding/json"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagialisi-bot/log"
	"github.com/fsufitch/tagialisi-bot/web/auth"
	"github.com/fsufitch/tagialisi-bot/web/usersession"
)

type whoAmIResponse struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	FullName      string `json:"fullname"`
	AvatarURL     string `json:"avatar_url"`
}

// WhoAmIHandler returns details about the current authenticated user
type WhoAmIHandler struct {
	Log            *log.Logger
	SessionStorage auth.SessionStorage
	AuthCookie     auth.CookieSupport
}

func (h WhoAmIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user *discordgo.User
	if sessionID, err := h.AuthCookie.GetSessionID(r); err != nil {
		http.Error(w, "could not get user session: "+err.Error(), http.StatusUnauthorized)
		return
	} else if token := h.SessionStorage.Get(sessionID); token == nil {
		http.Error(w, "session is invalid", http.StatusUnauthorized)
		return
	} else if session, err := usersession.NewIdentity(token.AccessToken); err != nil {
		http.Error(w, "could not initialize Discord session: "+err.Error(), http.StatusInternalServerError)
		return
	} else if user, err = session.User("@me"); err != nil {
		http.Error(w, "could not query user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := whoAmIResponse{
		ID:            user.ID,
		Username:      user.Username,
		Discriminator: user.Discriminator,
		FullName:      user.String(),
		AvatarURL:     user.AvatarURL(""),
	}

	data, _ := json.Marshal(response)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
