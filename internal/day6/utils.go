package day6

import (
	"log"
	"strings"

	"github.com/wincus/adventofcode2021/internal/common"
)

type lanternfish int

type school map[int]int

// Solve returns the solutions for day 6
func Solve(s []string, p common.Part) int {

	var days int

	switch p {
	case common.Part1:
		days = 80
	case common.Part2:
		days = 256
	default:
		log.Fatal("Unknown part")
	}

	n, err := common.ToInt(strings.Split(s[0], ","))

	if err != nil {
		log.Panicf("could not convert input to int: %v", err)
	}

	g := make(school)

	// initial values
	for _, i := range n {
		g[i]++
	}

	for i := 0; i < days; i++ {
		EvolveOneDay(g)
	}

	var sum int

	for _, v := range g {
		sum += v
	}

	return sum
}

func EvolveOneDay(g school) {

	gZero := g[0]

	g[0] = g[1]
	g[1] = g[2]
	g[2] = g[3]
	g[3] = g[4]
	g[4] = g[5]
	g[5] = g[6]
	g[6] = g[7] + gZero
	g[7] = g[8]
	g[8] = gZero

}
