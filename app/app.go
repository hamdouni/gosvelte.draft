// Le package `app` définie la couche applicative qui fait le lien entre tous
// les éléments de notre application : c'est la glue qui relie les entrées et
// les sorties, les traitements ou encore les contrôles nécessaires.

// Le package met à disposition les `services` nécessaires pour que les
// différentes couches coopèrent.

package app

import "admin/model"

func Config(uStore model.UserStorage, hStore model.HistoricStorage, secure model.Security) {
	model.UserStore = uStore
	model.HistoricStore = hStore
	model.Secure = secure
}
