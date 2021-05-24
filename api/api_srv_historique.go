package api

import (
	"net/http"
)

func (api *API) Historic(w http.ResponseWriter, r *http.Request) {
	message := api.biz.Historic()
	respondJSON(w, http.StatusOK, message)
}
