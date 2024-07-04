package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	got := add(7, 5)
	want := 12
	if got != want {
		log.Fatalf("Error - waant %v and got %v ", want, got)
	}

	got2 := paradise("Bali")
	want2 := "My idea of paradise is Bali"
	if got2 != want2 {
		log.Fatalf("error - want %v and got %v", want2, got2)
	}
}
