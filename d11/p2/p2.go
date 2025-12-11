//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p2

import (
	"strconv"
	"strings"
)

type state struct {
	node       string
	visitedFFT bool
	visitedDAC bool
}

func depthFirstSearch(paths map[string][]string, current string, visitedFFT, visitedDAC bool, memory map[state]int) int {
	if current == "out" {
		if visitedFFT && visitedDAC {
			return 1
		}
		return 0
	}
	s := state{current, visitedFFT, visitedDAC}
	if val, exists := memory[s]; exists {
		return val
	}
	result := 0
	for _, child := range paths[current] {
		result += depthFirstSearch(paths, child, visitedFFT || child == "fft", visitedDAC || child == "dac", memory)
	}
	memory[s] = result
	return result
}

func P2(input []string) string {
	paths := make(map[string][]string)
	for _, line := range input {
		machines := strings.Fields(line)
		key := machines[0][:len(machines[0])-1]
		paths[key] = machines[1:]
	}
	memory := make(map[state]int)
	result := depthFirstSearch(paths, "svr", false, false, memory)
	return strconv.Itoa(result)
}
