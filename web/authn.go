package web

import (
	"context"
	"net/http"
	"net/url"

	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/web/auth"
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

	oauthToken, err := (*oauth2.Config)(h.OAuth2Config).Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "failed to start Discord session with OAuth2 Code: "+err.Error(), http.StatusInternalServerError)
		return
	}
	sessionID := h.SessionStorage.New(oauthToken).ID

	u, _ := url.Parse(loginState.ReturnURL)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Set("sid", sessionID)
	u.RawQuery = q.Encode()

	http.Redirect(w, r, u.String(), http.StatusFound)
}

// LogoutHandler handles logout logic
type LogoutHandler struct {
	SessionStorage auth.SessionStorage
}

func (h LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sessionID := auth.GetSessionID(r)
	if sessionID == "" {
		http.Error(w, "request contained no session ID", http.StatusUnauthorized)
		return
	}
	session := h.SessionStorage.Get(sessionID)
	if session == nil {
		http.Error(w, "could not get session from session ID in request", http.StatusUnauthorized)
		return
	}

	h.SessionStorage.Clear(session.ID)
	// TODO: also revoke the token

	w.WriteHeader(http.StatusNoContent)
}
