package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/wincus/adventofcode2021/internal/common"
)

type command int

const (
	FORWARD command = iota
	UP
	DOWN
)

type instruction struct {
	c command
	v int
}

type position struct {
	horizontal int
	dept       int
	aim        int
}

// Solve returns the solutions of Day2
func Solve(s []string, p common.Part) int {

	var pos position

	for _, v := range s {

		if len(v) == 0 {
			continue
		}

		inst, err := parse(v)

		if err != nil {
			log.Panicf("could not parse the instruction: %v", err)
		}

		if err := runInstruction(inst, &pos, p); err != nil {
			log.Panicf("could not run the instruction: %v", err)
		}
	}

	return pos.dept * pos.horizontal

}

func runInstruction(i instruction, pos *position, p common.Part) error {

	if p != common.Part1 && p != common.Part2 {
		return fmt.Errorf("invalid part: %v", p)
	}

	// logic for Part1
	if p == common.Part1 {
		switch i.c {
		case FORWARD:
			pos.horizontal += i.v
		case UP:
			pos.dept -= i.v
		case DOWN:
			pos.dept += i.v
		default:
			return fmt.Errorf("unknown command %d", i.c)
		}
	}

	// logic for Part2
	if p == common.Part2 {
		switch i.c {
		case DOWN:
			pos.aim += i.v
		case UP:
			pos.aim -= i.v
		case FORWARD:
			pos.horizontal += i.v
			pos.dept += pos.aim * i.v
		default:
			return fmt.Errorf("unknown command %d", i.c)
		}
	}
	return nil
}

func parse(s string) (instruction, error) {

	var i instruction

	if len(s) == 0 {
		return i, fmt.Errorf("empty string")
	}

	v := strings.Split(s, " ")

	if len(v) != 2 {
		return i, fmt.Errorf("invalid instruction: %s", s)
	}

	switch v[0] {
	case "forward":
		i.c = FORWARD
	case "down":
		i.c = DOWN
	case "up":
		i.c = UP
	default:
		return i, fmt.Errorf("invalid command: %s", v[0])
	}

	n, err := strconv.Atoi(v[1])

	if err != nil {
		return i, fmt.Errorf("invalid value: %s", v[1])
	}

	i.v = n

	return i, nil
}
