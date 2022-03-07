package web

import (
	"admin/app"
	"net/http"
)

func Historic(w http.ResponseWriter, r *http.Request) {
	message := app.ListeHistorique()
	respondJSON(w, http.StatusOK, message)
}
