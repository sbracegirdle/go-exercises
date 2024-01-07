package main

import "fmt"

// BTree represents a B-tree
type BTree struct {
	root *BTreeNode
	t    int // Minimum degree of the B-tree.
}

// BTreeNode represents a node in the B-tree
type BTreeNode struct {
	isLeaf bool
	keys   []int
	child  []*BTreeNode
}

// searchKey returns the index of the first key that is equal or greater than k in the keys array.
// If all keys are less than k, then it returns len(b.keys)
func (b *BTreeNode) searchKey(k int) int {
	idx := 0
	for idx < len(b.keys) && b.keys[idx] < k {
		idx++
	}
	return idx
}

// Search returns the node containing the key and the index of the key in the keys array.
func (b *BTreeNode) Search(k int) (*BTreeNode, int) {
	idx := b.searchKey(k)
	// if the key is found in this node, return this node and the index of the key.
	if idx < len(b.keys) && b.keys[idx] == k {
		return b, idx
	} else if b.isLeaf {
		// if the node is a leaf node and the key is not in this node, return nil.
		return nil, -1
	} else {
		// if the node is not a leaf, search the appropriate child node.
		return b.child[idx].Search(k)
	}
}

// Insert inserts a key into the tree.
func (t *BTree) Insert(k int) {
	root := t.root
	if len(root.keys) == (2*t.t)-1 {
		// Create a new root because the root is full
		temp := &BTreeNode{}
		t.root = temp
		temp.child = append(temp.child, root)
		// Split the old root and move 1 key to the new root
		t.splitChild(temp, 0)
		t.insertNonFull(temp, k)
	} else {
		// If the root is not full, call insertNonFull for the root
		t.insertNonFull(root, k)
	}
}

// insertNonFull inserts a key into a non-full node.
func (t *BTree) insertNonFull(x *BTreeNode, k int) {
	// Start from the rightmost key in the node
	i := len(x.keys) - 1

	// If the node is a leaf node
	if x.isLeaf {
		// Append a new key at the end of the keys
		x.keys = append(x.keys, 0)

		// Shift all keys greater than k to the right
		for i >= 0 && k < x.keys[i] {
			x.keys[i+1] = x.keys[i]
			i--
		}

		// Insert the new key at the found position
		x.keys[i+1] = k
	} else {
		// If the node is not a leaf, find the child which is going to hold the new key
		for i >= 0 && k < x.keys[i] {
			i--
		}
		i++

		// If the found child is full
		if len(x.child[i].keys) == (2*t.t)-1 {
			// Split the child
			t.splitChild(x, i)

			// After split, the middle key of the child moves up and the child is split into two.
			// Check which of the two children is going to hold the new key
			if k > x.keys[i] {
				i++
			}
		}

		// Insert the key into the child
		t.insertNonFull(x.child[i], k)
	}
}

// splitChild splits the child y of x.
func (t *BTree) splitChild(x *BTreeNode, i int) {
	// tt is the minimum degree of the B-tree
	tt := t.t

	// y is the i-th child of x that is going to be split
	y := x.child[i]

	// z is the new node, created to store tt-1 keys of y
	z := &BTreeNode{isLeaf: y.isLeaf}

	// Make space for the new child
	x.child = append(x.child, nil)
	copy(x.child[i+2:], x.child[i+1:])
	x.child[i+1] = z

	// Make space for the new key in x
	x.keys = append(x.keys, 0)
	copy(x.keys[i+1:], x.keys[i:])

	// Move the middle key of y to x
	x.keys[i] = y.keys[tt-1]

	// Move the last tt-1 keys of y to z
	z.keys = append(z.keys, y.keys[tt:]...)
	y.keys = y.keys[:tt-1]

	// If y is not a leaf, move the last tt children of y to z
	if !y.isLeaf {
		z.child = append(z.child, y.child[tt:]...)
		y.child = y.child[:tt]
	}
}

func createBTree(keys []int, t int) *BTree {
	tree := &BTree{
		root: &BTreeNode{
			isLeaf: true,
			keys:   make([]int, 0),
			child:  make([]*BTreeNode, 0),
		},
		t: t,
	}

	for _, key := range keys {
		tree.Insert(key)
	}

	return tree
}

func main() {
	fmt.Println("B-Tree implementation in Go.")
}
