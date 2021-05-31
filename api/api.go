package api

import (
	"encoding/json"
	"net/http"
	"time"
)

const tokenCookieName = "jeton"
const tokenTimeLayout = time.RFC3339
const tokenDuration = time.Duration(10) * time.Second

type API struct {
	biz    business
	secret secure
}

func (api API) Init(b business, s secure) {
	api.biz = b
	api.secret = s
	api.InitRoutes()
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
