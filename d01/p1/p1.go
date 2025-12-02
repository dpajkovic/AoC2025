package p1

import (
	"fmt"
	"strconv"
)

func P1(input []string) int {
	count := 0
	pointer := 50

	for _, l := range input {
		move, err := transformDirection(l)
		if err != nil {
			panic(err)
		}
		pointer = (pointer + move + 100) % 100
		if pointer == 0 {
			count++
		}
	}

	return count
}

// transformDirection parses a string like "L32" or "R48" and returns the corresponding integer.
// "L" means negative, "R" means positive.
func transformDirection(s string) (int, error) {
	if len(s) < 2 {
		return 0, fmt.Errorf("invalid input: %s", s)
	}
	num, err := strconv.Atoi(s[1:])
	if err != nil {
		return 0, err
	}
	switch s[0] {
	case 'L':
		return -num, nil
	case 'R':
		return num, nil
	default:
		return 0, fmt.Errorf("invalid direction: %c", s[0])
	}
}
