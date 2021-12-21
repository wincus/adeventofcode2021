package day11

import (
	"log"
	"strconv"

	"github.com/wincus/adventofcode2021/internal/common"
)

type octopus struct {
	v       int
	flashed bool
}

type grid struct {
	field []octopus
	rows  int
	cols  int
}

// Solve returns the solutions for day 11
func Solve(s []string, p common.Part) int {

	var count int
	var i int

	g := Parse(s)

	for {
		i++
		count += g.step()

		if p == common.Part1 && i >= 100 {
			return count
		}

		if p == common.Part2 && g.isSyncronized() {
			return i
		}
	}
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
			g.field = append(g.field, octopus{v, false})
		}
	}

	return g
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
	// check top left
	if row > 0 && col > 0 {
		neighbors = append(neighbors, i-g.cols-1)
	}
	// check top right
	if row > 0 && col < g.cols-1 {
		neighbors = append(neighbors, i-g.cols+1)
	}
	// check bottom left
	if row < g.rows-1 && col > 0 {
		neighbors = append(neighbors, i+g.cols-1)
	}
	// check bottom right
	if row < g.rows-1 && col < g.cols-1 {
		neighbors = append(neighbors, i+g.cols+1)
	}

	return neighbors

}

// step simulates a step in the grid, returns the number of
// flashes that have occurred
func (g grid) step() int {

	var count int

	// reset all once step has finished
	defer g.reset()

	// increase energy level by 1 for each octopus
	for i := 0; i < len(g.field); i++ {
		g.field[i].v++
	}

	for {
		// check if the grid is sable ( aka no octopuses
		// are left to flash )
		if g.isStable() {
			return count
		}

		for i := 0; i < len(g.field); i++ {

			// if energy level is bigger than 9 a flash occurs
			if g.field[i].v > 9 {
				count++
				g.field[i].flashed = true
				g.field[i].v = 0

				// a flash is propagated to its neighbors
				n := g.GetNeighbors(i)

				for _, n := range n {
					// increase neighbor energy level by 1
					// only if hasn't flashed on this step
					if !g.field[n].flashed {
						g.field[n].v++
					}
				}
			}
		}
	}
}

// isStable returns false if the grid is not stable
// aka some octopuses are pending to flash
func (g grid) isStable() bool {

	for i := 0; i < len(g.field); i++ {
		if g.field[i].v > 9 {
			return false
		}
	}

	return true
}

// isSyncronized returns true if the grid is syncronized
// aka all octopuses have the same energy level == 0
func (g grid) isSyncronized() bool {
	for i := 0; i < len(g.field); i++ {
		if g.field[i].v != 0 {
			return false
		}
	}

	return true
}

// reset resets all octopuses to their initial state
func (g grid) reset() {
	for i := 0; i < len(g.field); i++ {
		g.field[i].flashed = false
	}
}
