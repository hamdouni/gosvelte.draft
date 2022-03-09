package model_test

import (
	"admin/model"
	"testing"
)

func TestToken(t *testing.T) {
	model.Init(nil, nil)
	token, err := model.NewToken("fakeuser", "fakelocalisation")
	if err != nil {
		t.Fatalf("Not expecting error:%s", err)
	}
	ok, err := model.CheckToken(token, "otherlocalisation")
	if err != nil {
		t.Fatalf("Not expecting error:%s", err)
	}
	if ok {
		t.Fatalf("Expecting check token to return false got %v", ok)
	}
	ok, err = model.CheckToken(token, "fakelocalisation")
	if err != nil {
		t.Fatalf("Not expecting error:%s", err)
	}
	if !ok {
		t.Fatalf("Expecting check token to return true got %v", ok)
	}
}
