package api

import (
	"encoding/base64"
	"net"
	"net/http"
	"webtoolkit/biz/secure"
	"strings"
)

const authCookieName = "jeton"

// auth enveloppe les handlers nécessitant une authentification
func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isAuth(r) {
			respond(w, http.StatusUnauthorized, "not authorized")
			return
		}
		next(w, r)
	}
}

// isAuth vérifie l'authentification basée sur le cookie
func isAuth(r *http.Request) bool {
	cookie, err := r.Cookie(authCookieName)
	if err != nil {
		return false
	}

	token, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return false
	}

	auth, err := secure.CheckToken(string(token), ipAddress(r))
	if err != nil {
		return false
	}
	return auth
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
func ipAddress(r *http.Request) string {
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
