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
	nodes := make([]*Node, len(items))
	var node, prevNode *Node
	for i := range items {
		node = &Node{nil, nil, &items[i]}
		if i != 0 {
			prevNode.next = node
		}
		nodes[i] = node
		prevNode = node
	}

	// Build layers until left with only a head node
	for len(nodes) != 1 {

		node = nodes[0]
		nodesAbove := []*Node{&Node{nil, node, node.item}}

		pos := 0
		for node.next != nil {
			if rand.Float64() < p {
				nodesAbove = append(nodesAbove, &Node{
					next:  nil,
					below: nodes[pos],
					item:  nodes[pos].item,
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
	node := n
	fmt.Printf("%5d ", n.item.key)
	for node.next != nil {
		fmt.Printf("%5d ", node.next.item.key)
		node = node.next
	}
	fmt.Print("\n")

	if n.below != nil {
		n.below.PrintKeys()
	}
}

func (n *Node) Get(key int) (*Item, error) {
	if n.item.key == key {
		return n.item, nil
	} else if n.item.key > key {
		return nil, errors.New("index error")
	} else if n.next == nil || n.next.item.key > key {
		return n.below.Get(key)
	} else {
		return n.next.Get(key)
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
func (n *Node) Insert(item *Item, p float64) error {

	// Handle the first case
	if item.key < n.item.key {
		n.below = insertLeftCol(item, n)

		// prune the previous left column
		pruneSecond(n.below, p)
		return nil
	}

	// Handle the second case
	nodeInsertedBelow := insertRight(item, n.below, p)
	if nodeInsertedBelow != nil {
		// The head node needs to be replaced because we permit only one node at the
		// top level (why?)
		n.below = &Node{n.next, n.below, n.item}
		n.next = nil
	}
	return nil
}

// insertRight is the recursive helper function called by Insert. It takes an
// item to insert, a node to the left of where the item should go, and a
// probability that the node will appear in the list index above.
//
// It returns a non-nil pointer to the inserted node iff a reference to the node
// should be added to the index list above.
func insertRight(item *Item, n *Node, p float64) (nodeInsertedBelow *Node) {
	if n.below != nil {
		nodeInsertedBelow = insertRight(item, n.below, p)
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

// pruneSecond looks at the item to the left of the head node (second item in
// the list) and decides retroactively whether to propagate it upward
func pruneSecond(n *Node, p float64) bool {
	if n.below == nil {
		return rand.Float64() < p
	}
	hoisted := pruneSecond(n.below, p)
	if hoisted {
		return rand.Float64() < p
	}
	n.next = n.next.next // if the test below failed, delete the reference to the right
	return false
}

// insertLeftCol adds a new tower of nodes to the left side of the skip-list.
// Unlike other items in the list, the left most item is guaranteed to alway
// bubble up to the next level
func insertLeftCol(item *Item, n *Node) (below *Node) {
	if n.below == nil {
		below = &Node{n.next, nil, item}
	} else {
		below = insertLeftCol(item, n.below)
		below = &Node{below.next, below.below, item}
	}
	return
}
