package api

import (
	"net/http"

	"wtk/internal/api/handlers"
)

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
	{"/login", handlers.Login, false},
	{"/logout", handlers.Logout, false},
	{"/hello", handlers.Hello, true},
	{"/list", handlers.ListUsers, true},
	{"/check", handlers.Check, false},
}

// Routes initalise les routes
func Routes(static string) {
	// Pour les traitements dynamiques, on déclare les
	// routes en gérant le fait qu'elles doivent être
	// authentifiées ou non.
	for _, route := range routes {
		if route.auth {
			http.HandleFunc(route.endpoint, handlers.Auth(route.handler))
		} else {
			http.HandleFunc(route.endpoint, route.handler)
		}
	}

	// Pour les fichiers statiques (html, js, images, ...), la librairie standard
	// propose une fonction FileServer.
	fs := http.FileServer(http.Dir(static))
	http.Handle("/", fs)
}
