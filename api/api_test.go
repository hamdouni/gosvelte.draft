package api_test

import (
	"admin/api"
	"admin/model"
	"admin/store/ram"
)

func init() {
	fakeStore, _ := ram.New()
	model.Init(&fakeStore, &fakeStore)
	model.AddUser("samething", "samething", 1)
	api.NewRoutes(".")
}
