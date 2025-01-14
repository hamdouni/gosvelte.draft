package handlers

import (
	"encoding/base64"
	"net/http"

	"wtk/user"
)

// handleLoginCheck check the login connexion
// If a cookie exists and is valid return OK status.
// Return Forbidden if not.
func Check(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}
