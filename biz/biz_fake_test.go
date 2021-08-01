package biz_test

type fakeStoreType struct{}

func (fakeStoreType) StockHistorique(string) {}
func (fakeStoreType) ListeHistorique() []string {
	return []string{attendu}
}
func (fakeStoreType) AddUser(string, string) {}
func (fakeStoreType) GetPasswordUser(p string) string {
	return p
}

type fakeSecType struct{}

func (fakeSecType) Encrypt(b string) (string, error) {
	return b, nil
}
func (fakeSecType) Decrypt(b string) (string, error) {
	return b, nil
}
