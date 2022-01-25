package secure

// Secure store key
type Secure struct {
	key *[32]byte
}

// Init initialize secret key
func (s *Secure) Init() (err error) {
	s.key, err = newEncryptionKey()
	return err
}
