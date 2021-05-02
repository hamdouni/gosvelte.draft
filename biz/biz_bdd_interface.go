package biz

type historicStore interface {
	Stock(string)
	Liste() []string
}
