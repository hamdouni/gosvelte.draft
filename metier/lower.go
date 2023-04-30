package metier

import "strings"

// Lower met en minuscule la chaîne de caractères
func Lower(s string) string {
	res := strings.ToLower(s)
	return res
}
