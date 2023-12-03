package main

import (
	"testing"
)

const (
	GOOD_GAME string = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	BAD_GAME  string = "Game 2: 3 blue, 15 red; 1 red, 2 green, 6 blue; 2 green"
)

func TestGetGameId(t *testing.T) {
	got := GetGameId(GOOD_GAME)

	if got != 1 {
		t.Errorf("got %d expected 1", got)
	}
}

func TestNewSet(t *testing.T) {
	got := NewSet(GOOD_GAME)

	if got[0].Red != 4 {
		t.Errorf("got %d expected 4", got[0].Red)
	}
}

func TestIsPossible(t *testing.T) {
	gotSad := NewGame(BAD_GAME)
	gotHappy := NewGame(GOOD_GAME)

	sadPossible := gotSad.IsPossible()
	happyPossible := gotHappy.IsPossible()

	if sadPossible == true {
		t.Errorf("got %t expected true", sadPossible)
	}

	if happyPossible == false {
		t.Errorf("got %t expected false", happyPossible)
	}
}
