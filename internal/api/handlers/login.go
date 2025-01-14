package handlers

import (
	"encoding/base64"
	"net/http"

	"wtk/internal/api/jayson"
	"wtk/user"
)

/*
Login est en charge de l'url "/login".
*/
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		jayson.Respond(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	realm := Realm(r)
	address := IPAddress(r)

	var creds struct {
		Username string
		Password string
	}

	if err := jayson.DecodeJson(r, &creds); err != nil {
		jayson.Respond(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := user.Auth(realm, creds.Username, creds.Password, address)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	cookie := http.Cookie{
		Name:     AuthCookieName,
		Value:    base64.StdEncoding.EncodeToString([]byte(token)),
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
