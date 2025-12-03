//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p2

import "strconv"

func P2(input []string) string {
	total := 0
	for _, line := range input {
		jolt := 0
		for i := 0; i < 12; i++ {
			largest, rest := findLargestDigit(line, 11-i)
			jolt = jolt*10 + largest
			line = rest
		}
		total += jolt
	}
	return strconv.Itoa(total)
}

func findLargestDigit(line string, d int) (largest int, rest string) {
	largest = 0
	rest = line
	pos := 0
	for i := 0; i < len(line)-d; i++ {
		digit := int(line[i] - '0')
		if digit > largest {
			largest = digit
			pos = i
		}
	}
	rest = line[pos+1:]
	return largest, rest
}
