package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	have := []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}

	got := Part1(have)
	want := 7

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	have := []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}

	got := Part2(have)
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
