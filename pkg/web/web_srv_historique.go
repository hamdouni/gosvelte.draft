package web

import (
	"net/http"
)

func (web *WEB) Historic(w http.ResponseWriter, r *http.Request) {
	message := web.data.ListeHistorique()
	respondJSON(w, http.StatusOK, message)
}
