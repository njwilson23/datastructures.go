package mergesort

// MergeSortRecursive implements a "top-down" recursive merge sort algorithm
func MergeSortRecursive(sortable []int) []int {
	n := len(sortable)
	if n == 1 {
		return sortable
	}
	left := MergeSortRecursive(sortable[:n/2])
	right := MergeSortRecursive(sortable[n/2:])
	return merge(left, right)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MergeSort implements an "bottom-up" non-recursive merge sort algorithm
func MergeSort(sortable []int) []int {
	var left, right []int
	mergeSize := 1
	n := len(sortable)
	for mergeSize <= n {
		i := 0
		for i < n {
			left = sortable[i : i+mergeSize]
			right = sortable[i+mergeSize : min(n, i+2*mergeSize)]
			sortable = append(sortable[:i], append(merge(left, right), sortable[min(n, i+2*mergeSize):]...)...)
			i = i + 2*mergeSize
		}
		mergeSize = mergeSize * 2
	}
	return sortable
}

// merge combines two sorted slices into a single sorted slice
func merge(left, right []int) []int {
	posLeft := 0
	posRight := 0
	n := len(left) + len(right)
	merged := []int{}

	for len(merged) != n {
		if posLeft == len(left) {
			merged = append(merged, right[posRight])
			posRight++
		} else if posRight == len(right) {
			merged = append(merged, left[posLeft])
			posLeft++
		} else if left[posLeft] < right[posRight] {
			merged = append(merged, left[posLeft])
			posLeft++
		} else {
			merged = append(merged, right[posRight])
			posRight++
		}
	}
	return merged
}
