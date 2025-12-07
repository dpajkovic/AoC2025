//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p2

import (
	"strconv"
	"strings"
)

func P2(input []string) string {

	beams := make([]int, len(input[0]))

	for _, line := range input {
		if strings.Contains(line, "S") {
			beams[strings.IndexRune(line, 'S')] = 1
		}
		if strings.Contains(line, "^") {
			newBeams := make([]int, len(line))
			for x, c := range line {
				if c == '^' {
					newBeams[x-1] += beams[x]
					newBeams[x+1] += beams[x]
				} else {
					newBeams[x] += beams[x]
				}
			}
			beams = newBeams
		}
	}

	return strconv.Itoa(sumInts(beams))
}

func sumInts(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}
