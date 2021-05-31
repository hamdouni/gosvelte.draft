package api

import (
	"encoding/base64"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func (api *API) auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(tokenCookieName)
		if err != nil {
			api.Login(w, r)
			return
		}
		if api.isAuth(cookie.Value, r) {
			next(w, r)
		} else {
			api.Logout(w, r)
		}
	}
}

func (api *API) isAuth(cookie string, r *http.Request) bool {
	decoded, err := base64.StdEncoding.DecodeString(cookie)
	if err != nil {
		return false
	}
	decoded, err = api.secret.Decrypt(decoded)
	if err != nil {
		return false
	}
	code := string(decoded)
	log.Printf("Check decrypting cookie to : %s", code)
	parts := strings.Split(code, "|")

	curTime := time.Now()
	loginTime, err := time.Parse(tokenTimeLayout, parts[2])
	if err != nil {
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
	val, err := api.secret.Encrypt([]byte(phrase))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(val), nil
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
