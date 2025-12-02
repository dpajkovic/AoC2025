package p2

import (
	"fmt"
	"strconv"
)

func P2(input []string) int {
	count := 0
	pointer := 50

	for _, l := range input {
		move, err := transformDirection(l)
		if err != nil {
			panic(err)
		}
		step := intSign(move)
		for i := 0; i < intAbs(move); i++ {
			pointer = (pointer + step + 100) % 100
			if pointer == 0 {
				count++
			}
		}
	}

	return count
}

func intSign(n int) int {
	if n < 0 {
		return -1
	}
	return 1
}

func intAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

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
