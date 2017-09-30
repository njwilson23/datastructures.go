package skiplist

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

/*
 A skip-list is a probabalistic data structure similar to a binary tree.
 Unlike a balanced binary tree (see rbtree), a skip-list is approximately
 balanced.

 The bottom layer of a skip-list is a linked list (singly-linked, in this
 implementation). Subsequent layers are built on top of this layer by
 including each list node with a probability *p*.

	L3	head
	L2	*					 *			   *
	L1	*      *             *             *     *	  	(randomly included nodes)
	L0	********************************************* 	(data in a linked list)

 When searching the skip-list, we start at the top (head) node, and move
 downward and to the right. When moving to the right would cause the key of
 the next node to exceed the key we're searching for, we move downward
 instead.
*/

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

// NewSkipList assembles a skip-list from a list of Items, where each node is
// contained in the layer above with probability p. The final layer contains a
// single head node, which is the return value.
func NewSkipList(items ItemSlice, p float64) *Node {

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
