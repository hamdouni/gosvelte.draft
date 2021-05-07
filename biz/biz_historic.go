package biz

func (b BIZ) Historic() []string {
	return b.store.ListeHistorique()
}
