package main

import (
	"app/api"
	"app/bdd"
	"app/biz"
	"app/sec"
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
	var storage bdd.RamStore
	storage.Init()

	// on crée notre biz avec notre store
	var biz biz.BIZ
	biz.Init(&storage)
	if err := biz.CreateUser("maximilien", "motdepasse"); err != nil {
		log.Fatalf("Impossible de créer un utilisateur de test : %v", err)
	}

	// on crée notre sécurité
	var sec sec.Secure
	if err := sec.Init(); err != nil {
		log.Fatalf("Impossible d'initialiser la sécurité de l'application : %v", err)
	}

	// on crée une api avec notre biz et notre sec
	var api api.API
	api.Init(biz, sec)

	log.Printf("Le service démarre sur le port %v \n", port)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Erreur de démarrage du service sur le port %v : %v\n", port, err)
	}
}
