package bdd

type RamStore struct {
	ramhistoric []string
}

func (rs *RamStore) Stock(s string) {
	rs.ramhistoric = append(rs.ramhistoric, s)
}

func (rs RamStore) Liste() []string {
	return rs.ramhistoric
}
