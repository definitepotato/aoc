package main

import "testing"

func TestMix(t *testing.T) {
	t.Run("mix secret", func(t *testing.T) {
		got := Mix(15, 42)
		want := 37

		if got != want {
			t.Errorf("got {%v}, wanted {%v}", got, want)
		}
	})
}

func TestPrune(t *testing.T) {
	t.Run("prune secret", func(t *testing.T) {
		have := 100000000
		got := Prune(have)
		want := 16113920

		if got != want {
			t.Errorf("got {%v}, wanted {%v}", got, want)
		}
	})
}

func TestNextSecret(t *testing.T) {
	t.Run("next secret", func(t *testing.T) {
		have := 123
		got := NextSecret(have)
		want := 15887950

		if got != want {
			t.Errorf("got {%v}, wanted {%v}", got, want)
		}
	})
}
