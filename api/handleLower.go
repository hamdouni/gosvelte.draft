package api

import (
	"net/http"
	"webtoolkit/metier"
	"webtoolkit/metier/historic"
)

/*
	handleLower est en charge de l'url "/lower".
*/
func handleLower(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nom := r.Form.Get("nom")
	message := metier.Lower(nom)
	historic.Save(message)
	respond(w, http.StatusOK, message)
}
