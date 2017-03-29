/*
 * A linked list is a way to chain items together using pointers.
 * Searching for an item in a linked list is an O(n) operation.
 * Adding a node to the beginning of a linked list is O(1), but adding
 * to the end or inserting in the middle is O(n). Removing a node from
 * a linked list is O(n), unless it is the first node, which is O(1).
 */

package linkedlist

import "errors"

var INDEX_ERROR = errors.New("out-of-range index error")

// Node is a linked list item
type Node struct {
	Prev  *Node
	Next  *Node
	Value interface{}
}

// LinkedList contains the header Node of an acyclic doubly-linked list
type LinkedList struct {
	Head   *Node
	length int
}

// New creates a new LinkedList with *initialValue* in the prev position
func New() *LinkedList {
	return &LinkedList{nil, 0}
}

// Length returns the length of a linked list
func (lst *LinkedList) Length() int {
	return lst.length
}

// Get returns the value at position *index*.
// If *index* is out of bounds, returns an error.
func (lst *LinkedList) Get(index int) (interface{}, error) {
	node := lst.Head
	if node == nil {
		return 0, INDEX_ERROR
	}
	if index < 0 || index >= lst.length {
		return 0, INDEX_ERROR
	}
	for i := 0; i != index; i++ {
		node = node.Next
	}
	return node.Value, nil
}

// Set sets the value at position *index*
// If *index* is out of bounds, returns an error.
func (lst *LinkedList) Set(index int, value interface{}) error {
	node := lst.Head
	if node == nil {
		return INDEX_ERROR
	}
	if index < 0 || index >= lst.length {
		return INDEX_ERROR
	}
	for i := 0; i != index; i++ {
		node = node.Next
	}
	node.Value = value
	return nil
}

// Append adds a node to the end of the linked list and returns
// the new length
func (lst *LinkedList) Append(value interface{}) int {
	if lst.Head == nil {
		lst.Head = &Node{nil, nil, value}
		lst.length++
		return 1
	}

	node := lst.Head
	index := 0
	for node.Next != nil {
		node = node.Next
		index++
	}
	node.Next = &Node{node, nil, value}
	lst.length++
	return lst.length
}

// Prepend adds a node to the beginning of the linked list and
// returns the new list length
func (lst *LinkedList) Prepend(value interface{}) int {
	if lst.Head == nil {
		lst.Head = &Node{nil, nil, value}
		lst.length++
		return 0
	}

	node := lst.Head
	lst.Head = &Node{nil, node, value}
	node.Prev = lst.Head
	lst.length++
	return lst.length
}

// Insert places a new Node in the middle of a linked list, or returns an error
func (lst *LinkedList) Insert(index int, value interface{}) error {
	if index < 0 || index >= lst.length {
		return INDEX_ERROR
	}

	node := lst.Head
	for i := 1; i != index; i++ {
		node = node.Next
	}

	newNode := &Node{node, node.Next, value}
	if node.Next != nil {
		node.Next.Prev = newNode
	}
	node.Next = newNode
	lst.length++
	return nil
}

// Delete removes the node at *index* and returns the deleted
// nodes' value. If *index* is out of bounds, returns an error.
func (lst *LinkedList) Delete(index int) (interface{}, error) {
	if lst.Head == nil {
		return 0, INDEX_ERROR
	}
	if index < 0 {
		return 0, INDEX_ERROR
	}

	node := lst.Head

	if index == 0 {
		lst.Head = lst.Head.Next
		lst.length = 0
		return node.Value, nil
	}

	for i := 0; i != index; i++ {
		if node.Next == nil {
			return 0, INDEX_ERROR
		}
		node = node.Next
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	lst.length--
	return node.Value, nil
}
