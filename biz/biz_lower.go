package biz

import "strings"

/*
	Maj est une fonction business qui met tous les caractères passés en paramètre en majuscule
*/
func (b BIZ) Maj(s string) string {
	return strings.ToUpper(s)
}
