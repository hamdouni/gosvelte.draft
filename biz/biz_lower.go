package biz

import "strings"

/*
	Maj est une fonction business qui met tous les caractères passés en paramètre en minuscule
*/
func (b BIZ) Min(s string) string {
	res := strings.ToLower(s)
	b.store.StockHistorique(res)
	return res
}
