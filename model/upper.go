package model

import "strings"

// Upper met en majuscule la chaîne de caractères
func Upper(s string) string {
	res := strings.ToUpper(s)
	return res
}
