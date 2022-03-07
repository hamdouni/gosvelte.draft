package web

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const tokenCookieName = "jeton"
const tokenTimeLayout = time.RFC3339
const tokenDuration = time.Duration(24) * time.Hour

type API struct {
}

func New(htmlDirectory string) error {
	InitRoutes(htmlDirectory)
	return nil
}

/*
	InitRoutes initialise toutes les routes (url) en spécifiant les fonctions qui les prennent en charge, soit des fonctions à nous qu'on déclare avec HandleFunc, soit la fonction qui s'occupe de servir les fichiers statiques.
*/
func InitRoutes(htmlDirectory string) {
	/*
		Voici des exemples de déclaration de routes avec l'url pour y accéder et la fonction qui va la traiter.
		Par exemple l'url "/hello" est prise en charge par api.Hello que l'on doit implémenter (cf fichier api_hello.go).
	*/
	logHandleFunc("/hello", handleHello)
	logHandleFunc("/login", handleLogin)
	logHandleFunc("/logout", handleLogout)
	logHandleFunc("/check", handleLogCheck)
	logHandleFunc("/historic", handleAuth(Historic))
	logHandleFunc("/upper", handleAuth(Upper))
	logHandleFunc("/lower", handleAuth(Lower))

	/*
		Pour les fichiers statiques (html, js, images, ...), la librairie standard propose une fonction FileServer qui reçoit le dossier contenant nos fichiers statiques.
	*/
	fs := http.FileServer(http.Dir(htmlDirectory))
	http.Handle("/", fs)
}

/*
	Pour tracer les requêtes, on enrichie http.HandleFunc avec 2 fonctions :
	- logHandleFunc sera utilisé à la place de http.HandleFunc dans la déclaration des routes
	- logReq affiche la requête et continue le traitement sur le handler d'origine
*/
func logHandleFunc(pattern string, handler http.HandlerFunc) {
	http.HandleFunc(pattern, logReq(handler))
}
func logReq(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	})
}

// respondJSON retourne une reponse json avec le statut et le contenu passés en paramètre
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err = w.Write([]byte(response)); err != nil {
		log.Printf("respondJSON error %v", err)
	}
}
