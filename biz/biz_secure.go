package biz

/*
Ici biz ne fait que le passe plat.
La vraie implémentation est dans le packet infra/secure.
Mais on peut imaginer que la sécurité est une expertise métier
et dans ce cas, l'implémentation se ferait ici.
*/

func (b BIZ) Encrypt(plaintext string) (ciphertext string, err error) {
	enc, err := b.secret.Encrypt(plaintext)
	if err != nil {
		return "", err
	}
	return string(enc), nil
}

func (b BIZ) Decrypt(ciphertext string) (plaintext string, err error) {
	dec, err := b.secret.Decrypt(ciphertext)
	if err != nil {
		return "", err
	}
	return string(dec), nil
}
