package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type API struct {
	biz    business
	secret secure
}

func (api API) Init(b business, s secure) {
	api.biz = b
	api.secret = s
	api.InitRoutes()
}

func (api *API) auth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("id")
		if err != nil {
			api.Login(w, r)
			return
		}
		fmt.Println(cookie.Value)
		if api.isValid(cookie.Value) {
			f(w, r)
		} else {
			api.Logout(w, r)
		}
	}
}

func (api *API) isValid(cookie string) bool {
	decoded, err := base64.StdEncoding.DecodeString(cookie)
	if err != nil {
		return false
	}
	code, err := api.secret.Decrypt(decoded)
	if err != nil {
		return false
	}
	log.Printf("Check decrypting cookie to : %s", string(code))
	return true
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
