package day9

import (
	"log"
	"strconv"

	"github.com/wincus/adventofcode2021/internal/common"
)

type grid struct {
	field []int
	rows  int
	cols  int
}

// Solve returns the solutions for day 9
func Solve(s []string, p common.Part) int {

	var count int

	g := Parse(s)

	for i, v := range g.field {
		if g.IsLow(i) {
			count += v + 1
		}
	}
	return count
}

func Parse(s []string) grid {

	var g grid

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		g.rows++
		g.cols = 0
		for _, c := range line {
			g.cols++
			v, err := strconv.Atoi(string(c))
			if err != nil {
				log.Printf("could not parse int: %v", err)
				continue
			}
			g.field = append(g.field, v)
		}
	}

	return g
}

func (g grid) IsLow(i int) bool {

	ns := g.GetNeighbours(i)

	for _, n := range ns {
		if g.field[n] <= g.field[i] {
			return false
		}
	}

	return true

}

func (g grid) GetNeighbours(i int) []int {

	var neighbours []int

	// get row
	row := i / g.cols
	// get column
	col := i % g.cols

	// check left
	if col > 0 {
		neighbours = append(neighbours, i-1)
	}
	// check right
	if col < g.cols-1 {
		neighbours = append(neighbours, i+1)
	}
	// check top
	if row > 0 {
		neighbours = append(neighbours, i-g.cols)
	}
	// check bottom
	if row < g.rows-1 {
		neighbours = append(neighbours, i+g.cols)
	}

	return neighbours

}
