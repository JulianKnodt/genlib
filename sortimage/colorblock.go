package main

import (
	"image"
	"image/color"
)

type ColorBlock struct {
	Width, Height int
	Data          []color.Color
}

func NewColorBlock(from image.Image) *ColorBlock {
	bounds := from.Bounds()
	x := bounds.Dx()
	y := bounds.Dy()
	data := make([]color.Color, y*x)
	out := &ColorBlock{
		Width:  bounds.Dx(),
		Height: bounds.Dy(),
		Data:   data,
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			out.Set(i, j, from.At(i, j))
		}
	}
	Shuffle(out.Data)
	return out
}

func (c ColorBlock) At(x, y int) color.Color {
	return c.Data[x+c.Width*y]
}

func (c *ColorBlock) Set(x, y int, col color.Color) {
	c.Data[x+c.Width*y] = col
}

func (c ColorBlock) ToImage() image.Image {
	x := c.Width
	y := c.Height
	out := image.NewRGBA64(image.Rect(0, 0, x, y))
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			out.Set(i, j, c.At(i, j))
		}
	}
	return out
}

func (c *ColorBlock) TrySwap(x1, y1, x2, y2 int, basis image.Image) (didSwap bool) {
	imgColor1 := c.At(x1, y1)
	imgColor2 := c.At(x2, y2)
	init1 := CompareSimple(basis.At(x1, y1), imgColor1)
	init2 := CompareSimple(basis.At(x2, y2), imgColor2)

	poss1 := CompareSimple(basis.At(x1, y1), imgColor2)
	poss2 := CompareSimple(basis.At(x2, y2), imgColor1)

	delta1 := poss1 - init1
	delta2 := poss2 - init2

	if delta1+delta2 <= 0 {
		c.Set(x1, y1, imgColor2)
		c.Set(x2, y2, imgColor1)
		return true
	}
	return false
}
