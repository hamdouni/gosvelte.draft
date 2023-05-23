package biz

import (
	"webtoolkit/biz/secure"
	"webtoolkit/biz/user"
)

func Intialize(s user.Storage) error {

	// initialise le composant de sécurité
	err := secure.Init()
	if err != nil {
		return err
	}

	user.WithRepo(s)

	return nil
}
