//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p1

import (
	"image"
	"strconv"
	"strings"
)

func P1(input []string) string {
	tiles := []image.Point{}
	for _, line := range input {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		tiles = append(tiles, image.Point{X: x, Y: y})

	}
	maxSize := 0
	for i := 0; i < len(tiles)-1; i++ {
		for j := i + 1; j < len(tiles); j++ {
			r := image.Rect(min(tiles[i].X, tiles[j].X), min(tiles[i].Y, tiles[j].Y),
				max(tiles[i].X, tiles[j].X), max(tiles[i].Y, tiles[j].Y))
			dX := abs(r.Dx())
			dY := abs(r.Dy())
			size := (dX + 1) * (dY + 1)

			// println("tiles", tiles[i].String(), " and ", tiles[j].String(), "dX:", dX, "dY:", dY, "size:", size)

			if size > maxSize {
				maxSize = size
			}
		}
	}

	return strconv.Itoa(maxSize)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
