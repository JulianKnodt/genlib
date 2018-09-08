package kdtree

import (
	"container/heap"
	"testing"
)

// This file exists as a sanity check for the ordering of the heap in knearest neighbors

func TestDistList(t *testing.T) {
	dl := &distList{
		distOf{1, 1.0},
		distOf{1, 5.0},
		distOf{1, 7.0},
	}
	heap.Init(dl)
	if (*dl)[0].distSqr > (*dl)[2].distSqr {
		t.Error("Incorrect order of dist list")
	}
}
