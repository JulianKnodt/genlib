package bezier

import (
	"math"
)

type Vector [3]float64

// v + o returns a new vector
func (v Vector) Add(o Vector) *Vector {
	return &Vector{v[0] + o[0], v[1] + o[1], v[2] + o[2]}
}

// v += o sets v to the new vector
func (v *Vector) AddSet(o Vector) *Vector {
	v[0] += o[0]
	v[1] += o[1]
	v[2] += o[2]
	return v
}

func MinMax(vs ...Vector) (min, max Vector) {
	min = Vector{math.Inf(1), math.Inf(1), math.Inf(1)}
	max = Vector{math.Inf(-1), math.Inf(-1), math.Inf(-1)}
	for _, v := range vs {
		max[0] = math.Max(max[0], v[0])
		max[1] = math.Max(max[1], v[1])
		max[0] = math.Max(max[0], v[0])

		min[0] = math.Min(min[0], v[0])
		min[1] = math.Min(min[1], v[1])
		min[2] = math.Min(min[2], v[2])
	}
	return
}

// v - o returns a new vector
func (v Vector) Sub(o Vector) *Vector {
	return &Vector{v[0] - o[0], v[1] - o[1], v[2] - o[2]}
}

// Multiplies v and returns a new vector
func (v Vector) SMul(k float64) *Vector {
	return &Vector{k * v[0], k * v[1], k * v[2]}
}

// Multiplies and Sets v
func (v *Vector) SMulSet(k float64) *Vector {
	v[0] *= k
	v[1] *= k
	v[2] *= k
	return v
}

// Divides and sets v
func (v *Vector) SDivSet(k float64) *Vector {
	v[0] /= k
	v[1] /= k
	v[2] /= k
	return v
}

// Returns magnitude of a vector
func (v Vector) Magn() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

func (v Vector) ApproxEqual(o Vector, epsilon float64) bool {
	return math.Abs(v[0]-o[0]) < epsilon &&
		math.Abs(v[1]-o[1]) < epsilon &&
		math.Abs(v[2]-o[2]) < epsilon
}
