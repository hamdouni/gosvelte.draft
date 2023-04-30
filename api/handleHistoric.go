package api

import (
	"net/http"
	"webtoolkit/metier/historic"
)

/*
	handleHistoric est en charge de l'url "/historic".
*/
func handleHistoric(w http.ResponseWriter, r *http.Request) {
	message := historic.Liste()
	respond(w, http.StatusOK, message)
}
