//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p2

import (
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

type intRange struct {
	min int
	max int
}

func P2(input []string) int {
	r := splitIntoGroups(input[0])
	sum := 0
	for _, v := range r {
		for i := v.min; i <= v.max; i++ {
			if isRepeating(i) {
				sum += i
			}
		}
	}
	return sum
}

func splitIntoGroups(input string) []intRange {
	if input == "" {
		return nil
	}

	parts := strings.Split(input, ",")
	out := make([]intRange, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		bounds := strings.Split(p, "-")
		if len(bounds) != 2 {
			continue
		}
		a, err1 := strconv.Atoi(bounds[0])
		b, err2 := strconv.Atoi(bounds[1])
		if err1 != nil || err2 != nil {
			continue
		}
		out = append(out, intRange{min: a, max: b})
	}
	return out
}

func isRepeating(i int) bool {
	s := strconv.Itoa(i)
	re := regexp2.MustCompile(`^(\d+)\1+$`, 0)
	matches, _ := re.MatchString(s)
	return matches
}
