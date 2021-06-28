package web

import (
	"encoding/json"
	"net/http"
	"time"
)

const tokenCookieName = "jeton"
const tokenTimeLayout = time.RFC3339
const tokenDuration = time.Duration(24) * time.Hour

type WEB struct {
	biz    business
	secret secure
}

func (web WEB) Init(b business, s secure) {
	web.biz = b
	web.secret = s
	web.InitRoutes()
}

// respondJSON retourne une reponse json avec le statut et le contenu passés en paramètre
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
