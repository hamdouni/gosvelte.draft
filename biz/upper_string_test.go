package biz_test

import (
	"app/biz"
	"testing"
)

func TestUpper(t *testing.T) {
	given := "majuscule"
	wait := "MAJUSCULE"
	got := biz.Upper(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}
