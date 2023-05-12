package api

import (
	"net/http"
	"strings"
	"webtoolkit/metier/user"
)

func handleListUsers(w http.ResponseWriter, r *http.Request) {
	realm := strings.Split(r.Host, ".")[0]
	users, err := user.List(realm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, http.StatusOK, users)
}
