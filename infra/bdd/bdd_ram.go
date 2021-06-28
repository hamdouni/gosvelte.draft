package bdd

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

func (rs *RamStore) AddUser(k, v string) {
	rs.ramusers[k] = v
}

func (rs *RamStore) GetPasswordUser(k string) (v string) {
	return rs.ramusers[k]
}
