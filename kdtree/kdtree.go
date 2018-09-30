package kdtree

import (
	"errors"
)

type Interface interface {
	// Returns the length of the interface
	Len() int

	// Returns squared distance between items at i, j
	DistSqr(i, j int) float64

	// Compares elements i, j on dimension dim
	// Which will be an int [0,maxDim)
	CompareDimension(i, j int, dim int) float64

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
	k.root = &Node{0, nil, nil}
	for i := 1; i < k.Data.Len(); i++ {
		k.root.Insert(i, 0, k)
	}
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
