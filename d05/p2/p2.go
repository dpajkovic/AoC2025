//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p2

import (
	"fmt"
	"strconv"
)

type numRange struct {
	min    int64
	max    int64
	active bool
}

func P2(input []string) string {
	ranges := make([]numRange, 0)

	for _, line := range input {
		if line == "" {
			break
		}
		var r numRange
		_, err := fmt.Sscanf(line, "%d-%d", &r.min, &r.max)
		if err != nil {
			panic(err)
		}
		r.active = true
		ranges = append(ranges, r)
	}

	mergedRanges := mergeRanges(ranges)

	count := int64(0)
	for _, r := range mergedRanges {
		count += r.max - r.min + 1
	}

	return strconv.FormatInt(count, 10)
}

func mergeRanges(ranges []numRange) []numRange {
	merged := make([]numRange, 0)
	changed := false
	for i := 0; i < len(ranges); i++ {
		if !ranges[i].active {
			continue
		}
		current := ranges[i]
		for j := i + 1; j < len(ranges); j++ {
			if !ranges[j].active {
				continue
			}
			if isInRange(ranges[j].min, current) || isInRange(ranges[j].max, current) ||
				isInRange(current.min, ranges[j]) || isInRange(current.max, ranges[j]) {
				if ranges[j].min < current.min {
					current.min = ranges[j].min
				}
				if ranges[j].max > current.max {
					current.max = ranges[j].max
				}
				ranges[j].active = false
				changed = true
			}
		}
		merged = append(merged, current)

	}
	if changed {
		return mergeRanges(merged)
	}
	return merged
}

func isInRange(n int64, r numRange) bool {
	return n >= r.min && n <= r.max
}
