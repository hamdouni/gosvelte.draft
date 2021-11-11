package biz_test

import (
	"app/biz"
	"testing"
)

func TestLower(t *testing.T) {
	given := "MINUSCULE"
	wait := "minuscule"
	got := biz.Lower(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}
