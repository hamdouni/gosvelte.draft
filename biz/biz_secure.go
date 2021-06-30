package biz

/*
Ici biz ne fait que le passe plat.
La vraie implémentation est dans le packet infra/secure.
Mais on peut imaginer que la sécurité est une expertise métier
et dans ce cas, l'implémentation se ferait ici.
*/

func (b BIZ) Encrypt(plaintext []byte) (ciphertext []byte, err error) {
	return b.secret.Encrypt(plaintext)
}

func (b BIZ) Decrypt(ciphertext []byte) (plaintext []byte, err error) {
	return b.secret.Decrypt(ciphertext)
}
