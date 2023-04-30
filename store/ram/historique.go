package ram

func (rs *RAM) Save(s string) {
	rs.historic = append(rs.historic, s)
}

func (rs RAM) List() []string {
	return rs.historic
}
