package main

import (
	"app/api"
	"app/biz/create"
	"app/store/ram"
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
	port := flags.Int("port", 8000, "specify port to listen on")
	static := flags.String("static", "./client/static", "folder to html css js static files")
	help := flags.Bool("help", false, "show command usage")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if *help {
		flags.Usage()
		return nil
	}

	addr := fmt.Sprintf("%s:%d", *host, *port)

	var storage ram.Store
	storage.Init()
	// On ajoute un user de test
	user, err := create.NewUser("test", "test", create.Administrator)
	if err != nil {
		return fmt.Errorf("impossible de créer un utilisateur de test : %v", err)
	}
	storage.AddUser(user.Username, user.Password, user.Role)

	var api api.API
	if err := api.Init(&storage, *static); err != nil {
		return fmt.Errorf("impossible d'initialiser l'api: %v", err)
	}

	log.Printf("Le service démarre sur %s:%v \n", *host, *port)
	return http.ListenAndServe(addr, nil)
}
