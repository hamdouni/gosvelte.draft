package web

type store interface {
	StockHistorique(string)
	ListeHistorique() []string
}
