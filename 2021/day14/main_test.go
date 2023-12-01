package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	have := []string{
		"NNCB",
		"",
		"CH -> B",
		"HH -> N",
		"CB -> H",
		"NH -> C",
		"HB -> C",
		"HC -> B",
		"HN -> C",
		"NN -> C",
		"BH -> H",
		"NC -> B",
		"NB -> B",
		"BN -> B",
		"BB -> N",
		"BC -> B",
		"CC -> N",
		"CN -> C",
	}

	got := Part1(have)
	want := 1588

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
