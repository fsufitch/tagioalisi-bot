package web

import (
	"context"
	"net/http"

	"github.com/fsufitch/tagialisi-bot/config"
	"github.com/fsufitch/tagialisi-bot/web/auth"
	"golang.org/x/oauth2"
)

// LoginHandler starts the authentication workflow
type LoginHandler struct {
	OAuth2Config config.OAuth2Config
	LoginStates  auth.LoginStates
}

func (h LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	returnURL := r.URL.Query().Get("return_url")
	if returnURL == "" {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no return_url specified"))
		return
	}

	loginState := h.LoginStates.New(returnURL)

	redirectURL := (*oauth2.Config)(h.OAuth2Config).AuthCodeURL(loginState.ID)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

// AuthCodeHandler handles the redirected successful OAuth2 login
type AuthCodeHandler struct {
	OAuth2Config   config.OAuth2Config
	LoginStates    auth.LoginStates
	SessionStorage auth.SessionStorage
	AuthCookie     auth.CookieSupport
}

func (h AuthCodeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if code == "" || state == "" {
		http.Error(w, "both code and state are required", http.StatusBadRequest)
		return
	}

	loginState := h.LoginStates.Get(state)
	if loginState == nil {
		http.Error(w, "expired/invalid state; go back and try again", http.StatusForbidden)
		return
	}
	h.LoginStates.Clear(state)

	//session, err := discordgo.New(code)
	oauthToken, err := (*oauth2.Config)(h.OAuth2Config).Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "failed to start Discord session with OAuth2 Code: "+err.Error(), http.StatusInternalServerError)
		return
	}
	sessionID := h.SessionStorage.Set(oauthToken)

	if err := h.AuthCookie.SetSessionID(w, sessionID); err != nil {
		http.Error(w, "failed encoding/setting JWT cookie: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, loginState.ReturnURL, http.StatusFound)
}

// LogoutHandler handles logout logic
type LogoutHandler struct {
	SessionStorage auth.SessionStorage
	AuthCookie     auth.CookieSupport
}

func (h LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		cookie *http.Cookie
		err    error
	)
	sessionID, err := h.AuthCookie.GetSessionID(r)
	if err != nil {
		http.Error(w, "could not get JWT from cookie: "+err.Error(), http.StatusForbidden)
		return
	}

	h.SessionStorage.Clear(sessionID)
	// TODO: also revoke the token

	clearedCookie := cookie
	clearedCookie.Value = ""
	http.SetCookie(w, clearedCookie)
	w.WriteHeader(http.StatusNoContent)
}
