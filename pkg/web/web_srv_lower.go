package web

import (
	"app/usecase"
	"log"
	"net/http"
)

/*
	Lower implémente la fonction en charge de l'url "/lower".
*/
func (web *WEB) Lower(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nom := r.Form.Get("nom")
	log.Printf("got nom: %s", nom)
	message := usecase.Lower(nom)
	web.data.StockHistorique(message)
	respondJSON(w, http.StatusOK, message)
}
