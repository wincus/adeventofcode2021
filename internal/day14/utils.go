package day14

import (
	"fmt"
	"log"
	"strings"

	"github.com/wincus/adventofcode2021/internal/common"
)

type pairs map[string]int
type rules map[string]string
type stats map[string]int

// Solve returns the solutions for day 14
func Solve(s []string, p common.Part) int {

	var iterations int
	var a pairs
	var r rules
	var t stats

	a = getPairs(s[0])

	r = getRules(s)

	switch p {
	case common.Part1:
		iterations = 10
	case common.Part2:
		iterations = 40
	default:
		log.Fatal("Invalid part")
	}

	for i := 0; i < iterations; i++ {
		a = process(a, r)
	}

	t = getStats(a)

	return getMax(t) - getMin(t)
}

func getMin(s stats) int {
	var min int

	for _, v := range s {
		if min == 0 || v < min {
			min = v
		}
	}

	return min
}

func getMax(s stats) int {
	var max int

	for _, v := range s {
		if v > max {
			max = v
		}
	}

	return max
}

func getStats(a pairs) stats {

	s := make(stats)

	for k, v := range a {
		s[string(k[0])] += v
		s[string(k[1])] += v
	}

	t := make(stats)

	for k, v := range s {
		if v%2 == 0 {
			t[k] = v / 2
		} else {
			t[k] = ((v - 1) / 2) + 1
		}
	}

	return t
}

func process(a pairs, r rules) pairs {

	n := make(pairs)

	for k, v := range a {

		insert, ok := r[k]

		if !ok {
			log.Panicf("No rule for %s", k)
		}

		n[fmt.Sprintf("%s%s", string(k[0]), string(insert))] += v
		n[fmt.Sprintf("%s%s", string(insert), string(k[1]))] += v

	}

	return n

}

func getPairs(s string) pairs {

	p := make(pairs)

	for i := 0; i < len(s)-1; i++ {
		p[fmt.Sprintf("%c%c", s[i], s[i+1])]++
	}
	return p
}

func getRules(s []string) rules {
	r := make(rules)

	for _, v := range s {
		if v == "" {
			continue
		}

		a := strings.Split(v, " -> ")

		if len(a) != 2 {
			continue
		}

		r[a[0]] = a[1]

	}
	return r
}
