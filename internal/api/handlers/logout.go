package handlers

import (
	"fmt"
	"net/http"
)

/*
Logout est en charge de l'url "/logout".
*/
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: AuthCookieName, Value: "", MaxAge: -1}
	http.SetCookie(w, &cookie)
	// if behind a proxy who change strip url prefix
	redirURL := r.Header.Get("X-Forwarded-Prefix")
	if redirURL == "" {
		redirURL = "/"
	}
	// set content-type so http.Redirect does not populate a body (see http.Redirect)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.Redirect(w, r, redirURL, http.StatusFound)
	fmt.Fprintln(w, "\"redirected\"") // json style
}
