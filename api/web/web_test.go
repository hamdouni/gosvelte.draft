package web_test

import (
	"app/api/web"
)

const attendu = "Fake Biz"

func init() {
	var fakeBiz fakeBizType
	var fakeWeb web.WEB
	fakeWeb.Init(fakeBiz)
}
