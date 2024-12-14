package main

import "testing"

func TestAbsolute(t *testing.T) {
	t.Run("Absolute of integer", func(t *testing.T) {
		have := -10
		got := absolute(have)
		want := 10

		if got != want {
			t.Errorf("got {%v}, wanted {%v}", got, want)
		}
	})
}

func TestAtoiNoError(t *testing.T) {
	t.Run("Ascii to integer without errors", func(t *testing.T) {
		have := "10"
		got := atoiNoError(have)
		want := 10

		if got != want {
			t.Errorf("got {%v}, wanted {%v}", got, want)
		}
	})
}

func TestQuadrant(t *testing.T) {
	tests := []struct {
		name string
		loc  [2]int
		want int
	}{
		{
			name: "Quadrant 1 (Upper Left)",
			loc:  [2]int{1, 1},
			want: 1,
		},
		{
			name: "Quadrant 2 (Lower Left)",
			loc:  [2]int{1, 6},
			want: 2,
		},
		{
			name: "Quadrant 3 (Upper Right)",
			loc:  [2]int{6, 1},
			want: 3,
		},
		{
			name: "Quadrant 4 (Lower Right)",
			loc:  [2]int{6, 6},
			want: 4,
		},
		{
			name: "Line Void (On Quadrant Line)",
			loc:  [2]int{6, 3},
			want: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			haveX := test.loc[0]
			haveY := test.loc[1]
			got := quadrant(haveX, haveY)

			if got != test.want {
				t.Errorf("got {%v}, wanted {%v}", got, test.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	t.Run("Multiply number of robots per quadrant", func(t *testing.T) {
		have := map[int]int{
			0: 3,
			1: 1,
			2: 4,
			3: 3,
			4: 1,
		}
		got := solve(have)
		want := 12

		if got != want {
			t.Errorf("got {%v}, wanted {%v}", got, want)
		}
	})
}
