package main

import (
	"fmt"
	"testing"
)

func TestSeek(t *testing.T) {
	have := "xsone7nineEFthree5onesix"
	first, last := Seek(have)

	fmt.Println(first, last)
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

func TestSeekWithWords(t *testing.T) {
	have := "xsone7nineEFthree5onesix"
	first, last := SeekWithWords(have)

	if first != "1" && last != "6" {
		t.Errorf("got %s,%s expected 1,6", first, last)
	}
}
