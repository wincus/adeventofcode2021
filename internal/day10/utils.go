package day10

import (
	"errors"
	"sort"

	"github.com/wincus/adventofcode2021/internal/common"
)

type chunk struct {
	open   rune
	close  rune
	cvalue int
	ivalue int
}

var (
	P = chunk{open: '(', close: ')', ivalue: 1, cvalue: 3}
	S = chunk{open: '[', close: ']', ivalue: 2, cvalue: 57}
	C = chunk{open: '{', close: '}', ivalue: 3, cvalue: 1197}
	L = chunk{open: '<', close: '>', ivalue: 4, cvalue: 25137}
)

var CORRUPTED = errors.New("line is corrupted")
var INCOMPLETE = errors.New("line is incomplete")
var UNKNOWN = errors.New("unknown chunk type")

// Solve returns the solutions for day 10
func Solve(s []string, p common.Part) int {

	var countCorrupted int
	var IncompleteList []int

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		c, err := parseLine(line)

		if err == CORRUPTED {
			countCorrupted += c
		}

		if err == INCOMPLETE {
			IncompleteList = append(IncompleteList, c)
		}
	}

	switch p {
	case common.Part1:
		return countCorrupted
	case common.Part2:
		sort.Ints(IncompleteList)
		return IncompleteList[(len(IncompleteList)-1)/2]
	default:
		return 0
	}
}

func parseLine(s string) (int, error) {

	var stack []rune
	var countIncomplete int

	for _, c := range s {

		switch c {
		// start of a chunk
		case P.open, S.open, C.open, L.open:
			stack = append(stack, c)
		case P.close, S.close, C.close, L.close:

			// end of a chunk
			if len(stack) == 0 {
				return 0, UNKNOWN
			}

			// pop the last chunk
			ch := stack[len(stack)-1]

			u, err := getChunk(ch)

			if err != nil {
				return 0, err
			}

			// check if the value corresponds
			// to the closing match for tha last
			// value in the stack
			if u.close != c {
				v, err := getChunk(c)

				if err != nil {
					return 0, err
				}
				return v.cvalue, CORRUPTED
			}

			stack = stack[:len(stack)-1]

		default:
			return 0, UNKNOWN
		}
	}

	// if the stack is not empty after processing
	// all runes, then we have gotten an incomplete line
	if len(stack) > 0 {

		// get the missing tokens in reverse to
		// have the correct order
		for e := len(stack) - 1; e >= 0; e-- {
			u, err := getChunk(stack[e])

			if err != nil {
				return 0, err
			}

			countIncomplete *= 5
			countIncomplete += u.ivalue
		}

		return countIncomplete, INCOMPLETE
	}

	return 0, nil
}

func getChunk(r rune) (chunk, error) {
	switch r {
	case P.open, P.close:
		return P, nil
	case S.open, S.close:
		return S, nil
	case C.open, C.close:
		return C, nil
	case L.open, L.close:
		return L, nil
	default:
		return chunk{}, UNKNOWN
	}
}
