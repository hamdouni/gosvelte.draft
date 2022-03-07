package app

import "admin/model"

func CheckPassword(username, password string) bool {
	hashed := app.store.GetPasswordUser(username)

	return app.secure.CheckPassword(password, hashed)
}

func StockHistorique(message string) {
	app.store.StockHistorique(message)
}

func ListeHistorique() []string {
	return app.store.ListeHistorique()
}

func AddUser(username, password string, role model.Role) error {
	hashed, err := app.secure.HashPassword(password)
	if err != nil {
		return err
	}
	app.store.AddUser(username, hashed, role)

	return nil
}
