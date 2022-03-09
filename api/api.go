package api

import (
	"encoding/json"
	"log"
	"net/http"
)

const tokenCookieName = "jeton"

/*
	NewRoutes initialise toutes les routes (url) en spécifiant les fonctions
	qui les prennent en charge, soit des fonctions à nous qu'on déclare avec
	HandleFunc, soit la fonction qui s'occupe de servir les fichiers statiques.
*/
func NewRoutes(htmlDirectory string) error {
	/*
		Pour les fichiers statiques (html, js, images, ...), la librairie
		standard propose une fonction FileServer qui reçoit le dossier
		contenant nos fichiers statiques.
	*/
	fs := http.FileServer(http.Dir(htmlDirectory))
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
		log.Printf("respondJSON error %s", err)
	}
}
