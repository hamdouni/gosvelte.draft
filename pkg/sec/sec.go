package sec

// Secure store key
type Secure struct {
	key     *[32]byte
	storage repo
}

// Init initialize secret key
func (s *Secure) Init(store repo) (err error) {
	s.storage = store
	s.key, err = newEncryptionKey()
	return err
}
