package web

import (
	"net/http"
)

func (web *WEB) Historic(w http.ResponseWriter, r *http.Request) {
	message := web.biz.Historic()
	respondJSON(w, http.StatusOK, message)
}
