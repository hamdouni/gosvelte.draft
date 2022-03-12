package api

import (
	"net/http"
)

// Server sert des pages dynamiques et statiques depuis le dossier RootDir en
// écoutant sur Address de la forme `host:port` par exemple `acme.com:80`.
type Server struct {
	RootDir string
	Address string
}

// Run exécute le serveur qui sert les fichiers statiques et les routes.
// Pour les fichiers statiques (html, js, images, ...), la librairie standard
// propose une fonction FileServer. Pour les traitements dynamiques, on déclare
// les routes en gérant le fait qu'elles doivent être authentifiées ou non.
func (s Server) Run() error {
	fs := http.FileServer(http.Dir(s.RootDir))
	http.Handle("/", fs)

	Routes()

	return http.ListenAndServe(s.Address, nil)
}
