package main

import (
	"image"
	"image/color"
//	"image/color/palette"
	"log"
)


func Palettize(img image.Image, p color.Palette) *image.Paletted {
	log.Println("Started converting palette")
	bounds := img.Bounds()
	result := image.NewPaletted(bounds, p)
	x, y := bounds.Dx(), bounds.Dy()
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			result.Set(i, j, img.At(i, j))
		}
	}
	log.Println("Finished converting palette")
	return result
}

const maxSize = 256
func GetPalette(img image.Image) color.Palette {
	uniqColors := make(map[color.Color]struct{}, 100)
	bounds := img.Bounds()
	x, y := bounds.Dx(), bounds.Dy()
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			uniqColors[img.At(i, j)] = struct{}{}
		}
	}

	result := make([]color.Color, 0, len(uniqColors))
	for c := range uniqColors {
		result = append(result, c)
	}
	return color.Palette(result[:maxSize])
}
