package api

import (
	"encoding/base64"
	"log"
	"net/http"
)

// Login service
func (api *API) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	user := r.Form.Get("username")
	pass := r.Form.Get("password")

	if !api.biz.CheckPassword(user, pass) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	val, err := api.secret.Encrypt([]byte(user))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Login error : %v", err)
		return
	}
	encoded := base64.StdEncoding.EncodeToString(val)
	cookie := http.Cookie{Name: "id", Value: string(encoded)}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
