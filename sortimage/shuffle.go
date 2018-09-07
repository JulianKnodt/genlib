package main

import (
	"image/color"
	"math/rand"
	"time"
)

func Shuffle(colors []color.Color) {
	var shufRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	curr := len(colors) - 1
	if curr < 0 {
		panic("empty colors")
	}
	for ; curr > 0; curr-- {
		next := shufRand.Intn(curr)
		colors[curr], colors[next] = colors[next], colors[curr]
	}
}
