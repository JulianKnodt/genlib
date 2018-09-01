# Bezier 

A go package implementing bezier curves.

Credit to: https://pomax.github.io/bezierinfo/#introduction

```
package bezier // import "github.com/julianknodt/bezier"

A package implementing some basic bezier curve functions

Based on: https://pomax.github.io/bezierinfo/

func CurveFrom(in ...[3]float64) *Curve
type Vector [3]float64
func DeCasteljaus(points []Vector, t float64) *Vector
  func LinearInterpolation(start, end Vector, t float64) *Vector
func MinMax(vs ...Vector) (min, max Vector)

type Curve struct {
  Points []Vector
}

func (c Curve) At(t float64) *Vector
func (c Curve) BackCurve() *Curve
func (c Curve) Degree() int
func (c Curve) For(numPartitions int) []Vector
func (c Curve) FrontCurve() *Curve
func (c Curve) Hull() (min, max Vector)
func (c Curve) RaiseDegree() *Curve
func (c Curve) Split(t float64) (*Curve, *Curve)
```
