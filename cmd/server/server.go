package main

import (
	"app/pkg/ram"
	"app/pkg/sec"
	"app/pkg/web"
	"app/usecase"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	port := flags.Int("p", 8000, "specify port to listen on")
	help := flags.Bool("h", false, "show command usage")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if *help {
		flags.Usage()
		return nil
	}

	addr := fmt.Sprintf("0.0.0.0:%d", *port)

	// composant de stockage des données
	var storage ram.RamStore
	storage.Init()

	// composant de sécurité
	var sec sec.Secure
	if err := sec.Init(&storage); err != nil {
		return fmt.Errorf("impossible d'initialiser la sécurité de l'application : %v", err)
	}

	// api utilise biz.
	var api web.WEB
	api.Init(&sec, &storage, "./html")

	// On ajoute un user de test
	user, err := usecase.NewUser("maximilien", "motdepasse")
	if err != nil {
		return fmt.Errorf("impossible de créer un utilisateur de test : %v", err)
	}
	hp, err := sec.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("cannot hass password: %v", err)
	}
	user.Password = hp
	storage.AddUser(*user)

	log.Printf("Le service démarre sur le port %v \n", *port)
	return http.ListenAndServe(addr, nil)
}
