package main

import (
	"app/api/web"
	"app/biz"
	"app/infra/bdd"
	"app/infra/sec"
	_ "embed"
	"log"
	"net/http"
)

const port = 8000

//go:embed VERSION
var version string

func main() {

	log.Printf("App version %v", version)

	/*
		Notre application aura ce schéma général :

		[API] - [BIZ] - [INFRA (DATA|SEC)]

		On commence par créer notre infrastructure avec 2 composants :
		- un store en RAM pour les données
		- une implémentation 256-bit AES-GCM pour la sécurité
	*/

	var storage bdd.RamStore
	storage.Init()

	var sec sec.Secure
	if err := sec.Init(); err != nil {
		log.Fatalf("Impossible d'initialiser la sécurité de l'application : %v", err)
	}

	/*
		On crée ensuite notre biz avec ces 2 composants
	*/
	var biz biz.BIZ
	biz.Init(&storage, sec)

	/*
		On ajoute un user de test
	*/
	if err := biz.CreateUser("maximilien", "motdepasse"); err != nil {
		log.Fatalf("Impossible de créer un utilisateur de test : %v", err)
	}

	/*
		Enfin, on crée l'api avec ce biz.
	*/
	var api web.WEB
	api.Init(biz)

	log.Printf("Le service démarre sur le port %v \n", port)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Erreur de démarrage du service sur le port %v : %v\n", port, err)
	}
}
