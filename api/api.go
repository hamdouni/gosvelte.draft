package api

import (
	"app/api/sec"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const tokenCookieName = "jeton"
const tokenTimeLayout = time.RFC3339
const tokenDuration = time.Duration(24) * time.Hour

type API struct {
	sec   secure
	store storage
}

func (api API) Init(d storage, htmlDirectory string) error {
	// composant de sécurité
	var s sec.Secure
	if err := s.Init(); err != nil {
		return err
	}
	api.sec = s
	api.store = d
	api.InitRoutes(htmlDirectory)
	return nil
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
