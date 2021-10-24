package web

import "net/http"

/*
	Upper implémente la fonction en charge de l'url "/upper".

	Pour être considérée comme une HandleFunc, elle doit obligatoirement accepter en paramètre une http.ResponseWriter et un pointeur sur une http.Request
*/
func (web *WEB) Upper(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nom := r.Form.Get("nom")
	message := web.biz.Maj(nom)
	web.data.StockHistorique(message)
	respondJSON(w, http.StatusOK, message)
}
