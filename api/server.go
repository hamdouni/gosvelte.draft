package api

import (
	"fmt"
	"net/http"
)

// server sert des pages dynamiques et statiques depuis le dossier RootDir en
// écoutant sur Address de la forme `host:port` par exemple `acme.com:80`.
type server struct {
	rootdir string
	address string
}

// New retourne un serveur prêt à servir des fichiers
// statiques ou des routes dynamiques
func New(static, host string, port int) server {

	Routes(static) // initialise les routes statiques et dynamiques

	return server{
		rootdir: static,
		address: fmt.Sprintf("%s:%d", host, port),
	}
}

// Run exécute le serveur
func (s server) Run() error {
	return http.ListenAndServe(s.address, nil)
}
