package ram

func (rs *RAM) StockHistorique(s string) {
	rs.historic = append(rs.historic, s)
}

func (rs RAM) ListeHistorique() []string {
	return rs.historic
}
