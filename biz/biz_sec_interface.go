package biz

type secure interface {
	Init() error
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}
