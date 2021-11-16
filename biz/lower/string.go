package lower

import "strings"

// String met en minuscule la chaîne de caractères
func String(s string) string {
	res := strings.ToLower(s)
	return res
}
