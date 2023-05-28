package api

import (
	"encoding/base64"
	"net/http"
	"wtk/biz/user"
)

/*
handleLogin est en charge de l'url "/login".
*/
func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		respond(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	realm := getRealm(r)
	address := ipAddress(r)

	var credential struct {
		Username string
		Password string
	}

	if err := decodeJson(r, &credential); err != nil {
		respond(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := user.Auth(realm, credential.Username, credential.Password, address)
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
