package api_test

import (
	"app/api"
	"app/store/ram"
)

func init() {
	var fakeStore ram.Store
	fakeStore.Init()
	var fakeWeb api.API
	fakeWeb.Init(&fakeStore, ".")
}
