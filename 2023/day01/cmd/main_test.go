package main

import (
	"testing"
)

func TestSeekFromLeft(t *testing.T) {
	have := "twothree5"
	got := SeekFromLeft(have)

	if got != "5" {
		t.Errorf("SeekFromLeft: got %s expected 5", got)
	}
}

func TestSeekFromRight(t *testing.T) {
	have := "twothree5"
	got := SeekFromRight(have)

	if got != "5" {
		t.Errorf("SeekFromRight: got %s expected 5", got)
	}
}

func TestIsNum(t *testing.T) {
	five, happy := IsNum("5")
	letterA, sad := IsNum("a")

	if happy != true && five != 5 {
		t.Errorf("got %v %d expected true 5", happy, five)
	}

	if sad != false && letterA == 0 {
		t.Errorf("got %v %d expected false 0", sad, letterA)
	}
}

func TestSeekWordsAsNum(t *testing.T) {
	have := "xsone7nineEFthree5onesix"
	first, last := SeekWordsAsNum(have)

	if first != "1" && last != "6" {
		t.Errorf("got %s,%s expected 1,6", first, last)
	}
}
