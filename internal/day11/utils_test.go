package day11

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
				"5483143223",
				"2745854711",
				"5264556173",
				"6141336146",
				"6357385478",
				"4167524645",
				"2176841721",
				"6882881134",
				"4846848554",
				"5283751526",
			},
			p:    common.Part1,
			want: 1656,
		},
		{
			input: []string{
				"5483143223",
				"2745854711",
				"5264556173",
				"6141336146",
				"6357385478",
				"4167524645",
				"2176841721",
				"6882881134",
				"4846848554",
				"5283751526",
			},
			p:    common.Part2,
			want: 195,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
