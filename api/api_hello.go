package api

import "net/http"

/*
	Hello implémente la fonction en charge de l'url "/hello". Pour être considérée comme une HandleFunc, elle doit obligatoirement accepter en paramètre une http.ResponseWriter et un pointeur sur une http.Request
*/
func (api *API) Hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	nom := r.Form.Get("nom")
	message := api.biz.Bonjour(nom)
	respondJSON(w, http.StatusOK, message)
}
