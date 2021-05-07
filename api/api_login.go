package api

import "net/http"

func (api *API) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	us := r.Form.Get("utilisateur")
	pw := r.Form.Get("pass")
	if api.biz.Login(us, pw) {
		respondJSON(w, http.StatusAccepted, "Autorisé")
		return
	}
	respondJSON(w, http.StatusUnauthorized, "Non autorisé")
}
