package biz_test

import "testing"

func TestLower(t *testing.T) {
	given := "MINUSCULE"
	wait := "minuscule"
	got := fakeBiz.Min(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}
