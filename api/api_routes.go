package api

import "net/http"

/*
	InitRoutes initialise toutes les routes (url) en spécifiant les fonctions qui les prennent en charge, soit des fonctions à nous qu'on déclare avec HandleFunc, soit la fonction qui s'occupe de servir les fichiers statiques.
*/
func (api *API) InitRoutes() {
	/*
		Voici un exemple de déclaration d'une route avec l'url pour y accéder "/hello" et la fonction qui va la prendre en charge api.Hello que l'on doit implémenter (cf fichier api_hello.go).
	*/
	http.HandleFunc("/hello", api.Hello)

	/*
		Les autres routes
	*/
	http.HandleFunc("/login", api.Login)
	http.HandleFunc("/logout", api.Logout)
	http.HandleFunc("/check", api.LogCheck)
	http.HandleFunc("/historic", api.auth(api.Historic))
	http.HandleFunc("/upper", api.auth(api.Upper))
	http.HandleFunc("/lower", api.auth(api.Lower))

	/*
		Pour les fichiers statiques (html, js, images, ...), la librairie standard propose une fonction FileServer qui reçoit le dossier contenant nos fichiers statiques.
	*/
	fs := http.FileServer(http.Dir("./ihm/public"))
	http.Handle("/", fs)
}
