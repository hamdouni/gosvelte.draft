package upper

import "strings"

// String met en majuscule la chaîne de caractères
func String(s string) string {
	res := strings.ToUpper(s)
	return res
}
