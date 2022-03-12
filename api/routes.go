package api

import "net/http"

/*
On définit les routes dans un tableau pour faciliter leur déclaration, en
précisant leurs caractéristiques :
- `endpoint` pour l'url d'accès à la route,
- `handler` pour la fonction en charge de cette route
- `auth` pour indiquer si la fonction doit être authentifiée ou non
*/
var routes = []struct {
	endpoint string           // url de la route
	handler  http.HandlerFunc // fonction en charge
	auth     bool             // indicateur fonction authentifiée
}{
	{"/hello", handleHello, false},
	{"/login", handleLogin, false},
	{"/logout", handleLogout, false},
	{"/check", handleLogCheck, false},
	{"/upper", Upper, true},
	{"/lower", Lower, true},
	{"/historic", Historic, true},
}

// Routes initalise les routes
func Routes() {
	for _, route := range routes {
		if route.auth {
			http.HandleFunc(route.endpoint, auth(route.handler))
		} else {
			http.HandleFunc(route.endpoint, route.handler)
		}
	}
}
