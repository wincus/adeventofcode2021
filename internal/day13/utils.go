package day13

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/wincus/adventofcode2021/internal/common"
)

const (
	COORD = `^(\d+),(\d+)$`
	FOLD  = `^fold\s+along\s+([xy])=(\d+)$`
)

type paper []point

type point struct {
	x int
	y int
}

type axis int

const (
	foldAlongX axis = iota
	foldAlongY
)

type instruction struct {
	a axis
	n int
}

// Solve returns the solutions for day 13
func Solve(s []string, p common.Part) int {

	f, i, err := Parse(s)

	if err != nil {
		log.Printf("error parsing input: %v", err)
		return 0
	}

	// Part 1, fold using only first instruction
	if p == common.Part1 {
		f = f.fold(i[0])
		return f.count()
	}

	// Part 2, fold using all instructions and
	// print the result to get the code
	if p == common.Part2 {
		for _, inst := range i {
			f = f.fold(inst)
		}

		f.print()

		return f.count()

	}

	return 0
}

func Parse(s []string) (paper, []instruction, error) {

	var i []instruction
	var p paper

	reCoord, err := regexp.Compile(COORD)

	if err != nil {
		return nil, nil, err
	}

	reFold, err := regexp.Compile(FOLD)

	if err != nil {
		return nil, nil, err
	}

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		if reCoord.MatchString(line) {
			c := reCoord.FindStringSubmatch(line)

			x, err := strconv.Atoi(c[1])

			if err != nil {
				return nil, nil, err
			}

			y, err := strconv.Atoi(c[2])

			if err != nil {
				return nil, nil, err
			}

			p = append(p, point{x, y})
		}

		if reFold.MatchString(line) {
			f := reFold.FindStringSubmatch(line)

			a, err := getAxis(f[1])

			if err != nil {
				return nil, nil, err
			}

			n, err := strconv.Atoi(f[2])

			if err != nil {
				return nil, nil, err
			}

			i = append(i, instruction{a, n})

		}
	}
	return p, i, nil
}

func getAxis(s string) (axis, error) {
	switch s {
	case "x":
		return foldAlongX, nil
	case "y":
		return foldAlongY, nil
	default:
		return 0, fmt.Errorf("invalid axis: %s", s)
	}
}

func (p *paper) count() int {

	m := make(map[point]bool)

	for _, t := range *p {
		m[t] = true
	}

	return len(m)
}

func (p *paper) fold(i instruction) paper {

	var newPaper paper

	switch i.a {

	case foldAlongX:
		for _, t := range *p {

			if t.x < i.n {
				newPaper = append(newPaper, t)
			}

			if t.x > i.n {
				newPaper = append(newPaper, point{2*i.n - t.x, t.y})
			}
		}

	case foldAlongY:
		for _, t := range *p {

			if t.y < i.n {
				newPaper = append(newPaper, t)
			}

			if t.y > i.n {
				newPaper = append(newPaper, point{t.x, 2*i.n - t.y})
			}
		}
	}

	return newPaper
}

func (p paper) print() {

	minX := 0
	maxX := p[0].x
	minY := 0
	maxY := p[0].y

	for _, t := range p {
		if t.x > maxX {
			maxX = t.x
		}
		if t.y > maxY {
			maxY = t.y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			fmt.Printf("%s", p.contains(point{x, y}))
		}
		fmt.Println()
	}
}

func (p paper) contains(t point) string {
	for _, v := range p {
		if v == t {
			return "#"
		}
	}
	return "."
}
