//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p2

import "strconv"

func P2(input []string) string {
	total := 0
	newCol := []int{}

	for x := len(input[0]) - 1; x >= 0; x-- {
		n := 0
		operated := false
		for y := 0; y < len(input); y++ {
			if input[y][x] == ' ' {
				continue
			} else if input[y][x] == '+' {
				result := n
				for i := 0; i < len(newCol); i++ {
					result += newCol[i]
				}
				total += result
				operated = true
				newCol = []int{}
				x--
				break
			} else if input[y][x] == '*' {
				result := n
				for i := 0; i < len(newCol); i++ {
					result *= newCol[i]
				}
				total += result
				operated = true
				newCol = []int{}
				x--
				break
			} else {
				fig, _ := strconv.Atoi(string(input[y][x]))
				n = n*10 + fig
			}
		}
		if !operated {
			newCol = append(newCol, n)
		}
	}

	return strconv.Itoa(total)
}
