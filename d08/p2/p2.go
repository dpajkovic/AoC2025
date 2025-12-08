//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p2

import (
	"math"
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

func P2(input []string) string {

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

	// keep adding closest junctions to circuits until all junctions are in one circuit.
	// note the junctions that last connected two circuits

	var inCircuit map[point]int = make(map[point]int)
	newestCircuitIndex := 1
	var lastAdded []int
	numCircuits := 0
	for {
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
				numCircuits--
			}
		} else if okI {
			inCircuit[junctions[minJ]] = circuitI
			// record last time we added a single junction to an existing circuit
			lastAdded = []int{minI, minJ}
		} else if okJ {
			inCircuit[junctions[minI]] = circuitJ
			lastAdded = []int{minJ, minI}
		} else {
			inCircuit[junctions[minI]] = newestCircuitIndex
			inCircuit[junctions[minJ]] = newestCircuitIndex
			newestCircuitIndex++
			numCircuits++
		}
		// remove this distance from the map
		delete(distances[minI], minJ)
		delete(distances[minJ], minI)

		// stop when all junctions have been assigned and there's a single circuit
		if len(inCircuit) == len(junctions) && numCircuits == 1 {
			break
		}
	}
	// multiply X coordinates of the last added junctions (no fallback)
	chosen := lastAdded
	product := junctions[chosen[0]].X * junctions[chosen[1]].X
	return strconv.Itoa(product)

}
