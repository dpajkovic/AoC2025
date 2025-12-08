//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p1

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	X, Y, Z int
}

func distance(a, b point) float64 {

	d := math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2) + math.Pow(float64(a.Z-b.Z), 2))
	return d
}

func P1(input []string, iter int) string {
	var junctions []point
	for _, line := range input {
		nums := strings.Split(line, ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		z, _ := strconv.Atoi(nums[2])
		junctions = append(junctions, point{X: x, Y: y, Z: z})
	}

	// calculate distances between all junctions

	var distances map[int]map[int]float64 = make(map[int]map[int]float64)
	for i := 0; i < len(junctions); i++ {
		distances[i] = make(map[int]float64)
		for j := 0; j < len(junctions); j++ {
			if i != j {
				distances[i][j] = distance(junctions[i], junctions[j])
			}
		}
	}

	var inCircuit map[point]int = make(map[point]int)
	newestCircuitIndex := 1

	// go through all distances and build circuits on each iteration between closest junctions. if they are already in a circuit, renuber all the junctions in the circuit to the lowest circuit index
	for it := 0; it < iter; it++ {
		minDistance := math.MaxFloat64
		var minI, minJ int
		for i := 0; i < len(junctions)-1; i++ {
			for j := i + 1; j < len(junctions); j++ {
				if val, ok := distances[i][j]; ok && val < minDistance {
					minDistance = distances[i][j]
					minI = i
					minJ = j
				}
			}
		}
		circuitI, okI := inCircuit[junctions[minI]]
		circuitJ, okJ := inCircuit[junctions[minJ]]
		if okI && okJ {
			// both junctions are already in a circuit, renumber all junctions in circuitJ to circuitI
			if circuitI != circuitJ {
				for k, v := range inCircuit {
					if v == circuitJ {
						inCircuit[k] = circuitI
					}
				}
			}
		} else if okI {
			// only junction i is in a circuit, add junction j to it
			inCircuit[junctions[minJ]] = circuitI
		} else if okJ {
			// only junction j is in a circuit, add junction i to it
			inCircuit[junctions[minI]] = circuitJ
		} else {
			// neither junction is in a circuit, create a new circuit
			inCircuit[junctions[minI]] = newestCircuitIndex
			inCircuit[junctions[minJ]] = newestCircuitIndex
			newestCircuitIndex++
		}
		// remove this distance from the map
		delete(distances[minI], minJ)
		delete(distances[minJ], minI)
	}

	// find sizes of 3 largest circuits and return product of their sizes
	circuitSizes := make(map[int]int)
	for _, v := range inCircuit {
		circuitSizes[v]++
	}
	var largestSizes []int
	for _, size := range circuitSizes {
		largestSizes = append(largestSizes, size)
	}
	// sort largestSizes descending
	sort.IntSlice(largestSizes).Sort()
	prod := 1
	for i := 0; i < 3 && i < len(largestSizes); i++ {
		prod *= largestSizes[len(largestSizes)-1-i]
	}

	return strconv.Itoa(prod)
}
