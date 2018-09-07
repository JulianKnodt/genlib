package main

import (
	"image/color"
)

func CompareSimple(a, b color.Color) float64 {
	r1, g1, b1, a1 := a.RGBA()
	r2, g2, b2, a2 := b.RGBA()
	deltaR := float64(r1) - float64(r2)
	deltaG := float64(g1) - float64(g2)
	deltaB := float64(b1) - float64(b2)
	deltaA := float64(a1) - float64(a2)
	return sqr(deltaR) + sqr(deltaG) + sqr(deltaB) + sqr(deltaA)
}

func sqr(a float64) float64 {
	return a * a
}
