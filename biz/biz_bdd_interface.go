package biz

type store interface {
	StockHistorique(string)
	ListeHistorique() []string
	AddUser(string, string)
	GetPasswordUser(string) string
}
