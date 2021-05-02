package api

import "net/http"

/*
	Lower implémente la fonction en charge de l'url "/lower".
*/
func (api *API) Lower(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	nom := r.Form.Get("nom")
	message := api.biz.Min(nom)
	respondJSON(w, http.StatusOK, message)
}
