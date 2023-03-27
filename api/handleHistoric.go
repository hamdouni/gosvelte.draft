package api

import (
	"admin/model"
	"net/http"
)

/*
	handleHistoric est en charge de l'url "/historic".
*/
func handleHistoric(w http.ResponseWriter, r *http.Request) {
	message := model.ListeHistorique()
	respond(w, http.StatusOK, message)
}
