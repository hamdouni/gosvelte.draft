package main

import (
	"app/biz"
	"app/store"
	"app/web"
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

	var storage store.Store
	storage.Init()
	// On ajoute un user de test
	user, err := biz.NewUser("maximilien", "motdepasse")
	if err != nil {
		return fmt.Errorf("impossible de créer un utilisateur de test : %v", err)
	}
	storage.AddUser(*user)

	var api web.WEB
	if err := api.Init(&storage, "./static"); err != nil {
		return fmt.Errorf("impossible d'initialiser la sécurité de l'application : %v", err)
	}

	log.Printf("Le service démarre sur le port %v \n", *port)
	return http.ListenAndServe(addr, nil)
}
