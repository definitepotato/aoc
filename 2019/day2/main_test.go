package day2

import (
	"testing"
)

func TestProcessOpcodes(t *testing.T) {
	testTable := []struct {
		name string
		have string
		want string
	}{
		{
			name: "one",
			have: "1,0,0,0,99",
			want: "2,0,0,0,99",
		},
		{
			name: "two",
			have: "2,3,0,3,99",
			want: "2,3,0,6,99",
		},
		{
			name: "three",
			have: "2,4,4,5,99,0",
			want: "2,4,4,5,99,9801",
		},
		{
			name: "four",
			have: "1,1,1,4,99,5,6,0,99",
			want: "30,1,1,4,2,5,6,0,99",
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			got := ProcessOpcodes(tt.have, 0)
			if got != tt.want {
				t.Errorf("got %s, want %s\n", got, tt.want)
			}
		})
	}
}
