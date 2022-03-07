package model_test

import (
	"admin/model"
	"testing"
)

func TestLower(t *testing.T) {
	given := "MINUSCULE"
	wait := "minuscule"
	got := model.Lower(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}
