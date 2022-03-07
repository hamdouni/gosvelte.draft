package model_test

import (
	"admin/model"
	"testing"
)

func TestUpper(t *testing.T) {
	given := "majuscule"
	wait := "MAJUSCULE"
	got := model.Upper(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}
