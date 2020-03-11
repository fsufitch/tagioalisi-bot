package auth

import "net/http"

const AuthCookieName = "tagi-jwt-auth"

type CookieSupport struct {
	JWT JWTSupport
}

func (c CookieSupport) GetSessionID(r *http.Request) (sessionID string, err error) {
	var cookie *http.Cookie
	if cookie, err = r.Cookie(AuthCookieName); err != nil {
		return
	}
	return c.JWT.DecodeSessionID(cookie.Value)
}

func (c CookieSupport) SetSessionID(w http.ResponseWriter, sessionID string) error {
	jwtString, err := c.JWT.EncodeSessionID(sessionID)
	if err != nil {
		return err
	}
	cookie := http.Cookie{
		Name:     AuthCookieName,
		Value:    jwtString,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	return nil
}
