package api

import (
	"net/http"
)

func (api *API) Historic(w http.ResponseWriter, r *http.Request) {
	message := api.store.ListeHistorique()
	respondJSON(w, http.StatusOK, message)
}
