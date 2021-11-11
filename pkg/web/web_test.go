package web_test

import (
	"app/pkg/ram"
	"app/pkg/web"
)

func init() {
	var fakeStore ram.Store
	fakeStore.Init()
	var fakeWeb web.WEB
	fakeWeb.Init(&fakeStore, ".")
}
