package api

import (
	"net/http"
	"strings"
)

// realm est la 1ère partie de l'url
func getRealm(r *http.Request) string {
	realm := strings.Split(r.Host, ".")[0]
	realm = strings.TrimSpace(realm)
	return realm
}
