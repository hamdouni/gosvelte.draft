package biz_test

import "testing"

func TestUpper(t *testing.T) {
	given := "majuscule"
	wait := "MAJUSCULE"
	got := fakeBiz.Maj(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}
