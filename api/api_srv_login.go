package api

import (
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

	jeton, err := api.getToken(user, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Login error : %v", err)
		return
	}

	cookie := http.Cookie{
		Name:     cookieTokenName,
		Value:    string(jeton),
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
