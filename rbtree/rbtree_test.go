package rbtree

import (
	"fmt"
	"testing"
)

func TestInsert1(t *testing.T) {
	tree := RedBlackTree{&Node{black, nil, nil, nil, 0}}

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(-1)
	tree.Insert(3)
	tree.Insert(-2)
	tree.Insert(-3)

	fmt.Println(tree)
}

func TestInsert2(t *testing.T) {
	tree := RedBlackTree{&Node{black, nil, nil, nil, 0}}
	for i := 0; i != 100; i++ {
		tree.Insert(i)
	}

	tree = RedBlackTree{&Node{black, nil, nil, nil, 0}}
	for i := 100; i != 0; i-- {
		tree.Insert(i)
	}
}

func TestRebalance1(t *testing.T) {
	sentinel := &Node{black, nil, nil, nil, 0}
	A := &Node{red, nil, nil, nil, 1}
	B := &Node{red, nil, nil, nil, 2}
	C := &Node{black, nil, nil, nil, 3}
	C.left = A
	C.right = sentinel
	C.p = sentinel
	B.left = sentinel
	B.right = sentinel
	B.p = A
	A.left = sentinel
	A.right = B
	A.p = C
	tree := RedBlackTree{C}
	tree.rebalanceInsert(B)
}
