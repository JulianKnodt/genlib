package bezier

// Curve represents a bezier curve
// maintaining a list of 3d vectors inside of it
type Curve struct {
	// The control points of this bezier curve
	Points []Vector
}

func LinearInterpolation(start, end Vector, t float64) *Vector {
	return start.Add(*end.Sub(start).SMulSet(t))
}

// Returns one less than the number of points
func (c Curve) Degree() int {
	return len(c.Points) - 1
}

func (c Curve) FrontCurve() *Curve {
	return &Curve{
		c.Points[:len(c.Points)-1],
	}
}

func (c Curve) BackCurve() *Curve {
	return &Curve{
		c.Points[1:len(c.Points)],
	}
}

func (c Curve) Hull() (min, max Vector) {
	return MinMax(c.Points...)
}

func CurveFrom(in ...[3]float64) *Curve {
	points := make([]Vector, len(in))
	for i, v := range in {
		points[i] = Vector(v)
	}
	return &Curve{points}
}

func sqr64(a float64) float64 {
	return a * a
}

// Returns point on curve for t
// in range [0,1]
func (c Curve) At(t float64) *Vector {
	switch len(c.Points) {
	case 1:
		return &c.Points[0]
	case 2:
		return LinearInterpolation(c.Points[0], c.Points[1], t)
	case 3:
		inv := 1 - t
		return c.Points[0].SMul(sqr64(inv)).
			AddSet(*c.Points[1].SMul(2 * inv * t)).
			AddSet(*c.Points[2].SMul(sqr64(t)))
	default:
		return c.FrontCurve().At(t).SMulSet(1 - t).AddSet(*c.BackCurve().At(t).SMul(t))
	}
}

func (c Curve) For(numPartitions int) []Vector {
	amt := 1 / float64(numPartitions)
	out := make([]Vector, numPartitions)
	count := 0
	for i := 0.0; i < 1.0; i += amt {
		out[count] = *c.At(i)
		count++
	}
	return out
}

func DeCasteljaus(points []Vector, t float64) *Vector {
	switch l := len(points); l {
	case 0:
		return nil
	case 1:
		return &points[0]
	default:
		newPoints := make([]Vector, l-1)
		for i := range newPoints {
			newPoints[i] = *points[i].SMul(1 - t).AddSet(*points[i+1].SMul(t))
		}
		return DeCasteljaus(newPoints, t)
	}
}

func (c Curve) Split(t float64) (*Curve, *Curve) {
	size := len(c.Points)
	left := &Curve{make([]Vector, size)}
	right := &Curve{make([]Vector, size)}
	split(c.Points, t, left, right)
	return left, right
}

func split(points []Vector, t float64, l, r *Curve) {
	switch s := len(points); s {
	case 1:
		l.Points = append(l.Points, points[0])
		r.Points = append(r.Points, points[0])
	default:
		l.Points = append(l.Points, points[0])
		r.Points = append(r.Points, points[s-1])
		newPoints := make([]Vector, s-1)
		for i := range newPoints {
			newPoints[i] = *points[i].SMul(1 - t).AddSet(*points[i+1].SMul(t))
		}
		split(newPoints, t, l, r)
	}
}

func (c Curve) RaiseDegree() *Curve {
	k := c.Degree() + 1 // new degree
	newPoints := make([]Vector, k+1)
	points := c.Points

	newPoints[0] = c.Points[0]
	newPoints[k] = c.Points[k-1]
	for i := 1; i < k; i++ {
		newPoints[i] = *points[i].
			SMul(float64(k - i)).
			AddSet(*points[i-1].SMul(float64(i))).
			SDivSet(float64(k))
	}
	return &Curve{newPoints}
}
