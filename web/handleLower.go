package web

import (
	"app/lower"
	"log"
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
	log.Printf("got nom: %s", nom)
	message := lower.String(nom)
	api.store.StockHistorique(message)
	respondJSON(w, http.StatusOK, message)
}
