package model

type Historic struct {
	Evenements []string
}

var HistoricStore HistoricStorage

type HistoricStorage interface {
	StockHistorique(string)
	ListeHistorique() []string
}

func StockHistorique(message string) {
	HistoricStore.StockHistorique(message)
}

func ListeHistorique() []string {
	return HistoricStore.ListeHistorique()
}
