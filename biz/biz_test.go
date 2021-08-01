package biz_test

import "app/biz"

const attendu = "Fake Biz"

var fakeBiz biz.BIZ

func init() {
	var fakeStore fakeStoreType
	var fakeSec fakeSecType
	fakeBiz.Init(fakeStore, fakeSec)
}
