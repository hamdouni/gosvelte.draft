package biz_test

import "testing"

func TestBonjour(t *testing.T) {
	given := "Jean"
	wait := "Bonjour Jean depuis le business !"
	got := fakeBiz.Bonjour(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}
