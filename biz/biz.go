package biz

type BIZ struct {
	store store
}

func (b *BIZ) Init(h store) {
	b.store = h
}
