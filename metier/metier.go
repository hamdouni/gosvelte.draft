package metier

import (
	"webtoolkit/metier/historic"
	"webtoolkit/metier/secure"
	"webtoolkit/metier/user"
)

func Configure(uStore user.Storage, hStore historic.Storage) error {
	// initialise le composant de sécurité
	err := secure.Init()
	if err != nil {
		return err
	}

	user.UseStore(uStore)
	historic.UseStore(hStore)

	return nil
}
