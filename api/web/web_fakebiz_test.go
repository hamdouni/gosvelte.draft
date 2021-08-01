package web_test

type fakeBizType struct{}

func (fakeBizType) Bonjour(string) string {
	return attendu
}
func (fakeBizType) Maj(string) string {
	return attendu
}
func (fakeBizType) Min(string) string {
	return attendu
}
func (fakeBizType) Historic() []string {
	return []string{attendu}
}

// si on envoi un username et password identique on renvoit toujours vrai
// sinon on renvoi faux
func (fakeBizType) CheckPassword(u, p string) bool {
	return u == p
}
func (fakeBizType) Encrypt(b string) (string, error) {
	return b, nil
}
func (fakeBizType) Decrypt(b string) (string, error) {
	return b, nil
}
