package main

import (
	"image/color"
)

var model = color.RGBA64Model

func CompareSimple(a, b color.Color) float64 {
	aRGBA := model.Convert(a).(color.RGBA64)
	bRGBA := model.Convert(b).(color.RGBA64)
	deltaR := float64(aRGBA.R) - float64(bRGBA.R)
	deltaG := float64(aRGBA.G) - float64(bRGBA.G)
	deltaB := float64(aRGBA.B) - float64(bRGBA.B)
	return sqr(deltaR) + sqr(deltaG) + sqr(deltaB)
}

func sqr(a float64) float64 {
	return a * a
}
