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
	dimX, dimY := len(input[0]), len(input)
	grid := image.NewGray(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: dimX, Y: dimY},
	})
	for y, l := range input {
		for x, c := range l {
			grid.SetGray(x, y, color.Gray{Y: uint8(c)})
		}
	}
	below := image.Point{X: 0, Y: 1}
	split := []image.Point{
		{X: -1, Y: 1},
		{X: 1, Y: 1},
	}
	for y := 0; y < dimY-1; y++ {
		for x := 0; x < dimX; x++ {
			current := grid.GrayAt(x, y).Y
			pointBelow := image.Point{X: x + below.X, Y: y + below.Y}
			if current == uint8('S') {
				grid.SetGray(pointBelow.X, pointBelow.Y, gray('|'))
				continue
			} else if current == uint8('|') {
				if grid.GrayAt(pointBelow.X, pointBelow.Y).Y == uint8('^') {
					for _, sp := range split {
						splitPoint := image.Point{X: x + sp.X, Y: y + sp.Y}
						grid.SetGray(splitPoint.X, splitPoint.Y, gray('|'))
					}
					continue

				} else {
					grid.SetGray(pointBelow.X, pointBelow.Y, gray('|'))
				}
				continue
			}
		}
	}
	counter := 0
	pointAbove := image.Point{X: 0, Y: -1}
	for y := 0; y < dimY; y++ {
		for x := 0; x < dimX; x++ {
			current := grid.GrayAt(x, y).Y
			if current == uint8('^') {
				above := image.Point{X: x + pointAbove.X, Y: y + pointAbove.Y}
				if grid.GrayAt(above.X, above.Y).Y == uint8('|') {
					counter++
				}
			}
		}
	}

	return strconv.Itoa(counter)
}

func gray(c rune) color.Gray {
	return color.Gray{Y: uint8(c)}
}
