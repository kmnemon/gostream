package gostream

import "testing"

func TestSplitSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	copySlices := splitSlice[int](a, 4)

	if len(copySlices) != 4 {
		t.Error("split slice count wrong")
	}
	if !equalSliceHelper[int](copySlices[0], []int{1, 2}) {
		t.Error("split slice wrong")
	}

	if !equalSliceHelper[int](copySlices[3], []int{7, 8, 9}) {
		t.Error("splice slice last group wrong")
	}

}
