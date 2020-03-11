package web

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/fsufitch/tagialisi-bot/config"
	"github.com/fsufitch/tagialisi-bot/oauth"
	"golang.org/x/oauth2"
)

// LoginHandler starts the authentication workflow
type LoginHandler struct {
	OAuth2Config config.OAuth2Config
	LoginStates  oauth.States
	JWTSecret    config.JWTHMACSecret
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

// TagiJWTCookieName is the cookie name to store the JWT
const TagiJWTCookieName = "tagi-jwt"

// AuthCodeHandler handles the redirected successful OAuth2 login
type AuthCodeHandler struct {
	OAuth2Config   config.OAuth2Config
	LoginStates    oauth.States
	SessionStorage oauth.SessionStorage
	JWTSecret      config.JWTHMACSecret
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
		http.Error(w, "failed to retrieve Oauth2 Token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	sessionID := h.SessionStorage.Set(oauthToken)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		Id: sessionID,
	})

	jwtString, err := jwtToken.SignedString([]byte(h.JWTSecret))
	if err != nil {
		http.Error(w, "failed signing JWT: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     TagiJWTCookieName,
		Value:    jwtString,
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(w, r, loginState.ReturnURL, http.StatusFound)
}

// LogoutHandler handles logout logic
type LogoutHandler struct {
	SessionStorage oauth.SessionStorage
	JWTSecret      config.JWTHMACSecret
}

func (h LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		cookie *http.Cookie
		err    error
	)
	if cookie, err = r.Cookie(TagiJWTCookieName); err != nil {
		http.Error(w, "could not recover auth cookie: "+err.Error(), http.StatusBadRequest)
		return
	}
	jwtToken, err := jwt.Parse(cookie.Value, func(*jwt.Token) (interface{}, error) {
		return []byte(h.JWTSecret), nil
	})
	if err != nil {
		http.Error(w, "could not parse JWT token: "+err.Error(), http.StatusForbidden)
		return
	}

	var sessionID string
	switch c := jwtToken.Claims.(type) {
	case jwt.StandardClaims:
		sessionID = c.Id
	case jwt.MapClaims:
		sessionID = c["jti"].(string)
	}

	h.SessionStorage.Clear(sessionID)
	// TODO: also revoke the token

	clearedCookie := cookie
	clearedCookie.Value = ""
	http.SetCookie(w, clearedCookie)
	w.WriteHeader(http.StatusNoContent)
}
