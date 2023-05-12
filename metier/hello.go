package metier

// Hello retourne bonjour et le nom passé en paramètre.
func Hello(s string) string {
	res := "Bonjour " + s + " depuis le métier !"
	return res
}
