package web

import (
	"net/http"
)

/*
	InitRoutes initialise toutes les routes (url) en spécifiant les fonctions qui les prennent en charge, soit des fonctions à nous qu'on déclare avec HandleFunc, soit la fonction qui s'occupe de servir les fichiers statiques.
*/
func (web *WEB) InitRoutes(publicDirectory string) {
	/*
		Voici des exemples de déclaration de routes avec l'url pour y accéder et la fonction qui va la traiter.
		Par exemple l'url "/hello" est prise en charge par api.Hello que l'on doit implémenter (cf fichier api_hello.go).
	*/
	web.logHandleFunc("/hello", web.Hello)
	web.logHandleFunc("/login", web.Login)
	web.logHandleFunc("/logout", web.Logout)
	web.logHandleFunc("/check", web.LogCheck)
	web.logHandleFunc("/historic", web.auth(web.Historic))
	web.logHandleFunc("/upper", web.auth(web.Upper))
	web.logHandleFunc("/lower", web.auth(web.Lower))

	/*
		Pour les fichiers statiques (html, js, images, ...), la librairie standard propose une fonction FileServer qui reçoit le dossier contenant nos fichiers statiques.
	*/
	fs := http.FileServer(http.Dir(publicDirectory))
	http.Handle("/", fs)
}

/*
	Pour tracer les requêtes, on enrichie http.HandleFunc avec 2 fonctions :
	- logHandleFunc sera utilisé à la place de http.HandleFunc dans la déclaration des routes
	- logReq affiche la requête et continue le traitement sur le handler d'origine
*/
func (web *WEB) logHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, logReq(handler))
}
func logReq(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	})
}
