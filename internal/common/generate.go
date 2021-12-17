package common

import (
	"fmt"
	"html/template"
	"os"
)

const UTILS = `package day{{.Day}}

import (
	"github.com/wincus/adventofcode2021/internal/common"
)

// Solve returns the solutions for day {{.Day}}
func Solve(s []string, p common.Part) int {
	return 0
}
`

const TEST = `package day{{.Day}}

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
			input: []string{},
			p:     common.Part1,
			want:  0,
		},
		{
			input: []string{},
			p:     common.Part2,
			want:  0,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
`

type data struct {
	Day string
}

func Generate(day string) error {

	if day == "" {
		panic("day number is required")
	}

	u, err := template.New("utils").Parse(UTILS)

	if err != nil {
		return err
	}

	t, err := template.New("test").Parse(TEST)

	if err != nil {
		return err
	}

	d := fmt.Sprintf("internal/day%s", day)

	err = os.MkdirAll(d, 0755)

	if err != nil {
		return err
	}

	fUtils, _ := os.Create(fmt.Sprintf("%s/utils.go", d))
	defer fUtils.Close()

	err = u.Execute(fUtils, data{day})

	if err != nil {
		return err
	}

	fTest, _ := os.Create(fmt.Sprintf("%s/utils_test.go", d))
	defer fTest.Close()

	err = t.Execute(fTest, data{day})

	if err != nil {
		return err
	}

	return nil

}
