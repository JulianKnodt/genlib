package kdtree

import (
	"container/heap"
)

type distOf struct {
	index   int
	distSqr float64
}

type distList []distOf

func (d distList) Len() int            { return len(d) }
func (d distList) Less(i, j int) bool  { return d[i].distSqr < d[j].distSqr }
func (d distList) Swap(i, j int)       { d[i], d[j] = d[j], d[i] }
func (d *distList) Push(x interface{}) { *d = append(*d, x.(distOf)) }
func (d *distList) Pop() interface{} {
	old := *d
	n := len(old)
	result := old[n-1]
	*d = old[:len(old)]
	return result
}

func (d *distList) PushCond(i, maxSize int, distSqr float64) {
	if d.Len() < maxSize {
		heap.Push(d, distOf{i, distSqr})
	} else if distSqr < (*d)[0].distSqr {
		(*d)[0].index = i
		(*d)[0].distSqr = distSqr
		heap.Fix(d, 0)
	}
}

func (d distList) ToIntSlice() []int {
	out := make([]int, d.Len())
	for i, v := range d {
		out[i] = v.index
	}
	return out
}

// Implements k nearest neighbors
func (n *Node) kNearest(to, k, depth int, on *KDTree, list *distList) {
	if n == nil || n.Value == to {
		return
	}

	comp := on.Data.CompareDimension(to, n.Value, depth%on.Dimensions)
	checkLeft := comp < 0.0

	if checkLeft {
		n.Left.kNearest(to, k, depth+1, on, list)
	} else {
		n.Right.kNearest(to, k, depth+1, on, list)
	}

	list.PushCond(n.Value, k, on.Data.DistSqr(to, n.Value))

	if comp*comp < (*list)[0].distSqr {
		if checkLeft {
			n.Right.kNearest(to, k, depth+1, on, list)
		} else {
			n.Left.kNearest(to, k, depth+1, on, list)
		}
	}
}
