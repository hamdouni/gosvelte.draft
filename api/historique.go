package api

import (
	"admin/model"
	"net/http"
)

func Historic(w http.ResponseWriter, r *http.Request) {
	message := model.ListeHistorique()
	respond(w, http.StatusOK, message)
}
