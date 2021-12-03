package day3

import (
	"log"
	"strconv"

	"github.com/wincus/adventofcode2021/internal/common"
)

// Solve returns the solutions of Day3
func Solve(s []string, p common.Part) int {

	var counter int

	// diagnostic byte size
	n := len(s[0])

	gammaByte := make([]byte, n)
	epsilonByte := make([]byte, n)

	// assume that all entries have the same length
	for i := 0; i < n; i++ {

		counter = 0

		for _, v := range s {
			p, err := strconv.ParseInt(v, 2, 64)

			if err != nil {
				log.Panicf("could not convert binary: %v", err)
			}
			counter += (int(p) & (1 << i)) >> i
		}

		if counter > len(s)/2 {
			gammaByte[n-i-1] = '1'
			epsilonByte[n-i-1] = '0'
		} else {
			gammaByte[n-i-1] = '0'
			epsilonByte[n-i-1] = '1'
		}
	}

	gamma, err := strconv.ParseInt(string(gammaByte), 2, 64)

	if err != nil {
		log.Panicf("could not convert binary: %v", err)
	}

	epsilon, err := strconv.ParseInt(string(epsilonByte), 2, 64)

	if err != nil {
		log.Panicf("could not convert binary: %v", err)
	}

	return int(gamma) * int(epsilon)
}
