package web

import (
	"admin/app"
	"admin/model"
	"net/http"
)

/*
	Lower implémente la fonction en charge de l'url "/lower".
*/
func (api *API) Lower(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nom := r.Form.Get("nom")
	message := model.Lower(nom)
	app.StockHistorique(message)
	respondJSON(w, http.StatusOK, message)
}
