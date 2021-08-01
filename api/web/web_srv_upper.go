package web

import "net/http"

/*
	Lower implémente la fonction en charge de l'url "/lower".
*/
func (web *WEB) Lower(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nom := r.Form.Get("nom")
	message := web.biz.Min(nom)
	respondJSON(w, http.StatusOK, message)
}
