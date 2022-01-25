package web

import (
	"encoding/base64"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func (api *API) handleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(tokenCookieName)
		if err != nil {
			respondJSON(w, http.StatusUnauthorized, "Non autorisé (pas de cookie) : "+err.Error())
			return
		}
		if !api.isAuth(cookie.Value, r) {
			respondJSON(w, http.StatusUnauthorized, "Non autorisé.")
			return
		}
		next(w, r)
	}
}

func (api *API) isAuth(cookie string, r *http.Request) bool {
	decoded, err := base64.StdEncoding.DecodeString(cookie)
	if err != nil {
		return false
	}

	code, err := api.sec.Decrypt(string(decoded))
	if err != nil {
		return false
	}
	parts := strings.Split(code, "|")

	curTime := time.Now()
	loginTime, err := time.Parse(tokenTimeLayout, parts[2])
	if err != nil {
		log.Printf("When parsing loginTime %v got error %v", parts[2], err)
		return false
	}
	dur := curTime.Sub(loginTime)
	log.Printf("Last login %v %v diff %v", curTime, loginTime, dur)
	if dur > tokenDuration {
		return false
	}

	userIP := parts[1]
	reqIP := getIPAddress(r)
	return userIP == reqIP
}

func (api *API) getAuthToken(user string, r *http.Request) (token string, err error) {
	timestamp := time.Now().Format(tokenTimeLayout)
	phrase := user + "|" + getIPAddress(r) + "|" + timestamp
	val, err := api.sec.Encrypt(phrase)
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
// notrereverse-proxy)
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
