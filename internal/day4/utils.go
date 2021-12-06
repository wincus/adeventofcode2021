package day4

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/wincus/adventofcode2021/internal/common"
)

const (
	SIZE = 5
)

const (
	BOARDREGEXP = `^\s*(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)$`
)

type loc struct {
	value int
	gone  bool
}

type board map[int]loc

var index map[int][]int

// generate index for a board of size SIZE
func init() {
	// for a board of size 5:
	//  map[int][]int{
	// 	0: {0, 1, 2, 3, 4},
	// 	1: {5, 6, 7, 8, 9},
	// 	2: {10, 11, 12, 13, 14},
	// 	3: {15, 16, 17, 18, 19},
	// 	4: {20, 21, 22, 23, 24},
	// 	5: {0, 5, 10, 15, 20},
	// 	6: {1, 6, 11, 16, 21},
	// 	7: {2, 7, 12, 17, 22},
	// 	8: {3, 8, 13, 18, 23},
	// 	9: {4, 9, 14, 19, 24},
	// }

	index = make(map[int][]int)

	// rows
	for i := 0; i < SIZE; i++ {
		index[i] = make([]int, SIZE)
		for j := 0; j < SIZE; j++ {
			index[i][j] = i*SIZE + j
		}
	}

	// columns
	for i := SIZE; i < 2*SIZE; i++ {
		index[i] = make([]int, SIZE)
		for j := 0; j < SIZE; j++ {
			index[i][j] = (SIZE * j) + (i - SIZE)
		}
	}
}

// Solve returns the solutions for day 4
func Solve(s []string, p common.Part) int {

	n, err := ParseNumbers(s)

	if err != nil {
		log.Panicf("Error parsing input: %v", err)
	}

	boards, err := ParseBoards(s)

	if err != nil {
		log.Panicf("Error parsing input: %v", err)
	}

	winners := make(map[int]bool)

	for _, n := range n {
		for i, b := range boards {

			// play number n
			b.Set(n)

			// check if the board wins
			if b.Check() {

				// if part1 return the _first_ winner
				if p == common.Part1 {
					return n * b.SumNotGone()
				}

				// if part2 return the _last_ winner
				if p == common.Part2 {
					winners[i] = true

					if len(winners) == len(boards) {
						return n * b.SumNotGone()
					}
				}
			}
		}
	}

	return 0
}

func ParseBoards(s []string) ([]board, error) {

	boards := make([]board, 0)

	re, err := regexp.Compile(BOARDREGEXP)

	if err != nil {
		return nil, err
	}

	b := make(board)

	for _, line := range s {
		if len(line) == 0 {
			continue
		}

		result := re.FindStringSubmatch(line)

		if len(result) != 6 {
			continue
		}

		ints, err := common.ToInt(result[1:])

		if err != nil {
			return nil, err
		}

		start := len(b)

		for i, n := range ints {
			b[start+i] = loc{value: n, gone: false}
		}

		if len(b) >= SIZE*SIZE {
			boards = append(boards, b)
			b = make(board)
		}
	}

	return boards, nil

}

func ParseNumbers(s []string) ([]int, error) {

	for _, line := range s {
		if len(line) == 0 {
			continue
		}

		nString := strings.Split(line, ",")

		if len(nString) < 2 {
			continue
		}

		n, err := common.ToInt(nString)

		if err != nil {
			return nil, err
		}

		return n, nil
	}

	return nil, fmt.Errorf("No numbers found")

}

func (b *board) Set(v int) {
	for i := 0; i < SIZE*SIZE; i++ {
		if (*b)[i].value == v {
			(*b)[i] = loc{value: v, gone: true}
		}
	}
}

// Check checks if the has a row or
// a column solved
func (b *board) Check() bool {

	for i := 0; i < 2*SIZE; i++ {
		if b.checkRow(i) {
			return true
		}
	}

	return false
}

func (b *board) checkRow(i int) bool {
	for _, p := range index[i] {
		if !(*b)[p].gone {
			return false
		}
	}

	return true

}

func (b *board) SumNotGone() int {
	sum := 0
	for _, p := range *b {
		if !p.gone {
			sum += p.value
		}
	}
	return sum
}
