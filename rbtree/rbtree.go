/*
 * Red-black tree
 *
 * A red-black tree is a balanced binary tree, which guarantees O(log n)
 * complexity for adding, deleting, and searching for elements.
 *
 * In a red-black tree, the following conditions are met:
 *
 * - every node cas a color, red or black
 * - the root node is black
 * - every leaf is black
 * - the children of red nodes are black
 * - every downward path from a node has an equal number of black nodes
 *
 * Compared to a AVL tree, a red-black tree has a larger constant factor
 * on searching due to it being less strictly balanced. On the other hand,
 * the red-black tree has a smaller constant factor on insertion and deletion,
 * and so is more suited for volatile data.
 *
 * The implementation below works for integer keys, but is easily modified
 * for any other orderable key. No special attention is given to duplicate
 * keys.
 */

package rbtree

const (
	red   = iota
	black = iota
)

// Color is an integer coding the color of a red-black tree node
type Color int

// Node represents a rd-black tree node
type Node struct {
	color Color
	left  *Node
	right *Node
	p     *Node
	key   int
}

// RedBlackTree represents a red-black tree
type RedBlackTree struct {
	root *Node
}

// isSentinel returns true when a node represents a sentinal node
func (n *Node) isSentinel() bool {
	return n.left == nil && n.right == nil && n.p == nil
}

// Rotations
//
// Rotations are used to alter the pointer structure within the tree, which is
// necessary to keep the tree balanced following insertion and deletion
// operations. Rotations can be to the "left" or to the "right". In a left
// rotation:
//
//       [n]                          [y]
//      /   \                        /   \
//     a    [y]        becomes     [n]    c
//         /   \                  /   \
//        b     c                a     b
//
// and in a right rotation, the operation is reversed. Note that the order
// of child nodes a, b, and c remains the same, but that a is now deeper, and c
// is now shallower.
//
// A binary search tree with n nodes can be transformed into any other binary
// search tree in O(n) rotations.
//
// Implementation note:
// Node [n] might be the root of the tree. In this case, the rotation replaces
// the tree root. As the tree struct is not modified by  this method, it returns
// a non-nil pointer to the new root in the event that it changed. The caller is
// responsible for checking whether the returned pointer is non-nil, and if so,
// for updating the tree root pointer.
func (n *Node) rotateLeft() *Node {
	var root *Node
	y := n.right
	n.right = y.left
	if !y.left.isSentinel() {
		y.left.p = n
	}
	y.p = n.p
	if n.p.isSentinel() {
		root = y
	} else if n.p.left == n {
		n.p.left = y
	} else {
		n.p.right = y
	}
	y.left = n
	n.p = y
	return root
}

func (n *Node) rotateRight() *Node {
	var root *Node
	y := n.left
	n.left = y.right
	if !y.right.isSentinel() {
		y.right.p = n
	}
	y.p = n.p
	if n.p.isSentinel() {
		root = y
	} else if n.p.right == n {
		n.p.right = y
	} else {
		n.p.left = y
	}
	y.right = n
	n.p = y
	return root
}

// Insert adds a node with value *key* to a red black tree
// This proceeds exactly the same as in an ordinary binary search tree, except
// that the inserted node is given a color (red) and the tree is rebalanced
// afterward to restore red-black properties by calling
// `RedBlackTree.rebalanceInsert()`
func (tree *RedBlackTree) Insert(key int) {
	childNode := tree.root
	parentNode := &Node{black, nil, nil, nil, 0}
	var newNode *Node
	// Follow tree until a leaf node is found
	for !childNode.isSentinel() {
		parentNode = childNode
		if key < childNode.key {
			childNode = childNode.left
		} else {
			childNode = childNode.right
		}
	}
	newNode = &Node{red, nil, nil, parentNode, key}
	if parentNode.isSentinel() {
		// This can only happen when childNode is the root node, i.e. the tree is empty
		tree.root = newNode
	} else if key < parentNode.key {
		parentNode.left = newNode
	} else {
		parentNode.right = newNode
	}
	// Place sentinel nodes below newNode
	newNode.left = &Node{black, nil, nil, nil, 0}
	newNode.right = &Node{black, nil, nil, nil, 0}
	tree.rebalanceInsert(newNode)
}

// rebalanceInsert restores red-black properties to a tree following the
// insertion of a new node.
//
// As the inserted node (which is the function parameter) is red, one of two
// things may have happened. If the tree was empty, then the new node because
// the root, and it is not black, which a red-black tree requires.
// Alternatively, if the tree was not empty, but the parent of the new node was
// red, then there is a sequence of red nodes, which is also a red-black
// violation.
//
// In the first case, the loop is skipped, the root color is set to black, and
// the tree is now valid. In the second case, restoration of red-black
// properties is somewhat more involved.
func (tree *RedBlackTree) rebalanceInsert(z *Node) {
	var y, t *Node

	// With every cycle of this loop, one of two things will happen.
	//
	// 1. The pointer to the active node moves to its parent
	// 2. Rotations are performed
	//
	// The loop stops when the parent of z is black which, since z is red, means
	// that the requirement that children of red nodes are black is upheld.
	for z.p.color == red {
		if z.p == z.p.p.left {
			// The parent of z is the left child of its parent.
			// y is the uncle of z
			y = z.p.p.right
			if y.color == red {
				// This is the first possible case: z's uncle is red
				// The situation looks like this:
				//
				//        black
				//      /      \
				//    red      red (y)
				//    /
				//  red (z)
				//
				// We fix it by setting the parent of z and the uncle to be black, and
				// refocusing on the grandparent of z, which we paint red.
				z.p.color = black
				y.color = black
				z.p.p.color = red
				z = z.p.p
			} else if z == z.p.right {
				// This is the second case, which is that the uncle is black and z is
				// a right child.
				//
				//        black
				//      /      \
				//    red      black (y)
				//      \
				//     red (z)
				//
				// By rotating z's parent, we obtain
				//
				//        black
				//      /      \
				//    red (t)  black (y)
				//    /
				//  red (z)
				//
				// which is case 3.
				z = z.p
				t = z.rotateLeft()
				if t != nil {
					tree.root = t
				}
			} else {
				// In the third case, the uncle is black and z is a left child.
				//
				//        black
				//      /      \
				//    red (p)  black (y)
				//    /
				//  red (z)
				//
				// A right rotation on z's grandparent accompanied by a color swap fixes
				// the situation.
				//
				//      black (p)
				//      /       \
				//    red (z)   red
				//                \
				//              black (y)
				//
				// This completes the rebalancing.
				z.p.color = black
				z.p.p.color = red
				t = z.p.p.rotateRight()
				if t != nil {
					tree.root = t
				}
			}
		} else {
			// This mirrors the logic from above with the tree flipped
			y = z.p.p.left
			if y.color == red {
				z.p.color = black
				y.color = black
				z.p.p.color = red
				z = z.p.p
			} else if z == z.p.left {
				z = z.p
				t = z.rotateRight()
				if t != nil {
					tree.root = t
				}
			} else {
				z.p.color = black
				z.p.p.color = red
				t = z.p.p.rotateLeft()
				if t != nil {
					tree.root = t
				}
			}
		}
	}
	tree.root.color = black
}

// Delete removes a value from the red-black tree
func (tree *RedBlackTree) Delete(key int) {
	// Not implemented
	var z *Node
	tree.rebalanceDelete(z)
}

func (tree *RedBlackTree) rebalanceDelete(z *Node) {
	// Not implemented
}
