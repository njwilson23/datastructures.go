/*
 * Merge-sort is a divide-and-conquor algorithm for sorting arrays of objects.
 * Although not really a data structure in the sense of the other algorithms
 * in this repository, I've included merge sort because it provides a nice
 * demonstration of manipulating array data structures in both a top-down and a
 * bottom-up manner.
 *
 * In the top down approach, the array is sub-divided repeatedly until each
 * subarray is only one item long, and these are then merged back up the call
 * stack in a way that the pieces at every level are sorted.
 *
 * In the bottom up approach, we start with atomic array elements, and
 * agglomerate them into increasingly large sorted segments.
 *
 * Both versions of this algorithm have O(n log n) time performance, however
 * the bottom-up version below is implemented without needing a stack of calling
 * frames can so may be preferable in a language such as Go.
 */

package mergesort

// RecursiveMergeSort implements a "top-down" recursive merge sort algorithm
func RecursiveMergeSort(sortable []int) []int {
	n := len(sortable)
	if n == 1 {
		return sortable
	}
	left := RecursiveMergeSort(sortable[:n/2])
	right := RecursiveMergeSort(sortable[n/2:])
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
