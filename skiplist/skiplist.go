/*
 * Package skiplist implements a skip-list. A skip-list is a probabalistic data
 * structure similar to a binary tree. Unlike a balanced binary tree (see
 * package rbtree), a skip-list is approximately balanced.
 *
 * The bottom layer of a skip-list is a linked list (singly-linked, in this
 * implementation). Subsequent layers are built on top of this layer by
 * including each list node with a probability *p*.

 *    L3  *  (head node)
 *    L2  *                    *             *          (sparse linked lists of
 *    L1  *      *             *             *     *     randomly included nodes)
 *    L0  ********************************************* (data in a linked list)

 * When searching the skip-list, we start at the top (head) node, and move
 * downward and to the right. When moving to the right would cause the key of
 * the next node to exceed the key we're searching for, we move downward
 * instead.
 */

package skiplist

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

type Item struct {
	key   int
	value interface{}
}

type ItemSlice []Item

func (items ItemSlice) Len() int {
	return len(items)
}

func (items ItemSlice) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items ItemSlice) Less(i, j int) bool {
	return items[i].key < items[j].key
}

type Node struct {
	next  *Node
	below *Node
	item  *Item
}

// Depth indicates what level a node is on in the skip-list, with 0 denoting the base (data) level
func (n *Node) Depth() int {
	if n.below == nil {
		return 0
	}
	return 1 + n.below.Depth()
}

// New assembles a skip-list from a list of Items, where each node is contained
// in the layer above with probability p. The final layer contains a single head
// node, which is the return value.
func New(items ItemSlice, p float64) *Node {
	if !sort.IsSorted(items) {
		sort.Sort(items)
	}

	// build the bottom layer
	nodes := make([]*Node, len(items)+1)
	nodes[0] = &Node{nil, nil, nil}

	for i := range items {
		nodes[i+1] = &Node{nil, nil, &items[i]}
		nodes[i].next = nodes[i+1]
	}

	// Build layers until left with only a head node
	var node *Node
	for len(nodes) != 1 {

		node = nodes[1]
		nodesAbove := []*Node{&Node{nil, nodes[0], nil}}

		pos := 1
		for node != nil {
			if rand.Float64() < p {
				nodesAbove = append(nodesAbove, &Node{
					next:  nil,
					below: node,
					item:  node.item,
				})
				nodesAbove[len(nodesAbove)-2].next = nodesAbove[len(nodesAbove)-1]
			}
			node = node.next
			pos++
		}

		nodes = nodesAbove
	}

	return nodes[0]
}

func (n *Node) PrintKeys() {
	fmt.Print("*")
	node := n.next
	for node != nil {
		fmt.Printf("%5d ", node.item.key)
		node = node.next
	}
	fmt.Print("\n")

	if n.below != nil {
		n.below.PrintKeys()
	}
}

// Get returns an item from the skip-list by key, or a non-nil error if the item is not found
func (n *Node) Get(key int) (*Item, error) {
	if n.below.next.item.key > key {
		return n.below.Get(key)
	}
	return get(n.below.next, key)
}

func get(n *Node, key int) (*Item, error) {
	// check for a direct shortcut
	if n.item.key == key {
		return n.item, nil
	}

	// linear search on data layer
	if n.below == nil {
		if n.next == nil || n.next.item.key > key {
			return nil, errors.New("index error")
		}
		return get(n.next, key) // TODO: expand as loop
	}

	// choose whether to move down or right
	if n.next == nil || n.next.item.key > key {
		return get(n.below, key)
	} else {
		return get(n.next, key)
	}
}

// Insert adds a new item to the skip-list. There are two important
// possibilities: - the item may be to the left of the head node - the item may
// be to the right of the head node
//
// In the first case, we replace the head node, and then run up the tower of
// nodes below the previous head and determine whether each node above the base
// layer should continue to exist.
//
// In the second case, we call a recursive function that inserts the item into
// the bottom most linked list and then bubbles up whether the node is kept as
// an index in the layer above.
func (head *Node) Insert(item *Item, p float64) error {

	// Handle the second case
	nodeInsertedBelow := insert(item, head.below, p)
	if nodeInsertedBelow != nil {
		// The head node needs to be replaced because we permit only one node at the
		// top level (why?)
		head.below = &Node{head.next, head.below, head.item}
		head.next = nil
	}
	return nil
}

// insert is the recursive helper function called by Insert. It takes an item to
// insert, a node to the left of where the item should go, and a probability
// that the node will appear in the list index above.
//
// It returns a non-nil pointer to the inserted node iff a reference to the node
// should be added to the index list above.
func insert(item *Item, n *Node, p float64) (nodeInsertedBelow *Node) {
	if n.below != nil {
		nodeInsertedBelow = insert(item, n.below, p)
	}

	if n.below == nil {
		// This is the data level, so create node to insert here
		nodeInsertedBelow = &Node{nil, nil, item}
	}

	if nodeInsertedBelow != nil {
		for n.next != nil && n.next.item.key < item.key {
			n = n.next
		}
		n.next = &Node{n.next, nodeInsertedBelow.below, item}
		if rand.Float64() >= p {
			nodeInsertedBelow = nil
		}
	}
	return
}
