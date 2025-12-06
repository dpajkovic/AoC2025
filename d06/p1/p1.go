//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p1

import (
	"strconv"
	"strings"
)

func P1(input []string) string {
	var nums [][]int
	var operators []string

	for i, line := range input {
		if i == len(input)-1 {
			operators = strings.Fields(line)
			break
		}
		var row []int
		numsInLine := strings.Fields(line)
		for _, n := range numsInLine {
			num, _ := strconv.Atoi(n)
			row = append(row, num)
		}
		nums = append(nums, row)
	}

	total := 0
	for i, op := range operators {
		var col []int
		for j := 0; j < len(nums); j++ {
			col = append(col, nums[j][i])
		}
		if op == "+" {
			total += sum(col)
		} else if op == "*" {
			total += product(col)
		}
	}
	return strconv.Itoa(total)
}

func sum(s []int) int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}

func product(s []int) int {
	total := 1
	for _, v := range s {
		total *= v
	}
	return total
}
