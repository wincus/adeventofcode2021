package day8

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/wincus/adventofcode2021/internal/common"
)

// errors
var DISPLAYPARSEERROR = errors.New("segments on the display are not valid")
var UNKNOWNSEGMENT = errors.New("unknown segment")
var CANNOTPARSEREADING = errors.New("cannot parse reading")

// segments of a display are represented as bits
// of a uint8 (8 bits) values, where:
//
//    aaaa
//   b    c
//   b    c
//    dddd       // 0  1  2  3  4  5  6  7
//   e    f      // ^  a  b  c  d  e  f  g
//   e    f      // |
//    gggg       // ( msb is not used )
//
//
// 0 -> 01110111 -> 0x77
// 1 -> 00010010 -> 0x12
// 2 -> 01011101 -> 0x5D
// 3 -> 01011011 -> 0x5B
// 4 -> 00111010 -> 0x3A
// 5 -> 01101011 -> 0x6B
// 6 -> 01101111 -> 0x6F
// 7 -> 01010010 -> 0x52
// 8 -> 01111111 -> 0x7F
// 9 -> 01111011 -> 0x7B
type display uint8

type reading struct {
	input   []display
	output  []display
	mapping func(r rune) rune
}

// maps a segment ID to a mask
var segmentToMask = map[rune]display{
	'a': 0x40,
	'b': 0x20,
	'c': 0x10,
	'd': 0x08,
	'e': 0x04,
	'f': 0x02,
	'g': 0x01,
}

// maps a segment combination to a number
var segmentToInt = map[display]uint8{
	0x77: 0,
	0x12: 1,
	0x5D: 2,
	0x5B: 3,
	0x3A: 4,
	0x6B: 5,
	0x6F: 6,
	0x52: 7,
	0x7F: 8,
	0x7B: 9,
}

// Solve returns the solutions for day 8
func Solve(s []string, p common.Part) int {

	var sum int

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		// part1 is trivial, we just count the number of characters
		// in the output
		if p == common.Part1 {
			v, err := CountUniq(line)

			if err != nil {
				log.Fatal(err)
			}

			sum += v
		}

		// part2 is far more complex, we need to find a mapping
		// that solves the segment combination.
		//
		// In this first implementation I just brute force it,
		// generating in runtime random combinations and checking
		// if the generated combination solves the problem.
		//
		// Personal TODO: improve this
		if p == common.Part2 {
			v, err := BruteForceLine(line)

			if err != nil {
				log.Fatal(err)
			}

			sum += v
		}
	}

	return sum
}

// count the number of numbers in the output
// with a uniq number of segments
func CountUniq(s string) (int, error) {
	var count int

	t := strings.Split(s, "|")

	if len(t) != 2 {
		return 0, fmt.Errorf("cannot parse input line")
	}

	q := strings.Split(t[1], " ")

	for _, c := range q {
		switch len(c) {
		case 0:
			continue
		case 2, 3, 4, 7:
			count++
		}
	}

	return count, nil
}

// BruteForceLine generates in runtime random combination of
// segments and checks if the combination solves the problem
// far from ideal but works, usually finds the solution in a
// couple of seconds.
func BruteForceLine(s string) (int, error) {

	var maxIterations = 1 << 16

	for {
		maxIterations--

		r, err := ParseReadingLine(s)

		if err != nil {
			log.Fatal(err)
		}

		// if check returns true we have found a mapping solution
		if r.Check() {
			return r.GetOutputReading()
		}

		if maxIterations == 0 {
			break
		}
	}

	return 0, fmt.Errorf("max iterations reached, could not find the answer")
}

func (r reading) Check() bool {

	for _, i := range r.input {
		_, err := i.GetNumber()

		if err != nil {
			return false
		}
	}

	return true

}

func (d *display) Set(c rune) error {

	mask, ok := segmentToMask[c]

	if !ok {
		return UNKNOWNSEGMENT
	}

	*d |= mask

	return nil
}

func (d *display) GetNumber() (uint8, error) {

	i, ok := segmentToInt[*d]

	if !ok {
		return 0, DISPLAYPARSEERROR
	}

	return i, nil

}

func Parse(s string) (display, error) {
	var d display

	for _, c := range s {
		if err := d.Set(c); err != nil {
			return d, err
		}
	}

	return d, nil
}

func ParseReadingLine(s string) (reading, error) {

	var r reading

	v := strings.Split(s, "|")

	if len(v) != 2 {
		return r, CANNOTPARSEREADING
	}

	input := strings.Split(v[0], " ")
	output := strings.Split(v[1], " ")

	r.mapping = GetMapping()

	for _, t := range input {

		if len(t) == 0 {
			continue
		}

		d, err := Parse(strings.Map(r.mapping, t))

		if err != nil {
			return r, err
		}

		r.input = append(r.input, d)
	}

	for _, t := range output {

		if len(t) == 0 {
			continue
		}

		d, err := Parse(strings.Map(r.mapping, t))

		if err != nil {
			return r, err
		}

		r.output = append(r.output, d)
	}

	return r, nil
}

func GetMapping() func(r rune) rune {
	rand.Seed(time.Now().UnixNano())

	a := []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g',
	}

	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	return func(r rune) rune {
		return a[r-'a']
	}
}

func (r *reading) GetOutputReading() (int, error) {

	var sum int

	l := len(r.output)

	for i := 0; i < l; i++ {

		v, err := r.output[i].GetNumber()

		if err != nil {
			return sum, err
		}

		sum += int(v) * int(math.Pow10(l-i-1))

	}

	return sum, nil

}
