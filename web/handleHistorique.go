package web

import (
	"admin/model"
	"net/http"
)

func Historic(w http.ResponseWriter, r *http.Request) {
	message := model.ListeHistorique()
	respondJSON(w, http.StatusOK, message)
}
