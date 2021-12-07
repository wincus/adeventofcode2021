package day5

import (
	"fmt"
	"log"
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
	SIZE = 10
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

func (f *field) connect(l line, allowDiagonal bool) {

	// swap coordinates if needed
	if l.start.x > l.end.x || l.start.y > l.end.y {
		l.start, l.end = l.end, l.start
	}

	// ignore non horizontal or vertical lines
	if l.start.x != l.end.x && l.start.y != l.end.y {

		if !allowDiagonal {
			return
		}

		// 45 degree diagonals are the only ones supported
		if (l.end.x - l.start.x) != (l.start.y + l.end.y) {
			return
		}
	}

	for x := l.start.y; x <= l.end.y; x++ {
		for y := l.start.x; y <= l.end.x; y++ {
			(*f)[x][y]++
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
