package day1

import (
	"log"

	"github.com/wincus/adventofcode2021/internal/common"
)

// Solve returns the solutions of Day1
func Solve(s []int, p common.Part) int {

	var previous int
	var sum int
	var counter int
	var n int

	switch p {
	case common.Part1:
		n = 1
	case common.Part2:
		n = 3
	default:
		log.Panicf("Invalid part: %v", p)
	}

	for i := 0; i < len(s)-n+1; i++ {

		sum = 0

		for e := 0; e < n; e++ {
			sum += s[i+e]
		}

		if sum > previous && i > 0 {
			counter++
		}

		previous = sum

	}

	return counter

}
