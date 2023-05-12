package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"webtoolkit/api"
	"webtoolkit/metier"
	"webtoolkit/metier/user"
	"webtoolkit/store/ram"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	host := flag.String("host", "0.0.0.0", "host name to listen on")
	port := flag.Int("port", 80, "port to listen on")
	static := flag.String("static", "./cmd/client/static", "static files folder")
	help := flag.Bool("help", false, "show command usage")
	flag.Parse()

	if *help {
		flag.Usage()
		return nil
	}

	// composant de stockage en RAM
	storage, err := ram.New()
	if err != nil {
		return err
	}

	// configure le métier avec le storage
	// ici on utilise le même storage pour l'historique et les users
	// mais on peut imaginer d'avoir un système de stockage différencié
	metier.Configure(&storage, &storage)

	// ajoute un user de test
	err = user.Add("test", "test", "test", user.Administrator)
	if err != nil {
		return fmt.Errorf("impossible de créer un utilisateur de test : %w", err)
	}

	// ajoute 5 fake user dans realm test
	for i := 1; i <= 5; i++ {
		err = user.Add("test", "test"+strconv.Itoa(i), "test", user.Customer)
	}
	// ajoute 5 fake user dans realm fakerealm
	for i := 1; i <= 5; i++ {
		err = user.Add("fakerealm", "test"+strconv.Itoa(i), "test", user.Customer)
	}

	// initialise le serveur api
	server := api.New(*static, *host, *port)

	// lance le server
	return server.Run()
}
