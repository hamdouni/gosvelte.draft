package api

import (
	"encoding/base64"
	"net/http"
	"wtk/biz/credential"
)

const authCookieName = "jeton"

// auth enveloppe les handlers nécessitant une authentification
func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isAuth(r) {
			respond(w, http.StatusUnauthorized, "not authorized")
			return
		}
		next(w, r)
	}
}

// isAuth vérifie l'authentification basée sur le cookie
func isAuth(r *http.Request) bool {
	cookie, err := r.Cookie(authCookieName)
	if err != nil {
		return false
	}

	token, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return false
	}

	auth, err := credential.CheckToken(string(token), ipAddress(r))
	if err != nil {
		return false
	}
	return auth
}
