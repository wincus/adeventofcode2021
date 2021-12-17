package day9

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
				"2199943210",
				"3987894921",
				"9856789892",
				"8767896789",
				"9899965678",
			},
			p:    common.Part1,
			want: 15,
		},
		{
			input: []string{
				"2199943210",
				"3987894921",
				"9856789892",
				"8767896789",
				"9899965678",
			},
			p:    common.Part2,
			want: 1134,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
