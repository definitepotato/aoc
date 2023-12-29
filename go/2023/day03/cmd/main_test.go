package main

import (
	"helpers"
	"testing"
)

func TestIsNumber(t *testing.T) {
	gotHappy := IsNumber("4")
	gotSad := IsNumber("a")

	if gotHappy != true {
		t.Errorf("got %v expected true", gotHappy)
	}

	if gotSad != false {
		t.Errorf("got %v expected false", gotSad)
	}
}

func TestScanParts(t *testing.T) {
	input := helpers.ReadFile("../test.txt")
	schematic := NewSchematic(input)

	if len(schematic.Parts) != 10 {
		t.Errorf("got %d, expected 10", len(schematic.Parts))
	}
}

func TestMarkParts(t *testing.T) {
	input := helpers.ReadFile("../test.txt")
	schematic := NewSchematic(input)

	sum := 0
	for _, part := range schematic.Parts {
		if part.IsPart {
			sum += part.Number
		}
	}

	if sum != 4361 {
		t.Errorf("got %d, expected 4361", sum)
	}
}
