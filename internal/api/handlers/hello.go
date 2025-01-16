package handlers

import (
	"net/http"

	"wtk/hello"
	"wtk/internal/api/jayson"
)

/*
Hello est en charge de l'url "/hello".

Pour être considérée comme une HandleFunc, elle doit obligatoirement accepter en paramètre une http.ResponseWriter et un pointeur sur une http.Request
*/
func Hello(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nom := r.Form.Get("nom")
	message := hello.Hello(nom)
	jayson.Respond(w, http.StatusOK, message)
}
