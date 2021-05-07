package main

import (
	"app/api"
	"app/bdd"
	"app/biz"
	_ "embed"
	"log"
	"net/http"
)

const port = 8000

//go:embed VERSION
var version string

func main() {

	log.Printf("App version %v", version)

	// on crée un historique en mémoire
	var historique bdd.RamStore

	// on crée notre biz avec notre histo
	var biz biz.BIZ
	biz.Init(&historique)

	// on crée une api avec notre biz
	var api api.API
	api.Init(biz)

	log.Printf("Le service démarre sur le port %v \n", port)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Erreur de démarrage du service sur le port %v : %v\n", port, err)
	}
}
