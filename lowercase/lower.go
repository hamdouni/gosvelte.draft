package app

import "strings"

// Min est une fonction business qui met tous les caractères passés en paramètre en minuscule
func Min(s string) string {
	res := strings.ToLower(s)
	return res
}
