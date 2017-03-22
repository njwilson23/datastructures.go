package mergesort

import (
	"fmt"
	"testing"
)

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	var _b interface{}
	for i, _a := range a {
		_b = b[i]
		if _a != _b {
			return false
		}
	}
	return true
}

func TestMerge(t *testing.T) {
	left := []int{2, 5, 8, 13, 18}
	right := []int{1, 4, 8, 11, 12, 16, 21}
	merged := merge(left, right)
	if !slicesEqual(merged, []int{1, 2, 4, 5, 8, 8, 11, 12, 13, 16, 18, 21}) {
		fmt.Println(merged)
		t.Fail()
	}

	left = []int{5}
	right = []int{7}
	merged = merge(left, right)
	if !slicesEqual(merged, []int{5, 7}) {
		fmt.Println(merged)
		t.Fail()
	}
}

func TestMergeSort(t *testing.T) {
	data := []int{43, 27, 8, 3, 75, 6, 32, 61, 3, 12, 6, 3}
	sortedData := MergeSort(data)
	if !slicesEqual(sortedData, []int{3, 3, 3, 6, 6, 8, 12, 27, 32, 43, 61, 75}) {
		t.Fail()
	}
}

func TestMergeSortRecursive(t *testing.T) {
	data := []int{43, 27, 8, 3, 75, 6, 32, 61, 3, 12, 6, 3}
	sortedData := MergeSortRecursive(data)
	if !slicesEqual(sortedData, []int{3, 3, 3, 6, 6, 8, 12, 27, 32, 43, 61, 75}) {
		t.Fail()
	}
}
