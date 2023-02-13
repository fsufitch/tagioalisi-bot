package web

import (
	"encoding/json"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagioalisi-bot/log"
	"github.com/fsufitch/tagioalisi-bot/security"
	"github.com/fsufitch/tagioalisi-bot/web/usersession"
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
	Log *log.Logger
	JWT security.JWTSupport
}

func (h WhoAmIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user *discordgo.User
	if jwt, err := h.JWT.ExtractJWTFromRequest(r); err != nil {
		switch err {
		case security.ErrJWTExpired:
			http.Error(w, "token expired", http.StatusUnauthorized)
		case security.ErrNoJWTFound:
			http.Error(w, "unauthorized", http.StatusUnauthorized)
		default:
			http.Error(w, "unknown error authorizing: "+err.Error(), http.StatusInternalServerError)
		}
		return
	} else if discordSession, err := usersession.NewIdentity(jwt.AccessToken); err != nil {
		http.Error(w, "could not initialize Discord session: "+err.Error(), http.StatusInternalServerError)
		return
	} else if user, err = discordSession.User("@me"); err != nil {
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
