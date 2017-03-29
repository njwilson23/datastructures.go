package heap

import (
	"errors"
)

var OVERFLOW_ERROR = errors.New("heap is at maximum size")

var EMPTY_HEAP_ERROR = errors.New("empty heap")

type Heap struct {
	value    []float64
	label    []int
	size     int
	capacity int
}

func New(capacity int) *Heap {
	return &Heap{make([]float64, capacity), make([]int, capacity), 0, capacity}
}

// MaxHeapify enforces the max-heap property of a Heap whose parent node is i.
func (h *Heap) MaxHeapify(i int) {
	var ilargest, ileft, iright int

	for {
		ileft = 2*(i+1) - 1
		iright = 2 * (i + 1)

		if h.size > ileft && h.value[ileft] > h.value[i] {
			ilargest = ileft
		} else {
			ilargest = i
		}

		if h.size > iright && h.value[iright] > h.value[ilargest] {
			ilargest = iright
		}

		if i != ilargest {
			h.value[i], h.value[ilargest] = h.value[ilargest], h.value[i]
			h.label[i], h.label[ilargest] = h.label[ilargest], h.label[i]
			i = ilargest
		} else {
			break
		}
	}
}

func (h *Heap) Maximum() (int, float64, error) {
	if h.size == 0 {
		return 0, 0.0, EMPTY_HEAP_ERROR
	}
	return h.label[0], h.value[0], nil
}

func (h *Heap) ExtractMaximum() (int, float64, error) {
	if h.size == 0 {
		return 0, 0.0, EMPTY_HEAP_ERROR
	}
	labelMax, valueMax, _ := h.Maximum()
	h.size--
	h.value = append(h.value[1:], 0.0)
	h.label = append(h.label[1:], 0)
	h.MaxHeapify(0)
	return labelMax, valueMax, nil
}

func BuildMaxHeap(values []float64, labels []int) *Heap {
	h := New(len(values))
	h.size = len(values)
	h.value = values
	h.label = labels
	for i := h.size / 2; i != -1; i-- {
		h.MaxHeapify(i)
	}
	return h
}
