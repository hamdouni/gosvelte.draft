package handlers

import (
	"encoding/base64"
	"net/http"

	"wtk/internal/api/jayson"
	"wtk/user"
)

const AuthCookieName = "jeton"

// Auth enveloppe les handlers nécessitant une authentification
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// récupère le jeton
		cookie, err := r.Cookie(AuthCookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := base64.StdEncoding.DecodeString(cookie.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !user.IsAuth(string(token), IPAddress(r)) {
			jayson.Respond(w, http.StatusUnauthorized, "not authorized")
			return
		}
		next(w, r)
	}
}
