package web

import (
	"net/http"
)

func (web *WEB) Historic(w http.ResponseWriter, r *http.Request) {
	message := web.store.ListeHistorique()
	respondJSON(w, http.StatusOK, message)
}
