package bezier

import (
	"testing"
)

func TestLinear(t *testing.T) {
	start := Vector{0, 0, 0}
	end := Vector{1, 1, 1}

	if *LinearInterpolation(start, end, 1) != end {
		t.Fail()
	}

	if *LinearInterpolation(start, end, 0) != start {
		t.Fail()
	}
}

func BenchmarkLinear(b *testing.B) {
	start := Vector{0, 0, 0}
	end := Vector{1, 1, 1}

	for i := 0; i < b.N; i++ {
		LinearInterpolation(start, end, 0.5)
	}
}

func TestQuartic(t *testing.T) {
	p0 := [3]float64{10, 3, 4}
	p1 := [3]float64{5, 5, 5}
	p2 := [3]float64{6, 3, 2}
	p3 := [3]float64{0, 0, 0}

	curve := CurveFrom(p0, p1, p2, p3)
	curve.At(0.5)
}

func BenchmarkQuartic(b *testing.B) {
	p0 := [3]float64{10, 3, 4}
	p1 := [3]float64{5, 5, 5}
	p2 := [3]float64{6, 3, 2}
	p3 := [3]float64{0, 0, 0}

	curve := CurveFrom(p0, p1, p2, p3)
	for i := 0; i < b.N; i++ {
		curve.At(0.5)
	}
}

func TestRaiseDegree(t *testing.T) {
	p0 := [3]float64{0, 0, 0}
	p1 := [3]float64{1, 1, 1}

	curve := CurveFrom(p0, p1)
	raised := curve.RaiseDegree()
	for i := 0.0; i < 1.0; i += 0.05 {
		if !(*curve.At(i)).ApproxEqual(*raised.At(i), 0.0000001) {
			t.Error("Raised degree not equal to original")
		}
	}
}

func BenchmarkRaiseDegree(b *testing.B) {
	p0 := [3]float64{10, 3, 4}
	p1 := [3]float64{5, 5, 5}
	p2 := [3]float64{6, 3, 2}
	p3 := [3]float64{0, 0, 0}

	curve := CurveFrom(p0, p1, p2, p3)
	for i := 0; i < b.N; i++ {
		curve.RaiseDegree()
	}
}

func TestSplit(t *testing.T) {
	p0 := [3]float64{10, 3, 4}
	p1 := [3]float64{5, 5, 5}
	p2 := [3]float64{6, 3, 2}
	p3 := [3]float64{0, 0, 0}

	curve := CurveFrom(p0, p1, p2, p3)
	_, _ = curve.Split(0.4)
}
