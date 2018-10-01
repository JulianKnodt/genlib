package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"

	"github.com/julianknodt/algos/kdtree"
)

func Sub(a, b []float64) []float64 {
	out := make([]float64, len(a))
	for i, v := range a {
		out[i] = v - b[i]
	}
	return out
}

func sqr(a float64) float64 {
	return a * a
}

func PiecewiseDist(a, b []float64) float64 {
	var sum float64
	for i, v := range a {
		sum += sqr(v - b[i])
	}
	return sum
}

func RandOf(l int) []float64 {
	out := make([]float64, l)
	for i := range out {
		out[i] = rand.Float64()
	}
	return out
}

func ConstOf(l int, val float64) []float64 {
	out := make([]float64, l)
	for i := range out {
		out[i] = val
	}
	return out
}

func Add(a, b []float64) []float64 {
	out := make([]float64, len(a))
	for i, v := range a {
		out[i] = v + b[i]
	}
	return out
}

func SMulSet(a []float64, k float64) []float64 {
	for i, v := range a {
		a[i] = v * k
	}
	return a
}

type pl [][]float64

func (p pl) Len() int                               { return len(p) }
func (p pl) DistSqr(i, j int) float64               { return PiecewiseDist(p[i], p[j]) }
func (p pl) CompareDimension(i, j, dim int) float64 { return p[i][dim] - p[j][dim] }
func (p *pl) Insert(v interface{}) (int, error) {
	if val, ok := v.([]float64); ok {
		index := p.Len()
		*p = append(*p, val)
		return index, nil
	}
	panic("Should only accept floats")
}
func (p *pl) Delete(i int) error {
	*p = append((*p)[:i], (*p)[i+1:]...)
	return nil
}

var (
	dims    = flag.Uint("d", 2, "Number of dimensions to use for points")
	points  = flag.Uint("p", 500, "Number of points to insert")
	iters   = flag.Uint("i", 10000, "Number of iterations to run")
	out     = flag.String("o", "out.png", "Where to output the image")
	numComp = flag.Uint("n", 5, "Number to compare for nearest neighbor")
	slFlag  = flag.Uint("sl", 1600, "Side length of out")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	numIters := *iters
	numPoints := *points
	numDims := int(*dims)
	comp := int(*numComp)
	sideLength := float64(*slFlag)
	img := image.NewRGBA(image.Rect(0, 0, int(sideLength), int(sideLength)))
	points := &pl{}
	tree := kdtree.NewKDTree(numDims, points)
	for i := 0; i < int(numIters); i++ {
		if numPoints > 0 {
			points.Insert(RandOf(numDims))
			numPoints--
		}
		newPoints := &pl{}
		for i := range *points {
			initial := (*points)[i]
			diff := initial

			for _, v := range tree.KNearestExisting(i, comp) {
				near := (*points)[v]
				diff = Add(diff, SMulSet(Sub(near, initial),
					1.0/PiecewiseDist(initial, near)))
			}
			*newPoints = append(*newPoints, diff)
		}
		*points = *newPoints
		tree.Fix()
		for j, v := range *points {
			img.Set(int(v[0]+sideLength/2), int(v[1]+sideLength/2),
				color.RGBA{uint8(j * 21 % 255), uint8(j * 37 % 255), uint8(j * 17 % 255),
					uint8(255 * float64(i) / float64(numIters))})
		}
	}
	f, err := os.Create(*out)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err = png.Encode(f, img); err != nil {
		panic(err)
	}
}
