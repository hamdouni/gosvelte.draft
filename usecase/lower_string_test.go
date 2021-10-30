package usecase_test

import (
	"app/usecase"
	"testing"
)

func TestLower(t *testing.T) {
	given := "MINUSCULE"
	wait := "minuscule"
	got := usecase.Lower(given)
	if got != wait {
		t.Fatalf("Waiting %v but got %v", wait, got)
	}
}
