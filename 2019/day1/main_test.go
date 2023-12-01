package day1

import "testing"

func TestCalculateFuel(t *testing.T) {
	testTable := []struct {
		name string
		have int
		want int
	}{
		{
			name: "one",
			have: 12,
			want: 2,
		},
		{
			name: "two",
			have: 14,
			want: 2,
		},
		{
			name: "three",
			have: 1969,
			want: 654,
		},
		{
			name: "four",
			have: 100756,
			want: 33583,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateFuel(tt.have)

			if tt.want != got {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestCalculateRecurseFuel(t *testing.T) {
	testTable := []struct {
		name string
		have int
		want int
	}{
		{
			name: "one",
			have: 14,
			want: 2,
		},
		{
			name: "two",
			have: 1969,
			want: 966,
		},
		{
			name: "three",
			have: 100756,
			want: 50346,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateRecurseFuel(tt.have, 0)

			if tt.want != got {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
