package day12

import (
	"log"
	"strings"

	"github.com/wincus/adventofcode2021/internal/common"
)

type caveKind int

const (
	small caveKind = iota
	large
)

type cave struct {
	name    string
	kind    caveKind
	connect []*cave
}

type caves map[string]*cave

// Solve returns the solutions for day 12
func Solve(s []string, p common.Part) int {

	// parse data
	c := Parse(s)

	return c.walk([]string{"start"}, p)

}

func (c caves) walk(s []string, p common.Part) int {

	var sum int

	last := s[len(s)-1]

	if last == "end" {
		return 1
	}

	for _, n := range c[last].connect {

		// checks if a given cave can be visited
		if !c.check(n.name, s, p) {
			continue
		}

		r := make([]string, len(s))
		copy(r, s)
		r = append(r, n.name)

		sum += c.walk(r, p)
	}

	return sum
}

// checks if a given cave can be visited following the rules of
// part 1 and part 2
func (c caves) check(f string, s []string, p common.Part) bool {

	// large caves are always accessible
	if c[f].kind == large {
		return true
	}

	switch p {
	case common.Part1:
		// for part 1 small caves can only be accessible once
		return !common.Contains(f, s)
	case common.Part2:

		// for part 2 ... start and end are exceptions and
		// can only be accessed once
		if f == "start" || f == "end" {
			return !common.Contains(f, s)
		}

		// allow a small cave to be visited once
		if !common.Contains(f, s) {
			return true
		}

		var count = make(map[string]int)

		// for the rest of small caves, only 1 can be accessed twice
		// while the rest can only be accessed once
		for _, d := range s {
			if c[d].kind == small {
				count[d]++
			}
		}

		for _, v := range count {
			if v > 1 {
				return false
			}
		}

		return true

	default:
		return false
	}
}

func Parse(s []string) caves {

	var f = make(caves)

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		c := strings.Split(line, "-")

		if len(c) != 2 {
			log.Printf("invalid line: %s", line)
			continue
		}

		if _, ok := f[c[0]]; !ok {
			if strings.ToUpper(c[0]) == c[0] {
				f[c[0]] = &cave{name: c[0], kind: large}
			} else {
				f[c[0]] = &cave{name: c[0], kind: small}
			}
		}

		if _, ok := f[c[1]]; !ok {
			if strings.ToUpper(c[1]) == c[1] {
				f[c[1]] = &cave{name: c[1], kind: large}
			} else {
				f[c[1]] = &cave{name: c[1], kind: small}
			}
		}

		f[c[0]].connect = append(f[c[0]].connect, f[c[1]])
		f[c[1]].connect = append(f[c[1]].connect, f[c[0]])
	}

	return f
}
