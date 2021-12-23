package day12

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
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d",
				"A-end",
				"b-end",
			},
			p:    common.Part1,
			want: 10,
		},
		{
			input: []string{
				"dc-end",
				"HN-start",
				"start-kj",
				"dc-start",
				"dc-HN",
				"LN-dc",
				"HN-end",
				"kj-sa",
				"kj-HN",
				"kj-dc",
			},
			p:    common.Part1,
			want: 19,
		},
		{
			input: []string{
				"fs-end",
				"he-DX",
				"fs-he",
				"start-DX",
				"pj-DX",
				"end-zg",
				"zg-sl",
				"zg-pj",
				"pj-he",
				"RW-he",
				"fs-DX",
				"pj-RW",
				"zg-RW",
				"start-pj",
				"he-WI",
				"zg-he",
				"pj-fs",
				"start-RW",
			},
			p:    common.Part1,
			want: 226,
		},
		{
			input: []string{
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d",
				"A-end",
				"b-end",
			},
			p:    common.Part2,
			want: 36,
		},
		{
			input: []string{
				"dc-end",
				"HN-start",
				"start-kj",
				"dc-start",
				"dc-HN",
				"LN-dc",
				"HN-end",
				"kj-sa",
				"kj-HN",
				"kj-dc",
			},
			p:    common.Part2,
			want: 103,
		},
		{
			input: []string{
				"fs-end",
				"he-DX",
				"fs-he",
				"start-DX",
				"pj-DX",
				"end-zg",
				"zg-sl",
				"zg-pj",
				"pj-he",
				"RW-he",
				"fs-DX",
				"pj-RW",
				"zg-RW",
				"start-pj",
				"he-WI",
				"zg-he",
				"pj-fs",
				"start-RW",
			},
			p:    common.Part2,
			want: 3509,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
