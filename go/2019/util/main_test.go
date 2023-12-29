package util

import "testing"

func TestStoi(t *testing.T) {
	t.Run("string to integer", func(t *testing.T) {
		have := "10"
		got := Stoi(have)
		want := 10

		if got != want {
			t.Errorf("got %d, wanted %d\n", got, want)
		}
	})
}

func TestMinMax(t *testing.T) {
	t.Run("minmax", func(t *testing.T) {
		have := []int{1, 2, 3, 4, 5}
		gotMin, gotMax := MinMax(have)
		wantMin := 1
		wantMax := 5

		if gotMin != wantMin {
			t.Errorf("got %d, wanted %d\n", gotMin, wantMin)
		}

		if gotMax != wantMax {
			t.Errorf("got %d, wanted %d\n", gotMax, wantMax)
		}
	})
}

func TestSliceStringToInt(t *testing.T) {
	t.Run("description", func(t *testing.T) {
		have := []string{"1", "2", "3", "4", "5"}
		got := SliceStringToInt(have)
		want := []int{1, 2, 3, 4, 5}

		for i := 0; i < len(got); i++ {
			if got[i] != want[i] {
				t.Errorf("got %d, wanted %d\n", got[i], want[i])
			}
		}
	})
}

func TestSortString(t *testing.T) {
	t.Run("sort string", func(t *testing.T) {
		have := "cdeba"
		got := SortString(have)
		want := "abcde"

		if got != want {
			t.Errorf("got %s, wanted %s\n", got, want)
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("reverse", func(t *testing.T) {
		have := []string{"a", "b", "c", "d"}
		got := Reverse(have)
		want := []string{"d", "c", "b", "a"}

		for i := 0; i < len(got); i++ {
			if got[i] != want[i] {
				t.Errorf("got %s, wanted %s\n", got[i], want[i])
			}
		}
	})
}

func TestUnique(t *testing.T) {
	t.Run("unique", func(t *testing.T) {
		have := []string{"a", "b", "c", "c", "d", "d"}
		got := Unique(have)
		want := []string{"a", "b", "c", "d"}

		for i := 0; i < len(got); i++ {
			if got[i] != want[i] {
				t.Errorf("got %s, wanted %s\n", got[i], want[i])
			}
		}
	})
}
