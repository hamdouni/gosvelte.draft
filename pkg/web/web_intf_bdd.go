package web

type storage interface {
	StockHistorique(string)
	ListeHistorique() []string
	GetPasswordUser(username string) (encryptedPassword string)
}
