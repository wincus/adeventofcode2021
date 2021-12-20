package day9

import (
	"log"
	"sort"
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
	var sizes []int

	g := Parse(s)

	for i, v := range g.field {
		// if i corresponds to a low point, get its value for part 1
		if g.IsLow(i) && p == common.Part1 {
			count += v + 1
		}

		// if i corresponds to a low point, get its basin size for part 2
		if g.IsLow(i) && p == common.Part2 {
			pi := new(int)
			g.getBasinSize(pi, i)
			sizes = append(sizes, *pi)
		}
	}

	if p == common.Part1 {
		return count
	}

	if p == common.Part2 {
		sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
		return sizes[0] * sizes[1] * sizes[2]
	}

	return 0

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

	ns := g.GetNeighbors(i)

	for _, n := range ns {
		if g.field[n] <= g.field[i] {
			return false
		}
	}

	return true

}

func (g grid) GetNeighbors(i int) []int {

	var neighbors []int

	// get row
	row := i / g.cols
	// get column
	col := i % g.cols

	// check left
	if col > 0 {
		neighbors = append(neighbors, i-1)
	}
	// check right
	if col < g.cols-1 {
		neighbors = append(neighbors, i+1)
	}
	// check top
	if row > 0 {
		neighbors = append(neighbors, i-g.cols)
	}
	// check bottom
	if row < g.rows-1 {
		neighbors = append(neighbors, i+g.cols)
	}

	return neighbors

}

// getBasinSize is a recursive function that returns the size of a basin
// starting with an initial low position.
// NOTE: this function is not thread safe.
// NOTE 2: this function makes changes on the field is processing, use it
// with a copy if you need to keep initial values.
func (g grid) getBasinSize(total *int, position int) {

	if g.field[position] != 9 {
		*total++              // counts current position
		g.field[position] = 9 // and avoids counting it again
	}

	n := g.GetNeighbors(position)

	for _, v := range n {

		if g.field[v] == 9 {
			continue
		}

		g.getBasinSize(total, v)
	}
}
