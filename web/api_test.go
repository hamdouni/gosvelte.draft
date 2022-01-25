package web_test

import (
	"app/store"
	"app/web"
)

func init() {
	var fakeStore store.Store
	fakeStore.Init()
	var fakeWeb web.API
	fakeWeb.Init(&fakeStore, ".")
}
