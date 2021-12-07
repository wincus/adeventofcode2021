package day5

import (
	"testing"

	"github.com/wincus/adventofcode2021/internal/common"
)

type Test struct {
	input []string
	p     common.Part
	want  int
}

func TestSolver(t *testing.T) {

	tests := []Test{
		{
			input: []string{
				"0,9 -> 5,9", // no changes
				"8,0 -> 0,8", // ?????
				"9,4 -> 3,4", // 3,4 -> 9,4
				"2,2 -> 2,1", // 2,1 -> 2,2
				"7,0 -> 7,4", // no changes
				"6,4 -> 2,0", // ignored
				"0,9 -> 2,9", // no changes
				"3,4 -> 1,4", // 1,4 -> 3,4
				"0,0 -> 8,8", // no changes
				"5,5 -> 8,2", // ignored
			},
			p:    common.Part1,
			want: 5,
		},
		{
			input: []string{
				"0,9 -> 5,9",
				"8,0 -> 0,8",
				"9,4 -> 3,4",
				"2,2 -> 2,1",
				"7,0 -> 7,4",
				"6,4 -> 2,0",
				"0,9 -> 2,9",
				"3,4 -> 1,4",
				"0,0 -> 8,8",
				"5,5 -> 8,2",
			},
			p:    common.Part2,
			want: 12,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
