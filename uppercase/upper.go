package app

import "strings"

// Maj est une fonction business qui met tous les caractères passés en paramètre en majuscule
func Maj(s string) string {
	res := strings.ToUpper(s)
	return res
}
