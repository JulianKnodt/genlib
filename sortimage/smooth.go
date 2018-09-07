package main

import (
	"image"
	"image/color"
)

func Average(colors ...color.Color) color.Color {
	var r, g, b, a float64
	length := float64(len(colors))
	for _, c := range colors {
		rgba := color.RGBAModel.Convert(c).(color.RGBA)
		r += float64(rgba.R) / length
		g += float64(rgba.G) / length
		b += float64(rgba.B) / length
		a += float64(rgba.A) / length
	}

	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func Smooth(img image.Image) image.Image {
	bounds := img.Bounds()
	out := image.NewRGBA(bounds)
	w, h := bounds.Dx(), bounds.Dy()
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			out.Set(i, j,
				Average(img.At(i, j), img.At(i+1, j), img.At(i, j+1), img.At(i-1, j), img.At(i, j-1)))
		}
	}
	return out
}
