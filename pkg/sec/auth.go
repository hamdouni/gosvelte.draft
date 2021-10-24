package sec

type repo interface {
	GetPasswordUser(username string) (encryptedPassword string)
}

func (s *Secure) AuthUser(u, p string) bool {
	h := s.storage.GetPasswordUser(u)
	return s.CheckPassword(p, h)
}
