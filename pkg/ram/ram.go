package ram

import app "app/create_user"

type RamStore struct {
	ramhistoric []string
	ramusers    map[string]string
}

func (rs *RamStore) Init() {
	rs.ramusers = make(map[string]string)
}

func (rs *RamStore) StockHistorique(s string) {
	rs.ramhistoric = append(rs.ramhistoric, s)
}

func (rs RamStore) ListeHistorique() []string {
	return rs.ramhistoric
}

func (rs *RamStore) AddUser(user app.User) error {
	rs.ramusers[user.Username] = user.Password
	return nil
}

func (rs *RamStore) GetPasswordUser(username string) (encryptedPassword string) {
	return rs.ramusers[username]
}
