package api_test

import (
	api "app/api"
	"app/store"
)

func init() {
	var fakeStore store.Store
	fakeStore.Init()
	var fakeWeb api.API
	fakeWeb.Init(&fakeStore, ".")
}
