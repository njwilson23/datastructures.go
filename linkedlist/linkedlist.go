/*
 * A linked list is a way to chain items together using pointers.
 * Searching for an item in a linked list is an O(n) operation.
 * Adding a node to the beginning of a linked list is O(1), but adding
 * to the end or inserting in the middle is O(n). Removing a node from
 * a linked list is O(n), unless it is the first node, which is O(1).
 */

package linkedlist

import (
	"errors"
)

// Node is a linked list item
type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

// LinkedList contains the header Node of an acyclic doubly-linked list
type LinkedList struct {
	head   *Node
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
	node := lst.head
	if node == nil {
		return 0, errors.New("empty list")
	}
	if index < 0 || index >= lst.length {
		return 0, errors.New("index error")
	}
	for i := 0; i != index; i++ {
		node = node.next
	}
	return node.value, nil
}

// Set sets the value at position *index*
// If *index* is out of bounds, returns an error.
func (lst *LinkedList) Set(index int, value interface{}) error {
	node := lst.head
	if node == nil {
		return errors.New("empty list")
	}
	if index < 0 || index >= lst.length {
		return errors.New("index error")
	}
	for i := 0; i != index; i++ {
		node = node.next
	}
	node.value = value
	return nil
}

// Append adds a node to the end of the linked list and returns
// the new length
func (lst *LinkedList) Append(value interface{}) int {
	if lst.head == nil {
		lst.head = &Node{nil, nil, value}
		lst.length++
		return 1
	}

	node := lst.head
	index := 0
	for node.next != nil {
		node = node.next
		index++
	}
	node.next = &Node{node, nil, value}
	lst.length++
	return lst.length
}

// Prepend adds a node to the beginning of the linked list and
// returns the new list length
func (lst *LinkedList) Prepend(value interface{}) int {
	if lst.head == nil {
		lst.head = &Node{nil, nil, value}
		lst.length++
		return 0
	}

	node := lst.head
	lst.head = &Node{nil, node, value}
	node.prev = lst.head
	lst.length++
	return lst.length
}

// Insert places a new Node in the middle of a linked list, or returns an error
func (lst *LinkedList) Insert(index int, value interface{}) error {
	if index < 0 || index >= lst.length {
		return errors.New("index error")
	}

	node := lst.head
	for i := 1; i != index; i++ {
		node = node.next
	}

	newNode := &Node{node, node.next, value}
	if node.next != nil {
		node.next.prev = newNode
	}
	node.next = newNode
	lst.length++
	return nil
}

// Delete removes the node at *index* and returns the deleted
// nodes' value. If *index* is out of bounds, returns an error.
func (lst *LinkedList) Delete(index int) (interface{}, error) {
	if lst.head == nil {
		return 0, errors.New("empty list")
	}
	if index < 0 {
		return 0, errors.New("index may not be negative")
	}

	node := lst.head
	for i := 0; i != index; i++ {
		if node.next == nil {
			return 0, errors.New("index error")
		}
		node = node.next
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	lst.length--
	return node.value, nil
}
