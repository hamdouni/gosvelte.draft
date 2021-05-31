package api

import (
	"net/http"
)

// LogCheck check the login connexion
// If a cookie name "id" exists and is valid return OK status.
// Return Forbidden if not.
func (api *API) LogCheck(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(cookieTokenName)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if !api.isValid(c.Value) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}
