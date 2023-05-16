package biz

import (
	"webtoolkit/biz/secure"
	"webtoolkit/biz/user"
)

func Configure(s user.Storage) error {

	// initialise le composant de sécurité
	err := secure.Init()
	if err != nil {
		return err
	}

	user.UseStore(s)

	return nil
}
