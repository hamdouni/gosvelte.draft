package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"wtk/biz"
	"wtk/biz/user"
	"wtk/cmd/api"
	"wtk/infra/store/ram"
	"wtk/infra/store/sqlite"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	host := flag.String("host", "0.0.0.0", "host name to listen on")
	port := flag.Int("port", 8000, "port to listen on")
	static := flag.String("static", "./cmd/cli/", "static files folder")
	help := flag.Bool("help", false, "show command usage")
	flag.Parse()

	if *help {
		flag.Usage()
		return nil
	}

	var initSchema = false

	// storage, initSchema, err := useRamDB()
	storage, initSchema, err := useSqliteDB()

	// configure le métier avec le storage
	// mais on peut imaginer d'avoir un système de stockage différencié
	biz.Initialize(storage)

	if initSchema {
		err = storage.InitSchema()
		if err != nil {
			return fmt.Errorf("impossible de créer le schéma de la base: %w", err)
		}
		// ajoute un user de test
		err = user.Add("test", "test", "test", user.Administrator)
		if err != nil {
			return fmt.Errorf("impossible de créer un utilisateur de test: %w", err)
		}

		// ajoute 5 fake user dans realm test
		for i := 1; i <= 5; i++ {
			err = user.Add("test", "test"+strconv.Itoa(i), "test", user.Customer)
		}
		// ajoute 5 fake user dans realm fakerealm
		for i := 1; i <= 5; i++ {
			err = user.Add("fakerealm", "test"+strconv.Itoa(i), "test", user.Customer)
		}
	}

	// initialise le serveur api
	server := api.New(*static, *host, *port)

	// lance le server
	return server.Run()
}

// composant de stockage en RAM
func useRamDB() (store user.Storage, needSchema bool, err error) {
	st, err := ram.New()
	if err != nil {
		return store, false, err
	}
	return &st, true, nil
}

// composant de stockage sqlite
func useSqliteDB() (store user.Storage, needSchema bool, err error) {
	const pragma = "_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=busy_timeout(8000)&_pragma=journal_size_limit(100000000)"
	const databasePath = "database.db"
	_, err = os.Stat(databasePath)
	if errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(databasePath)
		if err != nil {
			return store, false, err
		}
		needSchema = true
	} else if err != nil {
		return store, false, err
	}

	store, err = sqlite.New(databasePath, pragma)
	if err != nil {
		return store, false, err
	}

	return store, needSchema, err
}
