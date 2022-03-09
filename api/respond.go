package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// respond retourne une reponse json avec le statut et le contenu passés en paramètre
func respond(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err = w.Write([]byte(response)); err != nil {
		log.Printf("respondJSON error %s", err)
	}
}
