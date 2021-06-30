package biz

type BIZ struct {
	store  store
	secret secure
}

func (b *BIZ) Init(h store, s secure) {
	b.store = h
	b.secret = s
}
