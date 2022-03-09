package api

import (
	"net/http"
)

// handleLogCheck check the login connexion
// If a cookie exists and is valid return OK status.
// Return Forbidden if not.
func handleLogCheck(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(tokenCookieName)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if !isAuth(c.Value, r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}
