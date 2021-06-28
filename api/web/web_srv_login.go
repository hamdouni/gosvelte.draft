package web

import (
	"log"
	"net/http"
)

// Login service
func (web *WEB) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	user := r.Form.Get("username")
	pass := r.Form.Get("password")

	if !web.biz.CheckPassword(user, pass) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	jeton, err := web.getAuthToken(user, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Login error : %v", err)
		return
	}

	cookie := http.Cookie{
		Name:     tokenCookieName,
		Value:    string(jeton),
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
