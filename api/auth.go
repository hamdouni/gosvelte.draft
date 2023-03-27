package api

import (
	"admin/model"
	"admin/model/secure"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"strings"
)

const tokenCookieName = "jeton"

// handleLogin service
func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		respond(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	if err := r.ParseForm(); err != nil {
		respond(w, http.StatusBadRequest, err.Error())
		return
	}
	user := r.Form.Get("username")
	pass := r.Form.Get("password")
	address := ipAddress(r)

	token, err := model.Auth(user, pass, address)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	cookie := http.Cookie{
		Name:     tokenCookieName,
		Value:    base64.StdEncoding.EncodeToString([]byte(token)),
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}

// handleLogout clear id cookie
func handleLogout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: tokenCookieName, Value: "", MaxAge: -1}
	http.SetCookie(w, &cookie)
	// if behind a proxy who change strip url prefix
	redirURL := r.Header.Get("X-Forwarded-Prefix")
	if redirURL == "" {
		redirURL = "/"
	}
	// set content-type so http.Redirect does not populate a body (see http.Redirect)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.Redirect(w, r, redirURL, http.StatusFound)
	fmt.Fprintln(w, "\"redirected\"") // json style
}

// handleLogCheck check the login connexion
// If a cookie exists and is valid return OK status.
// Return Forbidden if not.
func handleLogCheck(w http.ResponseWriter, r *http.Request) {
	if !isAuth(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// auth wraps other handlers to check authentification.
func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isAuth(r) {
			respond(w, http.StatusUnauthorized, "Non autorisé.")
			return
		}
		next(w, r)
	}
}

// isAuth is a helper to check authentification based on the cookie
func isAuth(r *http.Request) bool {
	cookie, err := r.Cookie(tokenCookieName)
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
