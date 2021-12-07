package day5

import (
	"fmt"
	"log"
	"math"
	"regexp"

	"github.com/wincus/adventofcode2021/internal/common"
)

type line struct {
	start coord
	end   coord
}

type coord struct {
	x int
	y int
}

const (
	SIZE = 1000
)

type field [SIZE][SIZE]int

const (
	REGEXP = `^(\d+),(\d+)\s+->\s+(\d+),(\d+)$`
)

// Solve returns the solutions for day 5
func Solve(s []string, p common.Part) int {

	var f field
	var allowDiagonal bool

	lines, err := Parse(s)

	if err != nil {
		log.Panic(err)
	}

	switch p {
	case common.Part1:
		allowDiagonal = false
	case common.Part2:
		allowDiagonal = true
	default:
		log.Panicf("unknown part: %d", p)
	}

	for _, l := range lines {
		f.connect(l, allowDiagonal)
	}

	return getAnomalyCount(f)
}

func Parse(s []string) ([]line, error) {

	var lines []line

	for _, l := range s {

		if len(l) == 0 {
			continue
		}

		re, err := regexp.Compile(REGEXP)

		if err != nil {
			return nil, err
		}

		result := re.FindStringSubmatch(l)

		if len(result) != 5 {
			return nil, fmt.Errorf("could not parse coordinates")
		}

		n, err := common.ToInt(result[1:])

		if err != nil {
			return nil, err
		}

		q := line{
			start: coord{
				x: n[0],
				y: n[1],
			},
			end: coord{
				x: n[2],
				y: n[3],
			},
		}

		lines = append(lines, q)

	}
	return lines, nil

}

// connect draws a line on a map
// TODO: this is unnecesarly complex, Im sure there is an easier
// way to do this.
func (f *field) connect(l line, allowDiagonal bool) {

	// vertical line ( fixed at x = c )
	if l.start.x == l.end.x {

		// swap if needed for iteration
		if l.start.y > l.end.y {
			l.start, l.end = l.end, l.start
		}

		for i := l.start.y; i <= l.end.y; i++ {
			(*f)[i][l.start.x]++
		}

		return
	}

	// horizontal line ( fixed at y = c )
	if l.start.y == l.end.y {

		// swap if needed for iteration
		if l.start.x > l.end.x {
			l.start, l.end = l.end, l.start
		}

		for i := l.start.x; i <= l.end.x; i++ {
			(*f)[l.start.y][i]++
		}

		return
	}

	// if no diagonal lines are allowed, return
	if !allowDiagonal {
		return
	}

	ratioX := l.end.x - l.start.x
	ratioY := l.end.y - l.start.y

	// if diagonal lines are allowed, ensure are 45 degree lines
	if math.Abs(float64(ratioX)) != math.Abs(float64(ratioY)) {
		return
	}

	// inverted diagonal
	if ratioX == ratioY {
		if l.start.x > l.end.x {
			l.start, l.end = l.end, l.start
		}

		if ratioX < 0 {
			ratioX = -ratioX
		}

		for i := 0; i <= ratioX; i++ {
			(*f)[l.start.y+i][l.start.x+i]++
		}

		return

	} else {
		if l.start.x > l.end.x {
			l.start, l.end = l.end, l.start
		}

		if ratioX < 0 {
			ratioX = -ratioX
		}

		for i := 0; i <= ratioX; i++ {
			(*f)[l.start.y-i][l.start.x+i]++
		}
	}
}

func getAnomalyCount(f field) int {

	var count int

	for _, row := range f {
		for _, v := range row {
			if v > 1 {
				count++
			}
		}
	}

	return count

}
