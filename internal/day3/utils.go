package day3

import (
	"fmt"
	"log"
	"strconv"

	"github.com/wincus/adventofcode2021/internal/common"
)

type popularity int

const (
	ZEROISTHEMOSTCOMMON popularity = iota
	ONEISTHEMOSTCOMMON
	EQUAL
)

type rate int

const (
	GAMMA rate = iota
	EPSILON
	OXYGEN
	CO2
)

// Solve returns the solutions for day 3
func Solve(s []string, p common.Part) int {

	n := len(s[0])

	r, err := common.BinaryToDecimal(s)

	if err != nil {
		log.Fatal(err)
	}

	if p == common.Part1 {
		gamma, err := getRate(r, n, GAMMA)

		if err != nil {
			log.Fatal(err)
		}

		epsilon, err := getRate(r, n, EPSILON)

		if err != nil {
			log.Fatal(err)
		}

		return gamma * epsilon
	}

	if p == common.Part2 {

		oxygen, err := getRate(r, n, OXYGEN)

		if err != nil {
			log.Fatal(err)
		}

		co2, err := getRate(r, n, CO2)

		if err != nil {
			log.Fatal(err)
		}

		return oxygen * co2
	}

	log.Panicf("Unknown Part %v ", p)
	return 0
}

func getRate(values []int, n int, r rate) (int, error) {

	var result string

	for i := n - 1; i >= 0; i-- {

		// get the most common bit at index i
		bit := getMostCommonBit(values, i)

		switch r {
		case GAMMA:
			result = fmt.Sprintf("%v%v", result, bit)
		case EPSILON:
			result = fmt.Sprintf("%v%v", result, bit.Flip())
		case OXYGEN:
			if bit == EQUAL {
				values = filterByBit(values, i, 1)
			} else {
				values = filterByBit(values, i, int(bit))
			}
		case CO2:
			if bit == EQUAL {
				values = filterByBit(values, i, 0)
			} else {
				values = filterByBit(values, i, int(bit.Flip()))
			}
		default:
			return 0, fmt.Errorf("Unknown rate %v", r)
		}

		if r == OXYGEN || r == CO2 {
			if len(values) == 1 {
				return values[0], nil
			}
		}
	}

	// if we have gotten this far we have not found the rate
	// for OXYGEN or CO2
	if r == OXYGEN || r == CO2 {
		return 0, fmt.Errorf("No value found for rate %v", r)
	}

	q, err := strconv.ParseInt(result, 2, 64)

	if err != nil {
		return 0, err
	}

	return int(q), nil

}

// filterByBit removes from a slice of ints all elements where the the bit *not*
// equal to the given bit at the given index.
func filterByBit(r []int, index, bit int) []int {

	var new []int

	for _, v := range r {
		if getBitByIndex(v, index) == bit {
			new = append(new, v)
		}
	}

	return new
}

// for a slice of ints s return which bit is most popular at a given index i
func getMostCommonBit(s []int, i int) popularity {

	var counter int

	for _, v := range s {
		counter += getBitByIndex(v, i)
	}

	if 2*counter > len(s) {
		return ONEISTHEMOSTCOMMON
	}

	if 2*counter < len(s) {
		return ZEROISTHEMOSTCOMMON
	}

	return EQUAL
}

// getBitByIndex returns the bit value at the given index
func getBitByIndex(v, i int) int {
	return (v & (1 << i)) >> i
}

func (p popularity) Flip() popularity {
	switch p {
	case ZEROISTHEMOSTCOMMON:
		return ONEISTHEMOSTCOMMON
	case ONEISTHEMOSTCOMMON:
		return ZEROISTHEMOSTCOMMON
	default:
		return p
	}
}
