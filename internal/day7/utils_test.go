package day7

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
				"16,1,2,0,4,2,7,1,2,14",
			},
			p:    common.Part1,
			want: 37,
		},
		{
			input: []string{
				"16,1,2,0,4,2,7,1,2,14",
			},
			p:    common.Part2,
			want: 168,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
