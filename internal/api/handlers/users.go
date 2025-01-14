package handlers

import (
	"net/http"

	"wtk/internal/api/jayson"
	"wtk/user"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	realm := Realm(r)
	users, err := user.List(realm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jayson.Respond(w, http.StatusOK, users)
}
