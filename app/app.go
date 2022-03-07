// Le package `app` définie la couche applicative qui fait le lien entre tous
// les éléments de notre application : c'est la glue qui relie les entrées et
// les sorties, les traitements ou encore les contrôles nécessaires.

// Le package met à disposition les `services` nécessaires pour que les
// différentes couches coopèrent.

package app

import "admin/model"

var app struct {
	store  storage
	secure security
}

func Config(store storage, secure security) {
	app.store = store
	app.secure = secure
}

// Contrat avec le service de sécurité
type security interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
	HashPassword(pw string) (encryptedPassword string, err error)
	CheckPassword(pw, hashed string) bool
}

// Contrat avec le service de stockage
type storage interface {
	StockHistorique(string)
	ListeHistorique() []string
	GetPasswordUser(username string) (encryptedPassword string)
	AddUser(username, password string, role model.Role) error
}
