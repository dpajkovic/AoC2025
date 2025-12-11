//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p1

import (
	"strconv"
	"strings"
)

func depthFirstSearch(paths map[string][]string, current string, memory map[string]int, visited map[string]bool) int {
	if current == "out" {
		return 1
	}

	if val, exists := memory[current]; exists {
		return val
	}

	if visited[current] {
		return 0
	}

	visited[current] = true
	var ans int
	for _, child := range paths[current] {
		ans += depthFirstSearch(paths, child, memory, visited)
	}
	visited[current] = false

	memory[current] = ans
	return ans
}

func P1(input []string) string {
	paths := make(map[string][]string)
	for _, line := range input {
		machines := strings.Fields(line)
		key := machines[0][:len(machines[0])-1]
		paths[key] = machines[1:]
	}
	memory := make(map[string]int, len(paths))
	visited := make(map[string]bool, len(paths))
	result := depthFirstSearch(paths, "you", memory, visited)
	return strconv.Itoa(result)
}
