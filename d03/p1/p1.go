//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p1

import "strconv"

func P1(input []string) string {
	total := 0
	for _, l := range input {
		jolt := 0
		max := byte(0)
		pos := 0
		for i := 0; i < len(l)-1; i++ {
			if l[i] > max {
				max = l[i]
				pos = i
			}
		}
		jolt += int(max-'0') * 10
		max = 0
		for i := pos + 1; i < len(l); i++ {
			if l[i] > max {
				max = l[i]
			}
		}
		jolt += int(max - '0')
		total += jolt
	}
	return strconv.Itoa(total)
}
