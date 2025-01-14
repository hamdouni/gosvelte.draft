package handlers

import (
	"net/http"
	"strings"
)

// Realm est la 1ère partie de l'url
func Realm(r *http.Request) string {
	realm := strings.Split(r.Host, ".")[0]
	realm = strings.TrimSpace(realm)
	return realm
}
