package usecase_test

import (
	"app/usecase"
	"testing"
)

func TestUpper(t *testing.T) {
	given := "majuscule"
	wait := "MAJUSCULE"
	got := usecase.Upper(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}
