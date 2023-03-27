package api

import (
	"admin/model"
	"net/http"
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
	message := model.Lower(nom)
	model.StockHistorique(message)
	respond(w, http.StatusOK, message)
}
