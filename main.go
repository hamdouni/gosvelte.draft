package main

import (
	"admin/api"
	"admin/model"
	"admin/store/ram"
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
	host := flags.String("host", "0.0.0.0", "host name to listen on")
	port := flags.Int("port", 80, "port to listen on")
	static := flags.String("static", "./client/static", "static files folder")
	help := flags.Bool("help", false, "show command usage")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if *help {
		flags.Usage()
		return nil
	}

	addr := fmt.Sprintf("%s:%d", *host, *port)

	// composant de stockage en RAM
	storage, err := ram.New()
	if err != nil {
		return err
	}

	// configure le modèle avec le storage
	// ici on utilise le même storage pour l'historique et les users
	// mais on peut imaginer d'avoir un système de stockage différencié
	model.Init(&storage, &storage)

	// ajoute un user de test
	err = model.AddUser("test", "test", model.Administrator)
	if err != nil {
		return fmt.Errorf("impossible de créer un utilisateur de test : %s", err)
	}

	// initialise les routes de l'api
	err = api.Routes(*static)
	if err != nil {
		return fmt.Errorf("impossible d'initialiser l'api: %s", err)
	}

	// démarre le serveur
	log.Printf("Le service démarre sur %s:%d \n", *host, *port)
	return http.ListenAndServe(addr, nil)
}
