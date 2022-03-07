package web

import (
	"admin/app"
	"net/http"
)

func (api *API) Historic(w http.ResponseWriter, r *http.Request) {
	message := app.ListeHistorique()
	respondJSON(w, http.StatusOK, message)
}
