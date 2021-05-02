package biz

type BIZ struct {
	histo historicStore
}

func (b *BIZ) Init(h historicStore) {
	b.histo = h
}
