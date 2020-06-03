package web

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/security"
	"github.com/fsufitch/tagioalisi-bot/web/auth"
	"github.com/fsufitch/tagioalisi-bot/web/usersession"
	"golang.org/x/oauth2"
)

// LoginHandler starts the authentication workflow
type LoginHandler struct {
	OAuth2Config config.OAuth2Config
	AES          security.AESSupport
}

func (h LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	returnURL := r.URL.Query().Get("return_url")
	if returnURL == "" {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no return_url specified"))
		return
	}

	loginState := auth.NewLoginState(returnURL)
	stateStr, err := loginState.ToStateParam(h.AES)
	if err != nil {
		http.Error(w, fmt.Sprintf("unexpected error generating state: %v", err), http.StatusInternalServerError)
		return
	}

	redirectURL := (*oauth2.Config)(h.OAuth2Config).AuthCodeURL(stateStr)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

// AuthCodeHandler handles the redirected successful OAuth2 login
type AuthCodeHandler struct {
	OAuth2Config config.OAuth2Config
	JWT          security.JWTSupport
	AES          security.AESSupport
}

func (h AuthCodeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if code == "" || state == "" {
		http.Error(w, "both code and state are required", http.StatusBadRequest)
		return
	}

	// Unpack login state from the state param
	loginState, err := auth.LoginStateFromStateParam(state, h.AES)
	if err != nil {
		http.Error(w, "error unpacking state: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Acquire actual OAuth access token (and more) for this login
	oauthToken, err := (*oauth2.Config)(h.OAuth2Config).Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "failed to start Discord session with OAuth2 Code: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var user *discordgo.User
	// Query identity information
	if discordSession, err := usersession.NewIdentity(oauthToken.AccessToken); err != nil {
		http.Error(w, "could not initialize Discord session: "+err.Error(), http.StatusInternalServerError)
		return
	} else if user, err = discordSession.User("@me"); err != nil {
		http.Error(w, "could not query user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	jwtToken, err := h.JWT.CreateTokenString(security.JWTData{
		AccessToken: oauthToken.AccessToken,
		Expiration:  oauthToken.Expiry,
		SessionID:   loginState.ID,
		UserID:      user.ID,
		Username:    user.String(),
		AvatarURL:   user.AvatarURL(""),
	})

	if err != nil {
		http.Error(w, "error assembling jwt: "+err.Error(), http.StatusInternalServerError)
		return
	}

	u, _ := url.Parse(loginState.ReturnURL)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Set("jwt", jwtToken)
	u.RawQuery = q.Encode()

	http.Redirect(w, r, u.String(), http.StatusFound)
}

// LogoutHandler handles logout logic
type LogoutHandler struct{}

func (h LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// All auth is in the JWT, so nothing to do on server side
	w.WriteHeader(http.StatusNoContent)
}
