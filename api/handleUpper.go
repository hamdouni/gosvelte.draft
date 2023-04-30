package api

import (
	"net/http"
	"webtoolkit/metier"
	"webtoolkit/metier/historic"
)

/*
	handleUpper est en charge de l'url "/upper".
*/
func handleUpper(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nom := r.Form.Get("nom")
	message := metier.Upper(nom)
	historic.Save(message)
	respond(w, http.StatusOK, message)
}
