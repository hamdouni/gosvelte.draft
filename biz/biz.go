package biz

import (
	"wtk/biz/secure"
	"wtk/biz/user"
)

func Initialize(s Storage) error {

	// initialise le composant de sécurité
	err := secure.Init()
	if err != nil {
		return err
	}

	user.WithRepo(s)

	return nil
}
