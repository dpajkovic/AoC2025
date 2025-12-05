//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p1

import (
	"fmt"
	"strconv"
)

type numRange struct {
	min int64
	max int64
}

func P1(input []string) string {
	ranges := make([]numRange, 0)
	var nums []int64
	inRanges := true
	for _, line := range input {
		if line == "" {
			inRanges = false
			continue
		}
		if inRanges {
			var r numRange
			_, err := fmt.Sscanf(line, "%d-%d", &r.min, &r.max)
			if err != nil {
				panic(err)
			}
			ranges = append(ranges, r)
			continue
		}
		n, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	count := 0
	for _, n := range nums {
		for _, r := range ranges {
			if isInRange(n, r) {
				count++
				break
			}
		}
	}

	return strconv.Itoa(count)
}

func isInRange(n int64, r numRange) bool {
	return n >= r.min && n <= r.max
}
