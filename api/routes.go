package api

import "net/http"

/*
	Routes initialise toutes les routes (url) en spécifiant les fonctions
	qui les prennent en charge, soit des fonctions à nous qu'on déclare avec
	HandleFunc, soit la fonction qui s'occupe de servir les fichiers statiques.
*/
func Routes(dir string) error {
	/*
		Pour les fichiers statiques (html, js, images, ...), la librairie
		standard propose une fonction FileServer qui reçoit le dossier
		contenant nos fichiers statiques.
	*/
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	/*
		Pour les traitements dynamiques, on déclare la route depuis l'url vers
		la fonction de prise en charge.
		Par exemple l'url "/hello" est prise en charge par la fonction
		handleHello du fichier hello.go.
	*/
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/check", handleLogCheck)
	http.HandleFunc("/historic", handleAuth(Historic))
	http.HandleFunc("/upper", handleAuth(Upper))
	http.HandleFunc("/lower", handleAuth(Lower))

	return nil
}
