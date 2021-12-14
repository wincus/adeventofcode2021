package day7

import (
	"log"
	"math"
	"strings"

	"github.com/wincus/adventofcode2021/internal/common"
)

type group struct {
	min       int
	max       int
	positions []int
	atIndex   int
	cost      int
}

// Solve returns the solutions for day 7
func Solve(s []string, p common.Part) int {

	var nonlinear bool

	n, err := common.ToInt(strings.Split(s[0], ","))

	if err != nil {
		log.Panicf("could not convert input to int: %v", err)
	}

	g := NewGroup(n)

	if p == common.Part1 {
		nonlinear = false
	} else if p == common.Part2 {
		nonlinear = true
	} else {
		log.Panicf("unknown part: %v", p)
	}

	// for all possible positions, calculate the cost
	for i := g.min; i < g.max; i++ {
		g.calculateMinCost(i, nonlinear)
	}
	return g.cost
}

func NewGroup(s []int) *group {

	return &group{
		positions: s,
		min:       GetMin(s),
		max:       GetMax(s),
	}
}

func (g *group) calculateMinCost(p int, nonlinear bool) {

	var cost int
	var f, k int

	for _, v := range g.positions {

		if nonlinear {
			f = int(math.Abs(float64(v - p)))
			k = (f*f + f) / 2
		} else {
			k = int(math.Abs(float64(v - p)))
		}

		cost += k

		// abort if cost is bigger than recorded
		if g.cost != 0 && cost > g.cost {
			return
		}
	}

	g.cost = cost
	g.atIndex = p
}

func GetMin(v []int) int {

	min := math.MaxInt32

	for _, i := range v {
		if i < min {
			min = i
		}
	}
	return min

}

func GetMax(v []int) int {

	var max int

	for _, i := range v {
		if i > max {
			max = i
		}
	}
	return max

}
