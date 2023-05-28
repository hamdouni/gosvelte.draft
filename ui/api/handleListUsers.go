package api

import (
	"net/http"
	"wtk/biz/user"
)

func handleListUsers(w http.ResponseWriter, r *http.Request) {
	realm := getRealm(r)
	users, err := user.List(realm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, http.StatusOK, users)
}
