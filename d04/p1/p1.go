//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p1

import (
	"image"
	"image/color"
	"strconv"
)

func P1(input []string) string {
	maxX, maxY := len(input[0]), len(input)
	rect := image.Rect(0, 0, maxX, maxY)
	floorMap := image.NewGray(rect)

	for y, line := range input {
		for x, char := range line {
			floorMap.SetGray(x, y, color.Gray{Y: uint8(char)})
		}
	}

	roll := color.Gray{Y: uint8('@')}
	// empty := color.Gray{Y: uint8('.')}

	surroundingDeltas := []image.Point{
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: 0, Y: 1},
		{X: -1, Y: -1},
		{X: -1, Y: 1},
		{X: 1, Y: -1},
		{X: 1, Y: 1},
	}

	count := 0

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if floorMap.GrayAt(x, y) == roll {
				countPoint := 0
				for _, delta := range surroundingDeltas {
					surroundingPoint := image.Point{X: x + delta.X, Y: y + delta.Y}
					if surroundingPoint.In(rect) {
						if floorMap.GrayAt(surroundingPoint.X, surroundingPoint.Y) == roll {
							countPoint++
						}
					}
				}
				if countPoint < 4 {
					count++
				}
			}
		}
	}

	return strconv.Itoa(count)
}
