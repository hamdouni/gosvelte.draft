package api

import "net/http"

// Logout clear id cookie
func (api *API) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: tokenCookieName, Value: "", MaxAge: -1}
	http.SetCookie(w, &cookie)
	// if behind a proxy who change strip url prefix
	redirURL := r.Header.Get("X-Forwarded-Prefix")
	if redirURL == "" {
		redirURL = "/"
	}
	http.Redirect(w, r, redirURL, http.StatusFound)
}
