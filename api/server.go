package api

import (
	"fmt"
	"net/http"
)

// server sert des pages dynamiques et statiques depuis le dossier RootDir en
// écoutant sur Address de la forme `host:port` par exemple `acme.com:80`.
type server struct {
	RootDir string
	Address string
}

// New retourne un serveur à partir d'un dossier, d'un host et d'un port
func New(dir, host string, port int) server {
	return server{
		RootDir: dir,
		Address: fmt.Sprintf("%s:%d", host, port),
	}
}

// Run exécute le serveur qui sert les fichiers statiques et les routes.
// Pour les fichiers statiques (html, js, images, ...), la librairie standard
// propose une fonction FileServer. Pour les traitements dynamiques, on déclare
// les routes en gérant le fait qu'elles doivent être authentifiées ou non.
func (s server) Run() error {
	fs := http.FileServer(http.Dir(s.RootDir))
	http.Handle("/", fs)

	Routes()

	return http.ListenAndServe(s.Address, nil)
}
