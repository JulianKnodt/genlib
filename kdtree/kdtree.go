package kdtree

import (
	"errors"
)

type Interface interface {
	// Returns squared distance between items at i, j
	DistSqr(i, j int) float64

	// Compares elements i, j on dimension dim
	// Which will be an int [0,maxDim)
	// so for max dimension 3, it's 0, 1, 2.
	CompareDimension(i, j int, dim int) float64

	// Provides a way to iterate through the interface
	// Does not need to guarantee order and can implement
	// early returns when returns true
	For(func(int) bool)

	// Adds v if possible to Interface and
	// returns index, nil if success
	Insert(v interface{}) (index int, err error)

	// Removes i from Interface
	Delete(i int) error
}

type KDTree struct {
	// Total number of dimensions for this kdtree
	Dimensions int

	// The raw data stored in this kdtree
	// Must not be rearranged once inserted into the kd tree
	Data Interface
	root *Node
}

func NewKDTree(dim int, Data Interface) *KDTree {
	out := &KDTree{dim, Data, nil}
	out.Fix()
	return out
}

func (k *KDTree) Fix() {
  k.root = nil
	k.Data.For(func(i int) bool {
		if k.root == nil {
			k.root = &Node{i, nil, nil}
		} else {
			k.root.Insert(i, 0, k)
		}
		return false
	})
}

// Returns the nearest existing element in the kdtree to i
// Returns -1 for nil trees.
func (k *KDTree) NearestExisting(i int) int {
	if k == nil {
		return -1
	}
	n, _ := k.root.Nearest(i, 0, k)
	return n
}

var ErrorCannotCompare = errors.New("Cannot compare given value in KDTree")

func (k KDTree) Nearest(v interface{}) (nearestIndex int, err error) {
	index, err := k.Data.Insert(v)
	if err != nil {
		return -1, err
	}
	nearestIndex = k.NearestExisting(index)
	err = k.Data.Delete(index)
	return
}

func (k *KDTree) KNearest(v interface{}, num int) ([]int, error) {
	index, err := k.Data.Insert(v)
	if err != nil {
		return nil, err
	}
	out := &distList{}
	k.root.kNearest(index, num, 0, k, out)
	err = k.Data.Delete(index)
	return out.ToIntSlice(), err
}
