package model

import "admin/model/secure"

func Init(uStore UserStorage, hStore HistoricStorage) error {
	// initialise le composant de sécurité
	err := secure.Init()
	if err != nil {
		return err
	}
	UserStore = uStore
	HistoricStore = hStore

	return nil
}
