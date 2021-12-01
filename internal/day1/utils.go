package day1

import (
	"strconv"

	"github.com/wincus/adventofcode2021/internal/common"
)

// Solve returns the solutions of Day1
func Solve(s []string, p common.Part) int {

	var previous int
	var sum int
	var counter int
	var n int

	if p == common.Part1 {
		n = 1
	}

	if p == common.Part2 {
		n = 3
	}

	for i := 0; i < len(s)-n+1; i++ {
		for e := 0; e < n; e++ {
			d, _ := strconv.Atoi(s[i+e])
			sum += d
		}

		if sum > previous && i > 0 {
			counter++
		}

		previous = sum
		sum = 0

	}

	return counter

}
