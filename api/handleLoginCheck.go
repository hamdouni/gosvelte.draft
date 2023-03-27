package api

import "net/http"

// handleLoginCheck check the login connexion
// If a cookie exists and is valid return OK status.
// Return Forbidden if not.
func handleLoginCheck(w http.ResponseWriter, r *http.Request) {
	if !isAuth(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}
