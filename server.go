package main

import (
	"app/api"
	"app/biz"
	"app/store"
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
	host := flags.String("host", "127.0.0.1", "host name to listen on")
	port := flags.Int("port", 8000, "specify port to listen on")
	help := flags.Bool("help", false, "show command usage")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if *help {
		flags.Usage()
		return nil
	}

	addr := fmt.Sprintf("%s:%d", *host, *port)

	var storage store.Store
	storage.Init()
	// On ajoute un user de test
	user, err := biz.NewUser("test", "test")
	if err != nil {
		return fmt.Errorf("impossible de créer un utilisateur de test : %v", err)
	}
	storage.AddUser(*user)

	var api api.API
	if err := api.Init(&storage, "./client/static"); err != nil {
		return fmt.Errorf("impossible d'initialiser l'api: %v", err)
	}

	log.Printf("Le service démarre sur %s:%v \n", *host, *port)
	return http.ListenAndServe(addr, nil)
}
