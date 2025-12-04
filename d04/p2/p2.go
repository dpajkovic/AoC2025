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
	changed := color.Gray{Y: uint8('X')}

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
	removed := true
	for removed {
		removed = false
		saveImage(floorMap)
		for y := 0; y < maxY; y++ {
			for x := 0; x < maxX; x++ {
				coord := image.Point{X: x, Y: y}
				if floorMap.GrayAt(x, y) == changed {
					floorMap.SetGray(x, y, empty)
					continue
				}
				if floorMap.GrayAt(x, y) == roll {
					countPoint := 0
					for _, delta := range surroundingDeltas {
						surroundingPoint := delta.Add(coord)
						if surroundingPoint.In(rect) {
							if floorMap.GrayAt(surroundingPoint.X, surroundingPoint.Y) == roll {
								countPoint++
							}
						}
					}
					if countPoint < 4 {
						count++
						floorMap.SetGray(coord.X, coord.Y, changed)
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
var gifPalette color.Palette = color.Palette{color.Gray{Y: 0}, color.Gray{Y: 255}, color.RGBA{R: 255, G: 0, B: 0, A: 255}, color.RGBA{R: 255, G: 165, B: 0, A: 255}, color.RGBA{R: 255, G: 255, B: 0, A: 255}, color.RGBA{R: 0, G: 255, B: 0, A: 255}}

func saveImage(img *image.Gray) error {
	// Convert Gray image to Paletted for GIF
	bounds := img.Bounds()
	palettedImg := image.NewPaletted(bounds, gifPalette)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray := img.GrayAt(x, y)
			if gray.Y == uint8('@') {
				palettedImg.SetColorIndex(x, y, 1)
			} else {
				if gray.Y == uint8('X') {
					palettedImg.SetColorIndex(x, y, 2)
					img.SetGray(x, y, color.Gray{Y: uint8('R')})
				} else {
					if gray.Y == uint8('R') {
						palettedImg.SetColorIndex(x, y, 3)
						img.SetGray(x, y, color.Gray{Y: uint8('O')})
					} else {
						if gray.Y == uint8('O') {
							palettedImg.SetColorIndex(x, y, 4)
							img.SetGray(x, y, color.Gray{Y: uint8('Y')})

						} else {
							if gray.Y == uint8('Y') {
								palettedImg.SetColorIndex(x, y, 5)
								img.SetGray(x, y, color.Gray{Y: uint8('G')})

							} else {
								if gray.Y == uint8('G') {
									palettedImg.SetColorIndex(x, y, 0)
									img.SetGray(x, y, color.Gray{Y: uint8('.')})
								} else {
									palettedImg.SetColorIndex(x, y, 0)
								}
							}
						}
					}
				}
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
