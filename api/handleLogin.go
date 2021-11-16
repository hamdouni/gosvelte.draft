package api

import (
	"app/biz/hash"
	"log"
	"net/http"
)

// handleLogin service
func (api *API) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		respondJSON(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	if err := r.ParseForm(); err != nil {
		respondJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	user := r.Form.Get("username")
	pass := r.Form.Get("password")

	hashed := api.store.GetPasswordUser(user)
	if !hash.CheckPassword(pass, hashed) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jeton, err := api.getAuthToken(user, r)
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
