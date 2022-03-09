package api

import (
	"admin/model"
	"encoding/base64"
	"net"
	"net/http"
	"strings"
)

func handleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(tokenCookieName)
		if err != nil {
			respondJSON(w, http.StatusUnauthorized, "Non autorisé (pas de cookie) : "+err.Error())
			return
		}
		if !isAuth(cookie.Value, r) {
			respondJSON(w, http.StatusUnauthorized, "Non autorisé.")
			return
		}
		next(w, r)
	}
}

func isAuth(cookie string, r *http.Request) bool {
	token, err := base64.StdEncoding.DecodeString(cookie)
	if err != nil {
		return false
	}

	auth, err := model.CheckToken(string(token), getIPAddress(r))
	if err != nil {
		return false
	}
	return auth
}

func getAuthToken(user string, r *http.Request) (token string, err error) {
	val, err := model.NewToken(user, getIPAddress(r))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString([]byte(val)), nil
}

// Pour récupérer la vraie IP du user, il faut tenir compte du fait que l'ip
// au niveau de la connexion réseau sera faussée par l'utilisation d'un
// reverse-proxy comme Traefik. Heureusement, ces derniers surchargent le
// header http avec des entrées comme X-Fowarded-For et/ou X-Real-IP pour garder
// trace de l'IP réelle.
// Donc le principe c'est de vérifier si on trouve l'une de ces entrées,
// en prenant la dernière utilisée (c'est censé être l'IP avant de rentrer dans
// notre reverse-proxy)
// Et si on n'en trouve pas, on retourne l'adresse réseau.
func getIPAddress(r *http.Request) string {
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		addr := strings.Split(r.Header.Get(h), ",")
		for i := len(addr) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(addr[i])
			realIP := net.ParseIP(ip)
			if !realIP.IsGlobalUnicast() {
				continue
			}
			return ip
		}
	}
	// par défaut, on renvoi l'IP réseau après avoir enlever le port
	ip := strings.Split(r.RemoteAddr, ":")
	return ip[0]
}
