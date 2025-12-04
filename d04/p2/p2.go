//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p2

import (
	"image"
	"image/color"
	"image/gif"
	"os"
	"strconv"
)

func P2(input []string) string {
	maxX, maxY := len(input[0]), len(input)
	rect := image.Rect(0, 0, maxX, maxY)
	floorMap := image.NewGray(rect)

	for y, line := range input {
		for x, char := range line {
			floorMap.SetGray(x, y, color.Gray{Y: uint8(char)})
		}
	}

	roll := color.Gray{Y: uint8('@')}
	empty := color.Gray{Y: uint8('.')}

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
	imageNr := 0
	removed := true
	for removed {
		removed = false
		saveImage(floorMap, imageNr)
		imageNr++
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
						floorMap.SetGray(x, y, empty)
						removed = true
					}
				}
			}
		}
	}
	saveGifToFile("output.gif")
	return strconv.Itoa(count)
}

var gifImages []*image.Paletted
var gifDelays []int
var gifPalette color.Palette = color.Palette{color.Gray{Y: 0}, color.Gray{Y: 255}}

func saveImage(img *image.Gray, order int) error {
	// Convert Gray image to Paletted for GIF
	bounds := img.Bounds()
	palettedImg := image.NewPaletted(bounds, gifPalette)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray := img.GrayAt(x, y)
			if gray.Y == uint8('@') {
				palettedImg.SetColorIndex(x, y, 1)
			} else {
				palettedImg.SetColorIndex(x, y, 0)
			}
		}
	}
	gifImages = append(gifImages, palettedImg)
	gifDelays = append(gifDelays, 10) // 10 = 100ms per frame

	return nil
}

// Save the accumulated GIF frames to a file.

func saveGifToFile(filename string) error {
	if len(gifImages) == 0 {
		return nil
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	anim := gif.GIF{
		Image: gifImages,
		Delay: gifDelays,
	}
	return gif.EncodeAll(f, &anim)
}
