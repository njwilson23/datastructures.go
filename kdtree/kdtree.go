package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

var (
	POINT_NOT_FOUND = errors.New("point not found")
	EMPTYNODE       = errors.New("node is empty")
	DIMENSION_ERROR = errors.New("dimension mismatch")
)

type KDTree struct {
	k    int
	root *node
}

type node struct {
	dim         int
	split       float64
	left, right *node
	points      []point
	cap         int
}

type point []float64

func NewKDTree(k int) *KDTree {
	kdtree := new(KDTree)
	kdtree.k = k
	kdtree.root = &node{dim: 0, split: 0.0, left: nil, right: nil, points: []point{}, cap: 64}
	return kdtree
}

func (kdtree *KDTree) Insert(pt point) (err error) {
	return kdtree.root.insert(pt)
}

func (kdtree *KDTree) Delete(pt point) (err error) {
	return kdtree.root.delete(pt)
}

func (kdtree *KDTree) Search(bbox []float64) (found []point, err error) {
	if len(bbox) != kdtree.k*2 {
		err = DIMENSION_ERROR
		return
	}
	ch := make(chan point)
	go kdtree.root.search(bbox, ch)
	for pt := range ch {
		found = append(found, pt)
	}
	return
}

func (n *node) search(bounds []float64, c chan<- point) {
	if n.left != nil {
		cleft := make(chan point)
		cright := make(chan point)

		if n.split > bounds[n.dim] {
			go n.left.search(bounds, cleft)
		} else {
			close(cleft)
		}

		if n.split < bounds[n.dim+len(bounds)/2] {
			go n.right.search(bounds, cright)
		} else {
			close(cright)
		}

		for cleft != nil || cright != nil {
			select {
			case pt, ok := <-cleft:
				if ok {
					c <- pt
				} else {
					cleft = nil
				}
			case pt, ok := <-cright:
				if ok {
					c <- pt
				} else {
					cright = nil
				}
			}
		}
	} else {
		var within bool
		for _, pt := range n.points {
			within = true
			for i, v := range pt {
				if (v < bounds[i]) || (v > bounds[i+len(pt)]) {
					within = false
					break
				}
			}
			if within {
				c <- pt
			}
		}
	}
	close(c)
}

func (n *node) insert(pt point) (err error) {
	if n.left != nil {
		if pt[n.dim] < n.split {
			n.left.insert(pt)
		} else {
			n.right.insert(pt)
		}
	} else if len(n.points) == n.cap {
		err = n.makesplit(n.dim)
	} else {
		n.points = append(n.points, pt)
	}
	return
}

func (n *node) delete(pt point) (err error) {
	// find point
	if n.left != nil {
		if pt[n.dim] < n.split {
			n.left.delete(pt)
		} else {
			n.right.delete(pt)
		}
	} else {
		found := false
		for i, c := range n.points {
			if c.equals(pt) {
				n.points = append(n.points[:i], n.points[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			err = POINT_NOT_FOUND
		}
	}
	return
}

func (n *node) makesplit(dim int) error {
	if len(n.points) == 0 {
		return EMPTYNODE
	}
	vals := make([]float64, len(n.points))
	for i, pt := range n.points {
		vals[i] = pt[dim]
	}

	n.split = ninther(vals)

	n.left = &node{dim: (dim + 1) % len(n.points[0]),
		split:  0.0,
		left:   nil,
		right:  nil,
		points: []point{},
		cap:    n.cap}

	n.right = &node{dim: (dim + 1) % len(n.points[0]),
		split:  0.0,
		left:   nil,
		right:  nil,
		points: []point{},
		cap:    n.cap}

	for _, pt := range n.points {
		if pt[dim] < n.split {
			n.left.insert(pt)
		} else {
			n.right.insert(pt)
		}
	}
	n.points = []point{}
	return nil
}

func (pt point) equals(pt2 point) bool {
	for i, v := range pt {
		if v != pt2[i] {
			return false
		}
	}
	return len(pt) == len(pt2)
}

func median3(a, b, c float64) float64 {
	return math.Max(math.Min(a, b), math.Min(math.Max(a, b), c))
}

func ninther(pts []float64) float64 {
	n := len(pts)
	th := int(math.Floor(float64(n) / 9))
	return median3(median3(pts[0], pts[1*th], pts[2*th]),
		median3(pts[3*th], pts[4*th], pts[5*th]),
		median3(pts[6*th], pts[7*th], pts[8*th]))
}

func main() {

	kdtree := NewKDTree(2)
	for i := 0; i != 1000000; i++ {
		kdtree.Insert([]float64{rand.Float64(), rand.Float64()})
	}

	points2, err := kdtree.Search([]float64{0.4, 0.7, 0.5, 0.8})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("found %d points with k=2\n", len(points2))
	fmt.Printf("expected about 10000\n")

	kdtree = NewKDTree(5)
	for i := 0; i != 1000000; i++ {
		kdtree.Insert([]float64{rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()})
	}

	points5, err := kdtree.Search([]float64{
		0.4, 0.7, 0.1, 0.6, 0.2,
		0.5, 0.8, 0.2, 0.7, 0.3})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("found %d points with k=5\n", len(points5))
	fmt.Printf("expected about 10\n")
}
