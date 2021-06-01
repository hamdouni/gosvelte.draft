package api

import (
	"log"
	"net/http"
)

/*
	InitRoutes initialise toutes les routes (url) en spécifiant les fonctions qui les prennent en charge, soit des fonctions à nous qu'on déclare avec HandleFunc, soit la fonction qui s'occupe de servir les fichiers statiques.
*/
func (api *API) InitRoutes() {
	/*
		Voici des exemples de déclaration de routes avec l'url pour y accéder et la fonction qui va la traiter.
		Par exemple l'url "/hello" est prise en charge par api.Hello que l'on doit implémenter (cf fichier api_hello.go).
	*/
	api.logHandleFunc("/hello", api.Hello)
	api.logHandleFunc("/login", api.Login)
	api.logHandleFunc("/logout", api.Logout)
	api.logHandleFunc("/check", api.LogCheck)
	api.logHandleFunc("/historic", api.auth(api.Historic))
	api.logHandleFunc("/upper", api.auth(api.Upper))
	api.logHandleFunc("/lower", api.auth(api.Lower))

	/*
		Pour les fichiers statiques (html, js, images, ...), la librairie standard propose une fonction FileServer qui reçoit le dossier contenant nos fichiers statiques.
	*/
	fs := http.FileServer(http.Dir("./ihm/public"))
	http.Handle("/", fs)
}

/*
	Pour tracer les requêtes, on enrichie http.HandleFunc avec 2 fonctions :
	- logHandleFunc sera utilisé à la place de http.HandleFunc dans la déclaration des routes
	- logReq affiche la requête et continue le traitement sur le handler d'origine
*/
func (api *API) logHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, logReq(handler))
}
func logReq(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("path: %s", r.URL.Path)
		next(w, r)
	})
}
