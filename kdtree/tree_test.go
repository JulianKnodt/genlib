package kdtree

import (
	"errors"
	"math"
	"math/rand"
	"testing"
	"time"
)

type Point struct {
	X, Y, Z float64
}

func RandPoint() Point {
	return Point{
		rand.Float64(),
		rand.Float64(),
		rand.Float64(),
	}
}

func (p Point) SqrDist(o Point) float64 {
	return sqr(p.X-o.X) +
		sqr(p.Y-o.Y) +
		sqr(p.Z-o.Z)
}

type PL map[int]Point

func (p PL) NaiveMinDistance(to Point) int {
	min := math.Inf(1)
	index := -1
	for i, v := range p {
		if dist := v.SqrDist(to); dist < min {
			min = dist
			index = i
		}
	}
	return index
}

func sqr(a float64) float64           { return a * a }
func (p PL) Len() int                 { return len(p) }
func (p PL) DistSqr(i, j int) float64 { return p[i].SqrDist(p[j]) }
func (p PL) Delete(i int) error {
	delete(p, i)
	return nil
}
func (p PL) Insert(v interface{}) (int, error) {
	if point, ok := v.(Point); ok {
		curr := len(p)
		for {
			if _, has := p[curr]; !has {
				p[curr] = point
				return curr, nil
			}
			curr++
		}
	}
	return -1, errors.New("Not a point")
}

func (p PL) For(cb func(int) bool) {
	for i := range p {
		if cb(i) {
			return
		}
	}
}

func (p PL) CompareDimension(i, j, dim int) float64 {
	switch dim {
	case 0:
		return (p[i].X - p[j].X)
	case 1:
		return (p[i].Y - p[j].Y)
	case 2:
		return (p[i].Z - p[j].Z)
	default:
		panic("Unexpected dimension")
	}
}

func makePL() PL {
	rand.Seed(time.Now().UnixNano())
	out := make(map[int]Point, 500)
	for i := 0; i < 500; i++ {
		out[i] = RandPoint()
	}
	return PL(out)
}

func TestNearest(t *testing.T) {
	pl := makePL()

	kd := NewKDTree(3, pl)
	testPoint := RandPoint()
	index, err := kd.Nearest(testPoint)
	if err != nil {
		t.Error(err)
	}
	naiveIndex := pl.NaiveMinDistance(testPoint)
	if naiveIndex != index {
		t.Error("Nearest KDTree did not return same as naive check", naiveIndex, index)
	}
}

func TestKNearest(t *testing.T) {
	pl := makePL()

	kd := NewKDTree(3, pl)
	testPoint := RandPoint()
	index, err := kd.KNearest(testPoint, 1)
	if err != nil {
		t.Error(err)
	}
	naiveIndex := pl.NaiveMinDistance(testPoint)
	if index[0] != naiveIndex {
		t.Error("KNearest does not return correct result")
	}
}
