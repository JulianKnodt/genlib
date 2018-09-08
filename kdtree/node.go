package kdtree

import (
	"math"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(index, depth int, on *KDTree) {
	comp := on.Data.CompareDimension(index, n.Value, depth%on.Dimensions)
	if comp < 0.0 {
		if n.Left == nil {
			n.Left = &Node{index, nil, nil}
			return
		}
		n.Left.Insert(index, depth+1, on)
	} else {
		if n.Right == nil {
			n.Right = &Node{index, nil, nil}
			return
		}
		n.Right.Insert(index, depth+1, on)
	}
}

// Returns nearest value to on.Data[i] in on.Data other than i
func (n *Node) Nearest(i, depth int, on *KDTree) (int, float64) {
	if n == nil || n.Value == i {
		return -1, math.Inf(1)
	}
	bestIndex := n.Value
	bestDist := on.Data.DistSqr(i, n.Value)
	comp := on.Data.CompareDimension(i, n.Value, depth%on.Dimensions)
	checkLeft := comp < 0.0

	var possIndex int
	var possDist float64
	if checkLeft {
		possIndex, possDist = n.Left.Nearest(i, depth+1, on)
	} else {
		possIndex, possDist = n.Right.Nearest(i, depth+1, on)
	}

	if possDist < bestDist {
		bestDist = possDist
		bestIndex = possIndex
	}

	if comp*comp < bestDist {
		if checkLeft {
			possIndex, possDist = n.Right.Nearest(i, depth+1, on)
		} else {
			possIndex, possDist = n.Left.Nearest(i, depth+1, on)
		}
		if possDist < bestDist {
			bestDist = possDist
			bestIndex = possIndex
		}
	}

	return bestIndex, bestDist
}

func (n *Node) NumChildren() int {
	if n == nil {
		return 0
	}
	return 1 + n.Left.NumChildren() + n.Right.NumChildren()
}

const initSize = 16

func (n *Node) Children() map[int]struct{} {
	out := make(map[int]struct{})
	n.children(out)
	return out
}

func (n *Node) children(res map[int]struct{}) {
	if n == nil {
		return
	}
	res[n.Value] = struct{}{}
	n.Left.children(res)
	n.Right.children(res)
}
