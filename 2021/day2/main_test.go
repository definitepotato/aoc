package main

// 1: 150
// 2: 900

import (
	"testing"
)

func TestPart1(t *testing.T) {
	have := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	got := Part1(have)
	want := 150

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	have := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	got := Part2(have)
	want := 900

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
