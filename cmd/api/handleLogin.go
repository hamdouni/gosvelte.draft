package api

import (
	"encoding/base64"
	"net/http"
	"wtk/biz"
)

/*
handleLogin est en charge de l'url "/login".
*/
func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		respond(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	if err := r.ParseForm(); err != nil {
		respond(w, http.StatusBadRequest, err.Error())
		return
	}
	realm := getRealm(r)
	user := r.Form.Get("username")
	pass := r.Form.Get("password")
	address := ipAddress(r)

	token, err := biz.Auth(realm, user, pass, address)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	cookie := http.Cookie{
		Name:     authCookieName,
		Value:    base64.StdEncoding.EncodeToString([]byte(token)),
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
