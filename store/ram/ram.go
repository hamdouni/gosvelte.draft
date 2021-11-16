package ram

type Store struct {
	historic []string
	users    map[string]string
}

func (rs *Store) Init() {
	rs.users = make(map[string]string)
}

func (rs *Store) StockHistorique(s string) {
	rs.historic = append(rs.historic, s)
}

func (rs Store) ListeHistorique() []string {
	return rs.historic
}

func (rs *Store) AddUser(username, password string) error {
	rs.users[username] = password
	return nil
}

func (rs *Store) GetPasswordUser(username string) (encryptedPassword string) {
	return rs.users[username]
}
