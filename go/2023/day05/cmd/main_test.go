package main

import (
	"helpers"
	"testing"
)

const MAP string = "50 98 2"

func TestWriteMap(t *testing.T) {
	got := WriteMap(MAP)

	if got.Destination != 50 {
		t.Errorf("Destination is %d, expected 50", got.Destination)
	}

	if got.Source != 98 {
		t.Errorf("Source is %d, expected 98", got.Source)
	}

	if got.Range != 2 {
		t.Errorf("Range is %d, expected 2", got.Range)
	}
}

func TestCollectSeeds(t *testing.T) {
	input := helpers.ReadFile("../test.txt")
	got := CollectSeeds(input)

	if len(got) != 4 {
		t.Errorf("Len of seeds is %d, expected 4", len(got))
	}
}

func TestNewMap(t *testing.T) {
	input := []string{
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
	}
	got := NewMap(input)

	if got.Name == "" {
		t.Errorf("Name is empty, expected a value")
	}

	if len(got.Maps) == 0 {
		t.Errorf("No maps found, expected 2 maps")
	}
}

func TestNewAlamanac(t *testing.T) {
	input := helpers.ReadFile("../test.txt")
	got := NewAlmanac(input)

	s1 := got.GetSeedLocation(79)
	s2 := got.GetSeedLocation(14)
	s3 := got.GetSeedLocation(55)
	s4 := got.GetSeedLocation(13)

	if s1 != 82 {
		t.Errorf("got %d, expected 82", s1)
	}

	if s2 != 43 {
		t.Errorf("got %d, expected 43", s2)
	}

	if s3 != 86 {
		t.Errorf("got %d, expected 86", s3)
	}

	if s4 != 35 {
		t.Errorf("got %d, expected 35", s4)
	}
}
