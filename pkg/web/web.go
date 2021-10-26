package web

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const tokenCookieName = "jeton"
const tokenTimeLayout = time.RFC3339
const tokenDuration = time.Duration(24) * time.Hour

type WEB struct {
	sec  secure
	data store
}

func (web WEB) Init(s secure, d store, publicDirectory string) {
	web.sec = s
	web.data = d
	web.InitRoutes(publicDirectory)
}

// respondJSON retourne une reponse json avec le statut et le contenu passés en paramètre
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err = w.Write([]byte(response)); err != nil {
		log.Printf("respondJSON error %v", err)
	}
}
