package main

import (
	"app"
	"app/api/web"
	"app/biz"
	"app/infra/bdd"
	"app/infra/sec"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const port = 8000

func main() {
	log.Printf("App version %v", app.Version)
	if err := run(os.Args); err != nil {
		log.Fatalf("Erreur de démarrage du service sur le port %v : %v\n", port, err)
	}
}

func run(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		port = flags.Int("port", 8000, "port to listen on")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	addr := fmt.Sprintf("0.0.0.0:%d", *port)

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
		return fmt.Errorf("impossible d'initialiser la sécurité de l'application : %v", err)
	}

	// On crée ensuite notre biz avec ces 2 composants
	var biz biz.BIZ
	biz.Init(&storage, sec)

	// On ajoute un user de test
	if err := biz.CreateUser("maximilien", "motdepasse"); err != nil {
		return fmt.Errorf("Impossible de créer un utilisateur de test : %v", err)
	}

	// Enfin, on crée l'api avec ce biz.
	var api web.WEB
	api.Init(biz)

	log.Printf("Le service démarre sur le port %v \n", *port)
	return http.ListenAndServe(addr, nil)
}
