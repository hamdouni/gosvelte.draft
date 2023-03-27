package api

import (
	"admin/model"
	"net/http"
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
	message := model.Upper(nom)
	model.StockHistorique(message)
	respond(w, http.StatusOK, message)
}
